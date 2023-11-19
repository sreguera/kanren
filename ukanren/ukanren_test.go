package ukanren

import (
	"testing"
)

func TestWalk(t *testing.T) {
	var s Subst = emptySubst().extend(&Var{1}, &Str{"ys"}).extend(&Var{0}, &Var{1})
	if !walk(&Var{0}, s).Equal(&Str{"ys"}) {
		t.Fatal("could not find variable in walk")
	}
}

func TestWalkNotFound(t *testing.T) {
	var s Subst = emptySubst().extend(&Var{1}, &Str{"ys"}).extend(&Var{0}, &Var{1})
	if !walk(&Var{2}, s).Equal(&Var{2}) {
		t.Fatal("could not find variable in walk")
	}
}

func TestUnify(t *testing.T) {
	var t1 Term = List(&Var{0}, &Str{"y"})
	var t2 Term = Cons(&Str{"sx"}, &Var{1})
	var s = unify(t1, t2, emptySubst())
	if !s.find(&Var{0}).Equal(&Str{"sx"}) {
		t.Fatal("can't find var x")
	}
	if !s.find(&Var{1}).Equal(List(&Str{"y"})) {
		t.Fatal("can't find var y")
	}
}

func TestCallFresh(t *testing.T) {
	r := CallFresh(func(v *Var) Goal {
		return Equiv(v, &Int{5})
	})(EmptyState())
	if !r.head().s.find(&Var{0}).Equal(&Int{5}) {
		t.Fatal("call fresh didn't work")
	}
}
