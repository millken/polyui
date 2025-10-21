package core

// PolyUI holds the root component and settings for the prototype.
type PolyUI struct{
    Root Component
    Renderer Renderer
    Window Window
}

func NewPolyUI(root Component, r Renderer) *PolyUI{
    return &PolyUI{ Root: root, Renderer: r }
}

func (p *PolyUI) Init() error{
    if err := p.Renderer.Init(); err != nil { return err }
    p.Root.Setup(p)
    return nil
}

func (p *PolyUI) RenderFrame(width, height, pixelRatio float32){
    p.Renderer.BeginFrame(width,height,pixelRatio)
    p.Root.PreRender(p.Renderer)
    p.Root.Render(p.Renderer)
    p.Root.PostRender(p.Renderer)
    p.Renderer.EndFrame()
}
