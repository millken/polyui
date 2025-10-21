GLFW+NanoVG backend notes

Dependencies (suggested):
- github.com/go-gl/glfw/v3.3/glfw
- github.com/go-gl/gl/v4.6/gl
- NanoVG binding: there is no canonical maintained NanoVG Go binding; options:
  - use a maintained fork such as https://github.com/fyne-io/nanovg
  - write minimal cgo wrappers for the required NanoVG API surface used by Renderer

System requirements:
- OpenGL development headers
- GLFW development headers (libglfw3-dev on Debian/Ubuntu)

Build (example):

```bash
CGO_ENABLED=1 go build -tags glfw ./go-impl/impl/glfw
```

Implementation tips:
- Mirror method set from Kotlin `NVGRenderer.kt` to Go `NVGRenderer` methods.
- Keep default font/image fallbacks in the core module; renderer should accept raw bytes or file paths.
