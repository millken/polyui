package core

import "testing"

func TestRecordBindAccumulation(t *testing.T) {
	kb := NewKeyBinder()
	bind := &PolyBind{}
	collected := false
	kb.Record(bind, func(b *PolyBind) { collected = true })
	// simulate unmapped key press
	kb.AcceptUnmapped(2002, true, 0)
	// simulate mapped key press
	kb.Accept(KeyPressed{Key: 77, Mods: 0}, 0)
	// finish recording
	kb.Release()
	if !collected {
		t.Fatalf("expected recording callback to be invoked on Release")
	}
	// check that bind contains the recorded keys
	foundUnmapped := false
	for _, k := range bind.UnmappedKeys {
		if k == 2002 {
			foundUnmapped = true
		}
	}
	if !foundUnmapped {
		t.Fatalf("expected UnmappedKeys to include 2002")
	}
	foundMapped := false
	for _, k := range bind.Keys {
		if k == 77 {
			foundMapped = true
		}
	}
	if !foundMapped {
		t.Fatalf("expected Keys to include 77")
	}
}
