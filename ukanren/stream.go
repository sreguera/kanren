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
	for it := s; it != nil; {
		items = append(items, fmt.Sprintf("%v", it.v))
		it1, ok := it.next.(*stream)
		if !ok {
			items = append(items, "...")
		}
		it = it1
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

type ztream struct {
	f func() Stream
}

func Zzz(g func() Stream) Stream {
	return &ztream{g}
}

func (s *ztream) head() *State {
	panic(nil)
}

func (s *ztream) tail() Stream {
	panic(nil)
}

func (s *ztream) mplus(t Stream) Stream {
	x := func() Stream {
		return s.f().mplus(t)
	}
	return &ztream{x}
}

func (s *ztream) bind(g Goal) Stream {
	x := func() Stream {
		return s.f().bind(g)
	}
	return &ztream{x}
}

func (s *ztream) String() string {
	return "..."
}
