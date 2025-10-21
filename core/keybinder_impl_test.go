package core

import "testing"

func TestPolyBindImmediateAction_Impl(t *testing.T) {
	kb := NewKeyBinder()
	fired := false
	bind := &PolyBind{
		Keys:   []int{42},
		Action: func(down bool) bool { fired = true; return true },
	}
	kb.Add(bind)
	// simulate pressing key 42
	kb.downKeys = append(kb.downKeys, 42)
	if !kb.Update(0, 0, true) {
		t.Fatalf("expected Update to return true when bind action fired")
	}
	if !fired {
		t.Fatalf("expected bind action to have been executed")
	}
}

func TestAcceptUnmapped_Impl(t *testing.T) {
	kb := NewKeyBinder()
	kb.AcceptUnmapped(1001, true, 0)
	if !kb.downUnmappedKeys.Has(1001) {
		t.Fatalf("expected unmapped key 1001 to be present")
	}
	kb.AcceptUnmapped(1001, false, 0)
	if kb.downUnmappedKeys.Has(1001) {
		t.Fatalf("expected unmapped key 1001 to be removed")
	}
}
