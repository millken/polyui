package core

type KeyBinder struct {
	listeners                 []*PolyBind
	hasTimeSensitiveListeners bool
	downMouseButtons          IntSet
	downUnmappedKeys          IntSet
	downKeys                  []int

	recordingBind     *PolyBind
	recordingCallback func(*PolyBind)
}

func NewKeyBinder() *KeyBinder {
	return &KeyBinder{downMouseButtons: NewIntSet(), downUnmappedKeys: NewIntSet()}
}

// Accept handles mapped-key events represented as Event (KeyPressed/Released)
// Returns true if the event was consumed by a bind.
func (kb *KeyBinder) Accept(event Event, mods byte) bool {
	// We only care about key events (use Type to switch)
	switch event.Type() {
	case EventKindKeyPressed, EventKindKeyReleased:
		// TODO: convert Event -> key value; currently our Event implementations
		// in input.go are simple structs (KeyPressed/KeyReleased) so do a type assert.
		switch e := event.(type) {
		case KeyPressed:
			kb.downKeys = append(kb.downKeys, e.Key)
			if kb.recordingBind != nil {
				// add to recording Bind if absent
				found := false
				for _, k := range kb.recordingBind.Keys {
					if k == e.Key {
						found = true
						break
					}
				}
				if !found {
					kb.recordingBind.Keys = append(kb.recordingBind.Keys, e.Key)
				}
			}
		case KeyReleased:
			// remove from downKeys
			nk := make([]int, 0, len(kb.downKeys))
			for _, k := range kb.downKeys {
				if k != e.Key {
					nk = append(nk, k)
				}
			}
			kb.downKeys = nk
		}
	default:
		// other event kinds ignored by Accept
	}
	return false
}

// AcceptUnmapped handles unmapped key codes (raw scancodes etc.).
func (kb *KeyBinder) AcceptUnmapped(key int, down bool, mods byte) bool {
	if down {
		kb.downUnmappedKeys.Add(key)
		if kb.recordingBind != nil {
			// add if absent
			found := false
			for _, k := range kb.recordingBind.UnmappedKeys {
				if k == key {
					found = true
					break
				}
			}
			if !found {
				kb.recordingBind.UnmappedKeys = append(kb.recordingBind.UnmappedKeys, key)
			}
		}
	} else {
		kb.downUnmappedKeys.Remove(key)
	}
	return false
}

// Update runs the matching logic for all listeners. Returns true if any
// listener consumed the input and requested suppression.
func (kb *KeyBinder) Update(deltaNanos int64, mods byte, down bool) bool {
	consumed := false
	for _, b := range kb.listeners {
		// pass current down sets and keys
		if b.Update(kb.downUnmappedKeys, kb.downKeys, kb.downMouseButtons, Modifiers(mods), deltaNanos, down) {
			consumed = true
		}
	}
	return consumed
}

func (kb *KeyBinder) Add(bind *PolyBind) { kb.listeners = append(kb.listeners, bind) }

func (kb *KeyBinder) Remove(bind *PolyBind) {
	// remove by pointer equality
	out := kb.listeners[:0]
	for _, b := range kb.listeners {
		if b != bind {
			out = append(out, b)
		}
	}
	kb.listeners = out
}

// Record starts a recording session for a bind. The callback is invoked when
// recording finishes with the recorded bind.
func (kb *KeyBinder) Record(bind *PolyBind, callback func(*PolyBind)) {
	kb.recordingBind = bind
	kb.recordingCallback = callback
}

func (kb *KeyBinder) CancelRecord(reason string) {
	kb.recordingBind = nil
	kb.recordingCallback = nil
}

func (kb *KeyBinder) Release() {
	// clear state
	kb.downMouseButtons.Clear()
	kb.downUnmappedKeys.Clear()
	kb.downKeys = kb.downKeys[:0]
	// If we were recording, finish recording and call callback
	if kb.recordingBind != nil {
		cb := kb.recordingCallback
		b := kb.recordingBind
		kb.recordingBind = nil
		kb.recordingCallback = nil
		if cb != nil {
			cb(b)
		}
	}
}
