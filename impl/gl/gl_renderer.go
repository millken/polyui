//go:build glfw
// +build glfw

package gl

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"image"
	"image/color"
	"image/draw"
	_ "image/jpeg"
	"image/png"
	"os"
	"sync"

	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"

	gl "github.com/go-gl/gl/v4.6-core/gl"
	"github.com/millken/polyui/core"
)

// SimpleGLRenderer is a minimal OpenGL renderer for PolyUI prototype.
type SimpleGLRenderer struct {
	prog      uint32
	vao       uint32
	vbo       uint32
	textures  map[string]uint32
	texSizes  map[string][2]int
	fbW, fbH  int32
	useTexLoc int32
	// LRU cache metadata
	texOrder []string // most-recent at end
	texCap   int
	mu       sync.Mutex
}

func (s *SimpleGLRenderer) Init() error {
	// compile simple shaders
	vert := `#version 330 core
layout(location = 0) in vec2 pos;
layout(location = 1) in vec4 col;
layout(location = 2) in vec2 uv;
out vec4 vcol;
out vec2 vuv;
uniform int u_useTex;
void main() { vcol = col; vuv = uv; gl_Position = vec4(pos, 0.0, 1.0); }`
	frag := `#version 330 core
in vec4 vcol; in vec2 vuv; out vec4 FragColor; uniform sampler2D u_tex; uniform int u_useTex;
void main(){ if(u_useTex==1) { FragColor = texture(u_tex, vuv); } else { FragColor = vcol; } }`

	// compile and link program via helper
	p, err := createProgram(vert, frag)
	if err != nil {
		return err
	}
	s.prog = p
	s.useTexLoc = gl.GetUniformLocation(s.prog, gl.Str("u_useTex\x00"))

	// setup simple VAO/VBO for a quad (will update buffer per draw for simplicity)
	gl.GenVertexArrays(1, &s.vao)
	gl.GenBuffers(1, &s.vbo)
	s.textures = make(map[string]uint32)
	s.texSizes = make(map[string][2]int)
	s.texOrder = make([]string, 0, 64)
	s.texCap = 128 // default max textures

	return nil
}

func (s *SimpleGLRenderer) BeginFrame(width, height, pixelRatio float32) {
	gl.Viewport(0, 0, int32(width), int32(height))
	gl.ClearColor(0.1, 0.1, 0.1, 1.0)
	gl.Clear(gl.COLOR_BUFFER_BIT)
	s.fbW = int32(width)
	s.fbH = int32(height)
}

func (s *SimpleGLRenderer) EndFrame() {}

// Rect draws an axis-aligned rectangle in screen coordinates (pixels).
func (s *SimpleGLRenderer) Rect(x, y, w, h float32, color core.Color) {
	// convert to NDC clip space
	nx := func(px float32) float32 { return px*2.0/float32(s.fbW) - 1.0 }
	ny := func(py float32) float32 { return 1.0 - py*2.0/float32(s.fbH) }
	verts := []float32{
		nx(x), ny(y), color.R, color.G, color.B, color.A, 0, 0,
		nx(x + w), ny(y), color.R, color.G, color.B, color.A, 1, 0,
		nx(x + w), ny(y + h), color.R, color.G, color.B, color.A, 1, 1,
		nx(x), ny(y + h), color.R, color.G, color.B, color.A, 0, 1,
	}
	gl.BindVertexArray(s.vao)
	gl.BindBuffer(gl.ARRAY_BUFFER, s.vbo)
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(verts), gl.Ptr(verts), gl.STREAM_DRAW)

	gl.UseProgram(s.prog)
	// position: location 0, vec2; color: location 1, vec4
	gl.EnableVertexAttribArray(0)
	gl.VertexAttribPointer(0, 2, gl.FLOAT, false, 8*4, gl.PtrOffset(0))
	gl.EnableVertexAttribArray(1)
	gl.VertexAttribPointer(1, 4, gl.FLOAT, false, 8*4, gl.PtrOffset(2*4))
	gl.EnableVertexAttribArray(2)
	gl.VertexAttribPointer(2, 2, gl.FLOAT, false, 8*4, gl.PtrOffset(6*4))

	// draw quad as triangle fan
	gl.DrawArrays(gl.TRIANGLE_FAN, 0, 4)

	gl.DisableVertexAttribArray(0)
	gl.DisableVertexAttribArray(1)
	gl.DisableVertexAttribArray(2)
}

// loadTexture uploads an image file to GL and caches the texture id and size.
// helper: decode image.Image into RGBA
func decodeToRGBA(img image.Image) *image.RGBA {
	rgba := image.NewRGBA(img.Bounds())
	draw.Draw(rgba, rgba.Bounds(), img, image.Point{0, 0}, draw.Src)
	return rgba
}

