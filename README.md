# PolyUI Go prototype — Running the GLFW example

This supplemental README describes how to run the GLFW example and the native dependencies needed.

Quick start (Linux)

Prerequisites
- Go 1.20+ installed
- System OpenGL and GLFW development libraries installed. Example packages on Debian/Ubuntu:

```bash
sudo apt-get install libglfw3-dev libglu1-mesa-dev libgles2-mesa-dev libx11-dev
```

Run the GLFW example

From the repository root you can run the example with:

```bash
# run the GLFW-specific example file from the repo root
go run -tags glfw ./example/main_glfw.go
```

Notes
- The example uses a small OpenGL renderer (`impl/gl`) and a GLFW backend (`impl/glfw`).
- If you run in a headless environment (no X11/Wayland), the example may not display a window even though the render loop executes. Use a desktop session, X forwarding, or run on a host with a display.

Troubleshooting
- If `go build` or `go run` fails while fetching `go-gl` or `glfw` modules, run `go mod tidy` and ensure network access. If specific commit versions are unavailable in your environment, update `go.mod` to point to stable tag versions.
- If shaders fail to compile, `gl` renderer prints shader/program info logs to stdout to help debug compilation issues.

Module & Tests
- The Go module path used in the `go-impl` prototype is `github.com/millken/polyui` (package name `polyui`).

Run unit tests for the Go prototype from the repo root:

```bash
go test ./...
```

If you want me to fully replace the existing `README.md`, tell me and I'll try again (I kept the original file untouched to avoid accidental formatting changes).