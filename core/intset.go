package core

// IntSet is a tiny int set used for down keys/mouse buttons.
type IntSet map[int]struct{}

func NewIntSet() IntSet         { return make(IntSet) }
func (s IntSet) Add(v int)      { s[v] = struct{}{} }
func (s IntSet) Remove(v int)   { delete(s, v) }
func (s IntSet) Has(v int) bool { _, ok := s[v]; return ok }
func (s IntSet) Clear() {
	for k := range s {
		delete(s, k)
	}
}
func (s IntSet) IsEmpty() bool { return len(s) == 0 }
