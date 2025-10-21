package event

import "sync"

// Generic State container similar to the Kotlin State<T>.
// Listeners have the signature func(T) bool and may return true to veto
// the change. For convenience, Listen returns a function to remove the
// listener.

type listenerEntry[T comparable] struct {
	id int
	fn func(T) bool
}

type State[T comparable] struct {
	mu                 sync.RWMutex
	v                  T
	listeners          []listenerEntry[T]
	nextID             int
	instanceChangeOnly func(T) bool
}

func NewState[T comparable](value T) *State[T] {
	return &State[T]{v: value}
}

func (s *State[T]) Value() T {
	s.mu.RLock()
	v := s.v
	s.mu.RUnlock()
	return v
}

func (s *State[T]) Set(value T) {
	s.mu.RLock()
	if s.v == value {
		s.mu.RUnlock()
		return
	}
	inst := s.instanceChangeOnly
	// copy listeners under read lock
	listeners := make([]listenerEntry[T], len(s.listeners))
	copy(listeners, s.listeners)
	s.mu.RUnlock()

	if inst != nil && inst(value) {
		return
	}
	for _, e := range listeners {
		if e.fn(value) {
			return
		}
	}

	s.mu.Lock()
	s.v = value
	s.mu.Unlock()
}

// Notify triggers listeners with the current value. Returns true if any
// listener vetoed the notification.
func (s *State[T]) Notify() bool {
	v := s.Value()
	s.mu.RLock()
	listeners := make([]listenerEntry[T], len(s.listeners))
	copy(listeners, s.listeners)
	s.mu.RUnlock()

	for _, e := range listeners {
		if e.fn(v) {
			return true
		}
	}
	return false
}

// Listen adds a listener and returns a removal function. The listener should
// return true to cancel the change.
func (s *State[T]) Listen(listener func(T) bool) func() {
	s.mu.Lock()
	s.nextID++
	id := s.nextID
	s.listeners = append(s.listeners, listenerEntry[T]{id: id, fn: listener})
	s.mu.Unlock()
	return func() { s.removeListener(id) }
}

func (s *State[T]) removeListener(id int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i, e := range s.listeners {
		if e.id == id {
			last := len(s.listeners) - 1
			s.listeners[i] = s.listeners[last]
			s.listeners = s.listeners[:last]
			return
		}
	}
}

// ListenSimple is a convenience for listeners that don't return a bool.
// It returns a removal function.
func (s *State[T]) ListenSimple(listener func(T)) func() {
	return s.Listen(func(t T) bool { listener(t); return false })
}

// ListenToInstanceChange sets a listener that only triggers when the instance
// (identity) of the value changes. Returns the previous listener.
func (s *State[T]) ListenToInstanceChange(listener func(T) bool) (old func(T) bool) {
	s.mu.Lock()
	old = s.instanceChangeOnly
	s.instanceChangeOnly = listener
	s.mu.Unlock()
	return old
}
