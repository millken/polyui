package input

// KeybindHelper is a minimal builder for PolyBind objects.
type KeybindHelper struct {
	duration int64
	keys     []Key
	mods     Modifiers
	unmapped IntSet
	mouse    IntSet
	fn       func(bool) bool
}

func NewKeybindHelper() *KeybindHelper {
	return &KeybindHelper{unmapped: NewIntSet(), mouse: NewIntSet()}
}

func (k *KeybindHelper) Build() *PolyBind {
	if k.fn == nil {
		panic("function must be set")
	}
	b := &PolyBind{UnmappedKeys: nil, Keys: k.keys, Mouse: nil, Mods: k.mods, DurationNanos: k.duration, Action: k.fn}
	return b
}

func (k *KeybindHelper) Keys(keys ...Key) *KeybindHelper        { k.keys = append(k.keys, keys...); return k }
func (k *KeybindHelper) ModsByte(b byte) *KeybindHelper         { k.mods = Modifiers(b); return k }
func (k *KeybindHelper) Does(fn func(bool) bool) *KeybindHelper { k.fn = fn; return k }
func (k *KeybindHelper) Duration(d int64) *KeybindHelper        { k.duration = d; return k }