// ensureCapacity evicts least-recent textures when over capacity
func (s *SimpleGLRenderer) ensureCapacity() {
	for len(s.texOrder) > s.texCap {
		evict := s.texOrder[0]
		s.texOrder = s.texOrder[1:]
		if tex, ok := s.textures[evict]; ok {
			gl.DeleteTextures(1, &tex)
			delete(s.textures, evict)
			delete(s.texSizes, evict)
		}
	}
}

// touch marks a key as recently used
func (s *SimpleGLRenderer) touch(key string) {
	// find and remove if exists
	for i := len(s.texOrder) - 1; i >= 0; i-- {
		if s.texOrder[i] == key {
			s.texOrder = append(s.texOrder[:i], s.texOrder[i+1:]...)
			break
		}
	}
	s.texOrder = append(s.texOrder, key)
}

// loadTextureKey loads texture either from path or from provided bytes key
func (s *SimpleGLRenderer) loadTextureKey(key string, data []byte) (uint32, int, int, error) {
	s.mu.Lock()
	if t, ok := s.textures[key]; ok {
		s.touch(key)
		sz := s.texSizes[key]
		s.mu.Unlock()
		return t, sz[0], sz[1], nil
	}
	s.mu.Unlock()

	var rgba *image.RGBA
	if len(data) > 0 {
		img, _, err := image.Decode(bytes.NewReader(data))
		if err != nil {
			return 0, 0, 0, err
		}
		rgba = decodeToRGBA(img)
	} else {
		f, err := os.Open(key)
		if err != nil {
			return 0, 0, 0, err
		}
		defer f.Close()
		img, _, err := image.Decode(f)
		if err != nil {
			return 0, 0, 0, err
		}
		rgba = decodeToRGBA(img)
	}

	var tex uint32
	gl.GenTextures(1, &tex)
	gl.BindTexture(gl.TEXTURE_2D, tex)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	w := int32(rgba.Rect.Dx())
	h := int32(rgba.Rect.Dy())
	if len(rgba.Pix) > 0 {
		gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, w, h, 0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(&rgba.Pix[0]))
	} else {
		gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, w, h, 0, gl.RGBA, gl.UNSIGNED_BYTE, nil)
	}

	s.mu.Lock()
	s.textures[key] = tex
	s.texSizes[key] = [2]int{int(w), int(h)}
	s.touch(key)
	s.ensureCapacity()
	s.mu.Unlock()
	return tex, int(w), int(h), nil
}

func (s *SimpleGLRenderer) Image(img core.Image, x, y, w, h float32) {
	var key string
	var data []byte
	if len(img.Data) > 0 {
		sum := sha256.Sum256(img.Data)
		key = "mem:" + hex.EncodeToString(sum[:])
		data = img.Data
	} else if img.Path != "" {
		key = img.Path
	} else {
		return
	}
	tex, _, _, err := s.loadTextureKey(key, data)
	if err != nil {
		core.Debugf("failed to load texture %s: %v\n", key, err)
		return
	}
	gl.UseProgram(s.prog)
	gl.Uniform1i(s.useTexLoc, 1)
	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, tex)

	nx := func(px float32) float32 { return px*2.0/float32(s.fbW) - 1.0 }
	ny := func(py float32) float32 { return 1.0 - py*2.0/float32(s.fbH) }
	verts := []float32{
		nx(x), ny(y), 1, 1, 1, 1, 0, 0,
		nx(x + w), ny(y), 1, 1, 1, 1, 1, 0,
		nx(x + w), ny(y + h), 1, 1, 1, 1, 1, 1,
		nx(x), ny(y + h), 1, 1, 1, 1, 0, 1,
	}
	gl.BindVertexArray(s.vao)
	gl.BindBuffer(gl.ARRAY_BUFFER, s.vbo)
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(verts), gl.Ptr(verts), gl.STREAM_DRAW)
	gl.EnableVertexAttribArray(0)
	gl.VertexAttribPointer(0, 2, gl.FLOAT, false, 8*4, gl.PtrOffset(0))
	gl.EnableVertexAttribArray(1)
	gl.VertexAttribPointer(1, 4, gl.FLOAT, false, 8*4, gl.PtrOffset(2*4))
	gl.EnableVertexAttribArray(2)
	gl.VertexAttribPointer(2, 2, gl.FLOAT, false, 8*4, gl.PtrOffset(6*4))
	gl.DrawArrays(gl.TRIANGLE_FAN, 0, 4)
	gl.DisableVertexAttribArray(0)
	gl.DisableVertexAttribArray(1)
	gl.DisableVertexAttribArray(2)
	gl.Uniform1i(s.useTexLoc, 0)
}

