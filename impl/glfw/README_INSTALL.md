Local install notes (Ubuntu/Debian)

Install system dependencies required for building GLFW/OpenGL bindings:

```bash
sudo apt update
sudo apt install -y build-essential libgl1-mesa-dev libx11-dev libxcursor-dev libxrandr-dev libxi-dev libglfw3-dev
```

For macOS, install GLFW via Homebrew:

```bash
brew install glfw
```

Then build with CGO enabled and the 'glfw' tag:

```bash
CGO_ENABLED=1 go run -tags glfw ./go-impl/example
```
