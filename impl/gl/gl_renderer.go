//go:build glfw
// +build glfw

package gl

import (
	"fmt"
	"image"
	"image/draw"
	_ "image/jpeg"
	_ "image/png"
	"os"

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

	vsrc, freeV := gl.Strs(vert + "\x00")
	defer freeV()
	fsrc, freeF := gl.Strs(frag + "\x00")
	defer freeF()

	v := gl.CreateShader(gl.VERTEX_SHADER)
	gl.ShaderSource(v, 1, vsrc, nil)
	gl.CompileShader(v)
	var status int32
	gl.GetShaderiv(v, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		return fmt.Errorf("vertex shader compile failed")
	}

	f := gl.CreateShader(gl.FRAGMENT_SHADER)
	gl.ShaderSource(f, 1, fsrc, nil)
	gl.CompileShader(f)
	gl.GetShaderiv(f, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		return fmt.Errorf("fragment shader compile failed")
	}

	p := gl.CreateProgram()
	gl.AttachShader(p, v)
	gl.AttachShader(p, f)
	gl.LinkProgram(p)
	gl.GetProgramiv(p, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		return fmt.Errorf("program link failed")
	}

	s.prog = p
	s.useTexLoc = gl.GetUniformLocation(s.prog, gl.Str("u_useTex\x00"))

	// setup simple VAO/VBO for a quad (will update buffer per draw for simplicity)
	gl.GenVertexArrays(1, &s.vao)
	gl.GenBuffers(1, &s.vbo)
	s.textures = make(map[string]uint32)
	s.texSizes = make(map[string][2]int)

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
func (s *SimpleGLRenderer) loadTexture(path string) (uint32, int, int, error) {
	if t, ok := s.textures[path]; ok {
		sz := s.texSizes[path]
		return t, sz[0], sz[1], nil
	}
	f, err := os.Open(path)
	if err != nil {
		return 0, 0, 0, err
	}
	defer f.Close()
	img, _, err := image.Decode(f)
	if err != nil {
		return 0, 0, 0, err
	}
	rgba := image.NewRGBA(img.Bounds())
	draw.Draw(rgba, rgba.Bounds(), img, image.Point{0, 0}, draw.Src)

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
	s.textures[path] = tex
	s.texSizes[path] = [2]int{int(w), int(h)}
	return tex, int(w), int(h), nil
}

func (s *SimpleGLRenderer) Image(img core.Image, x, y, w, h float32) {
	if img.Path == "" {
		return
	}
	tex, _, _, err := s.loadTexture(img.Path)
	if err != nil {
		fmt.Printf("failed to load texture %s: %v\n", img.Path, err)
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
	if tex, ok := s.textures[img.Path]; ok {
		gl.DeleteTextures(1, &tex)
		delete(s.textures, img.Path)
		delete(s.texSizes, img.Path)
	}
}

func (s *SimpleGLRenderer) Text(x, y float32, text string, size float32, color core.Color) {}
func (s *SimpleGLRenderer) Push()                                                          {}
func (s *SimpleGLRenderer) Pop()                                                           {}
func (s *SimpleGLRenderer) PushScissor(x, y, w, h float32)                                 {}
func (s *SimpleGLRenderer) PopScissor()                                                    {}