func (s *SimpleGLRenderer) DeleteImage(img core.Image) {
	var key string
	if len(img.Data) > 0 {
		sum := sha256.Sum256(img.Data)
		key = "mem:" + hex.EncodeToString(sum[:])
	} else {
		key = img.Path
	}
	s.mu.Lock()
	if tex, ok := s.textures[key]; ok {
		gl.DeleteTextures(1, &tex)
		delete(s.textures, key)
		delete(s.texSizes, key)
		// remove from order
		for i := range s.texOrder {
			if s.texOrder[i] == key {
				s.texOrder = append(s.texOrder[:i], s.texOrder[i+1:]...)
				break
			}
		}
	}
	s.mu.Unlock()
}

// Text is implemented below; the earlier empty stub was removed to avoid duplicate declaration.
func (s *SimpleGLRenderer) Push()                          {}
func (s *SimpleGLRenderer) Pop()                           {}
func (s *SimpleGLRenderer) PushScissor(x, y, w, h float32) {}
func (s *SimpleGLRenderer) PopScissor()                    {}

// renderTextToPNGBytes renders given text into an RGBA image using basicfont and returns PNG bytes and pixel size.
func renderTextToPNGBytes(text string, col color.RGBA, scale float32) ([]byte, int, int, error) {
	face := basicfont.Face7x13

	// measure width
	adv := font.MeasureString(face, text)
	w := adv.Round()
	metrics := face.Metrics()
	h := (metrics.Ascent + metrics.Descent).Round()
	if w <= 0 {
		w = 1
	}
	if h <= 0 {
		h = 1
	}

	rgba := image.NewRGBA(image.Rect(0, 0, w, h))
	// transparent background
	draw.Draw(rgba, rgba.Bounds(), &image.Uniform{C: color.RGBA{0, 0, 0, 0}}, image.Point{}, draw.Src)

	d := &font.Drawer{
		Dst:  rgba,
		Src:  image.NewUniform(col),
		Face: face,
		Dot:  fixed.P(0, face.Metrics().Ascent.Round()),
	}
	d.DrawString(text)

	var buf bytes.Buffer
	if err := png.Encode(&buf, rgba); err != nil {
		return nil, 0, 0, err
	}
	return buf.Bytes(), w, h, nil
}

// Text renders text by rasterizing it into an RGBA image, uploading as a texture (cached) and drawing a quad.
func (s *SimpleGLRenderer) Text(x, y float32, text string, size float32, col core.Color) {
	// convert color
	c := color.RGBA{uint8(col.R * 255), uint8(col.G * 255), uint8(col.B * 255), uint8(col.A * 255)}

	data, tw, th, err := renderTextToPNGBytes(text, c, size)
	if err != nil {
		core.Debugf("failed to render text: %v\n", err)
		return
	}

	// key by hash of bytes
	sum := sha256.Sum256(data)
	key := "text:" + hex.EncodeToString(sum[:])
	tex, _, _, err := s.loadTextureKey(key, data)
	if err != nil {
		core.Debugf("failed to upload text texture: %v\n", err)
		return
	}

	// draw it using Image path (reuse quad drawing code)
	gl.UseProgram(s.prog)
	gl.Uniform1i(s.useTexLoc, 1)
	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, tex)

	// render at given x,y with pixel size tw x th
	nx := func(px float32) float32 { return px*2.0/float32(s.fbW) - 1.0 }
	ny := func(py float32) float32 { return 1.0 - py*2.0/float32(s.fbH) }

	fw := float32(tw)
	fh := float32(th)
	verts := []float32{
		nx(x), ny(y), 1, 1, 1, 1, 0, 0,
		nx(x + fw), ny(y), 1, 1, 1, 1, 1, 0,
		nx(x + fw), ny(y + fh), 1, 1, 1, 1, 1, 1,
		nx(x), ny(y + fh), 1, 1, 1, 1, 0, 1,
	}
	gl.BindVertexArray(s.vao)
	gl.BindBuffer(gl.ARRAY_BUFFER, s.vbo)
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(verts), gl.Ptr(verts), gl.STREAM_DRAW)
	gl.EnableVertexAttribArray(0)
	gl.VertexAttribPointer(0, 2, gl.FLOAT, false, 8*4, gl.PtrOffset(0))
	gl.EnableVertexAttribArray(1)
	gl.VertexAttribPointer(1, 4, gl.FLOAT, false, 8*4, gl.PtrOffset(2*4))
	gl.EnableVertexAttribArray(2)
	gl.VertexAttribPointer(2, 2, gl.FLOAT, false, 8*4, gl.PtrOffset(6*4))
	gl.DrawArrays(gl.TRIANGLE_FAN, 0, 4)
	gl.DisableVertexAttribArray(0)
	gl.DisableVertexAttribArray(1)
	gl.DisableVertexAttribArray(2)
	gl.Uniform1i(s.useTexLoc, 0)
}
