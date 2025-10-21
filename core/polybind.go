package core

type PolyBind struct {
	UnmappedKeys  []int
	Keys          []int
	Mouse         []int
	Mods          Modifiers
	DurationNanos int64
	Action        func(down bool) bool

	// internal
	Muted bool
	// internal runtime state
	elapsedNanos int64
	active       bool
}

func (p *PolyBind) Update(downUnmapped IntSet, downKeys []int, downMouse IntSet, mods Modifiers, deltaNanos int64, down bool) bool {
	if p.Muted || p.Action == nil {
		return false
	}

	// quick mods check: require all mods in p.Mods to be present
	if p.Mods != 0 && (mods&p.Mods) != p.Mods {
		// mods not satisfied
		p.elapsedNanos = 0
		p.active = false
		return false
	}

	// helper to check slice-of-int presence
	containsAll := func(have []int, want []int, set IntSet) bool {
		// if want is provided as a slice, check presence in have or set
		for _, w := range want {
			found := false
			if set != nil {
				if set.Has(w) {
					found = true
				}
			}
			if !found {
				for _, v := range have {
					if v == w {
						found = true
						break
					}
				}
			}
			if !found {
				return false
			}
		}
		return true
	}

	// match Keys (regular mapped keys) and UnmappedKeys and Mouse
	if !containsAll(downKeys, p.Keys, nil) {
		p.elapsedNanos = 0
		p.active = false
		return false
	}
	if !containsAll(nil, p.UnmappedKeys, downUnmapped) {
		p.elapsedNanos = 0
		p.active = false
		return false
	}
	if !containsAll(nil, p.Mouse, downMouse) {
		p.elapsedNanos = 0
		p.active = false
		return false
	}

	// if we reach here, the bind's required buttons/keys are all down
	if p.DurationNanos > 0 {
		p.elapsedNanos += deltaNanos
		if p.elapsedNanos < p.DurationNanos {
			// still accumulating
			p.active = true
			return false
		}
		// duration satisfied -> fire action
		p.elapsedNanos = 0
		p.active = false
		return p.Action(down)
	}

	// No duration required: invoke action immediately on matching.
	return p.Action(down)
}

func (p *PolyBind) ResetState() {}
