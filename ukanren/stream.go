package ukanren

import (
	"fmt"
	"strings"
)

type State struct {
	s Subst
	i int
}

func (s *State) String() string {
	return fmt.Sprintf("{%v, %d}", s.s, s.i)
}

func EmptyState() *State {
	return &State{emptySubst(), 0}
}

type Stream interface {
	head() *State
	tail() Stream
	mplus(Stream) Stream
	bind(Goal) Stream
}

type Goal func(*State) Stream

type stream struct {
	v    *State
	next Stream
}

func (s *stream) String() string {
	items := []string{}
	for it := s; it != nil; it = it.next.(*stream) {
		items = append(items, fmt.Sprintf("%v", it.v))
	}
	return "<" + strings.Join(items, ",\n ") + ">"
}

func (s *stream) head() *State {
	return s.v
}

func (s *stream) tail() Stream {
	return s.next
}

func (s *stream) mplus(t Stream) Stream {
	if s == nil {
		return t
	} else {
		return &stream{s.head(), s.tail().mplus(t)}
	}
}

func (s *stream) bind(g Goal) Stream {
	if s == nil {
		return mzero()
	} else {
		return g(s.head()).mplus(s.tail().bind(g))
	}
}

func unit(state *State) *stream {
	return &stream{state, mzero()}
}

func mzero() *stream {
	return nil
}
