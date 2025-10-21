# PolyUI Go prototype (go-impl)

This folder contains a minimal Go prototype of the PolyUI core API. It is intentionally small and uses a NoOp renderer to validate the core contracts (Renderer, Window, Component, Input flow) without any native or cgo dependencies.

Run the example:

```bash
go run ./example
```

This will print simulated draw calls to stdout.
