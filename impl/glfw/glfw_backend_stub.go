//go:build !glfw
// +build !glfw

package glfw

// This is a stub file that exists to document the planned GLFW+NanoVG backend.
// The real implementation should be placed in files built with the 'glfw' build tag
// and will require CGO and external Go OpenGL/GLFW bindings.

// To build the real backend, add files with `//go:build glfw` and ensure the
// appropriate dependencies and system headers are installed.
