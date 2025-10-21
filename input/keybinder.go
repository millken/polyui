package input

import (
	"sync"

	"github.com/millken/polyui/event"
)

// KeyBinder is a lightweight port of the Kotlin KeyBinder. It tracks down
// keys and forwards binds. This implementation is intentionally minimal and
// focuses on compatibility for the prototype.
type KeyBinder struct {
	settings                  interface{} // placeholder for Settings; keep as interface{} to avoid cycle
	listeners                 []*PolyBind
	mu                        sync.Mutex
	hasTimeSensitiveListeners bool
	downMouse                 IntSet
	downUnmapped              IntSet
	downKeys                  []Key
	recording                 *PolyBind
	recordingCallback         func(*PolyBind)
}

func NewKeyBinder(settings interface{}) *KeyBinder {
	return &KeyBinder{settings: settings, downMouse: NewIntSet(), downUnmapped: NewIntSet()}
}

func (kb *KeyBinder) AcceptEvent(e event.Event, mods byte) bool {
	// Note: a complete translation would assert on concrete event types.
	// For now, we return false to let callers handle forwarding.
	return false
}

func (kb *KeyBinder) AcceptUnmapped(key int, down bool, mods byte) bool {
	if down {
		kb.downUnmapped.Add(key)
		kb.completeRecording(mods)
		kb.Update(0, mods, true)
	} else {
		kb.downUnmapped.Remove(key)
		kb.Update(0, mods, false)
	}
	return false
}

func (kb *KeyBinder) Update(deltaNanos int64, mods byte, down bool) bool {
	if !kb.hasTimeSensitiveListeners && deltaNanos > 0 {
		return false
	}
	for _, b := range kb.listeners {
		if b.Update(kb.downUnmapped, kb.downKeys, kb.downMouse, mods, deltaNanos, down) {
			return true
		}
	}
	return false
}

func (kb *KeyBinder) Add(b *PolyBind) { kb.listeners = append(kb.listeners, b) }
func (kb *KeyBinder) Remove(b *PolyBind) {
	out := kb.listeners[:0]
	for _, x := range kb.listeners {
		if x != b {
			out = append(out, x)
		}
	}
	kb.listeners = out
}

func (kb *KeyBinder) Record(b *PolyBind, cb func(*PolyBind)) {
	kb.recording = b
	kb.recordingCallback = cb
}

func (kb *KeyBinder) CancelRecord(reason string) {
	kb.recording = nil
	kb.recordingCallback = nil
}

func (kb *KeyBinder) Release() {
	kb.downMouse.Clear()
	kb.downUnmapped.Clear()
	kb.downKeys = kb.downKeys[:0]
	if kb.recording != nil {
		cb := kb.recordingCallback
		b := kb.recording
		kb.recording = nil
		kb.recordingCallback = nil
		if cb != nil {
			cb(b)
		}
	}
}

func (kb *KeyBinder) completeRecording(mods byte) {
	// minimal no-op: finish recording if set
	if kb.recording == nil {
		return
	}
	b := kb.recording
	if kb.recordingCallback != nil {
		kb.recordingCallback(b)
	}
	kb.recording = nil
	kb.recordingCallback = nil
}
