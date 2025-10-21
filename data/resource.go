package data

import "fmt"

// Resource represents an external resource (file, URL). Loading implementations
// are intentionally minimal placeholders for the Go prototype.
type Resource struct {
	ResourcePath string
	LoadSync     bool
	Initialized  bool
	callbacks    []func()
}

func NewResource(path string, loadSync bool) *Resource {
	return &Resource{ResourcePath: path, LoadSync: loadSync}
}

func (r *Resource) String() string { return fmt.Sprintf("Resource(file=%s)", r.ResourcePath) }

func (r *Resource) OnInit(fn func()) { r.callbacks = append(r.callbacks, fn) }

func (r *Resource) ReportInit() {
	r.Initialized = true
	for _, cb := range r.callbacks {
		cb()
	}
	r.callbacks = nil
}
