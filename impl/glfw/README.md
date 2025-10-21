# GLFW + NanoVG backend (Go)

This folder will contain a GLFW + NanoVG renderer implementation for the Go prototype. It requires cgo and native bindings. Recommended libraries:

no maintained NanoVG Go binding was reliable in this environment. Prefer implementing a custom OpenGL
renderer (see ../gl) or using a maintained high-level Go UI library.
Notes:
- Building this backend requires system OpenGL/GLFW headers and CGO enabled. Prefer testing on Linux or macOS with proper dev packages installed.
Implementation plan:
1. Implement a `GLFWWindow` that satisfies `core.Window`, creates an OpenGL context and calls into a NanoVG wrapper.
2. Implement a `NVGRenderer` that fulfills `core.Renderer` by forwarding calls to NanoVG.
3. Add font/image loading fallbacks to match `Settings.resourcePolicy` semantics.
