package ukanren

type State struct {
	s Subst
	i int
}

func EmptyState() *State {
	return &State{emptySubst(), 0}
}

type Stream interface {
	head() *State
	tail() Stream
}

type stream struct {
	v    *State
	next *stream
}

func (s *stream) head() *State {
	return s.v
}

func (s *stream) tail() Stream {
	return s.next
}

func unit(state *State) *stream {
	return &stream{state, mzero()}
}

func mzero() *stream {
	return nil
}
