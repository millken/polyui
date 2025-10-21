package gl

import (
	"fmt"

	gl "github.com/go-gl/gl/v4.6-core/gl"
	"github.com/millken/polyui/core"
)

// createProgram compiles vertex and fragment shaders, links them into a program,
// prints info logs on failure and returns the program id or an error.
func createProgram(vertSrc, fragSrc string) (uint32, error) {
	vsrc, freeV := gl.Strs(vertSrc + "\x00")
	defer freeV()
	fsrc, freeF := gl.Strs(fragSrc + "\x00")
	defer freeF()

	v := gl.CreateShader(gl.VERTEX_SHADER)
	gl.ShaderSource(v, 1, vsrc, nil)
	gl.CompileShader(v)
	var status int32
	gl.GetShaderiv(v, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLen int32
		gl.GetShaderiv(v, gl.INFO_LOG_LENGTH, &logLen)
		if logLen > 0 {
			log := make([]byte, logLen+1)
			gl.GetShaderInfoLog(v, logLen, nil, &log[0])
			core.Debugf("vertex shader log: %s\n", string(log))
		}
		return 0, fmt.Errorf("vertex shader compile failed")
	}

	f := gl.CreateShader(gl.FRAGMENT_SHADER)
	gl.ShaderSource(f, 1, fsrc, nil)
	gl.CompileShader(f)
	gl.GetShaderiv(f, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLen int32
		gl.GetShaderiv(f, gl.INFO_LOG_LENGTH, &logLen)
		if logLen > 0 {
			log := make([]byte, logLen+1)
			gl.GetShaderInfoLog(f, logLen, nil, &log[0])
			core.Debugf("fragment shader log: %s\n", string(log))
		}
		return 0, fmt.Errorf("fragment shader compile failed")
	}

	p := gl.CreateProgram()
	gl.AttachShader(p, v)
	gl.AttachShader(p, f)
	gl.LinkProgram(p)
	gl.GetProgramiv(p, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		var logLen int32
		gl.GetProgramiv(p, gl.INFO_LOG_LENGTH, &logLen)
		if logLen > 0 {
			log := make([]byte, logLen+1)
			gl.GetProgramInfoLog(p, logLen, nil, &log[0])
			core.Debugf("program link log: %s\n", string(log))
		}
		return 0, fmt.Errorf("program link failed")
	}
	return p, nil
}
