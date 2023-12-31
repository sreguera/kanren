package ukanren

import (
	"fmt"
	"strings"
)

type Subst interface {
	extend(v *Var, t Term) Subst
	find(v *Var) Term
}

type subst struct {
	v    *Var
	t    Term
	next *subst
}

func (s *subst) String() string {
	items := []string{}
	for it := s; it.next != nil; it = it.next {
		items = append(items, fmt.Sprintf("(%v, %v)", it.v, it.t))
	}
	return "[" + strings.Join(items, ", ") + "]"
}

func emptySubst() Subst {
	return &subst{}
}

func (s *subst) extend(v *Var, t Term) Subst {
	return &subst{v, t, s}
}

func (s *subst) find(v *Var) Term {
	for entry := s; entry.next != nil; entry = entry.next {
		if entry.v != nil && entry.v.Equal(v) {
			return entry.t
		}
	}
	return nil
}
