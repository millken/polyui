package input

// PolyBind is a minimal port of Kotlin PolyBind used for keybinds.
type PolyBind struct {
	UnmappedKeys  []int
	Keys          []Key
	Mouse         []int
	Mods          Modifiers
	DurationNanos int64
	Action        func(down bool) bool

	// runtime state
	Muted        bool
	elapsedNanos int64
	ran          bool
}

func (p *PolyBind) Update(downUnmapped IntSet, downKeys []Key, downMouse IntSet, mods byte, deltaNanos int64, down bool) bool {
	if p.Muted || p.Action == nil {
		return false
	}
	// minimal test: if not bound, return false
	if !p.IsBound() {
		return false
	}
	if p.DurationNanos == 0 && deltaNanos > 0 {
		return false
	}
	// check unmapped keys
	for _, k := range p.UnmappedKeys {
		if !downUnmapped.Has(k) {
			return false
		}
	}
	// check keys
	for _, k := range p.Keys {
		found := false
		for _, dk := range downKeys {
			if dk == k {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	// check mouse
	for _, m := range p.Mouse {
		if !downMouse.Has(m) {
			return false
		}
	}
	// basic mods check omitted for brevity
	if p.DurationNanos != 0 {
		p.elapsedNanos += deltaNanos
		if p.elapsedNanos >= p.DurationNanos && !p.ran {
			p.ran = true
			return p.Action(true)
		}
	} else if down && !p.ran {
		p.ran = true
		return p.Action(true)
	}
	return false
}

func (p *PolyBind) ResetState() {
	if p.ran && p.Action != nil {
		p.Action(false)
	}
	p.elapsedNanos = 0
	p.ran = false
}

func (p *PolyBind) IsBound() bool {
	return len(p.UnmappedKeys) > 0 || len(p.Keys) > 0 || len(p.Mouse) > 0 || p.Mods != 0
}
