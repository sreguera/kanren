package ukanren

import (
	"testing"
)

func TestWalk(t *testing.T) {
	var s Subst = emptySubst().extend(&Var{"y"}, &Str{"ys"}).extend(&Var{"x"}, &Var{"y"})
	if !walk(&Var{"x"}, s).Equal(&Str{"ys"}) {
		t.Fatal("could not find variable in walk")
	}
}

func TestWalkNotFound(t *testing.T) {
	var s Subst = emptySubst().extend(&Var{"y"}, &Str{"ys"}).extend(&Var{"x"}, &Var{"y"})
	if !walk(&Var{"z"}, s).Equal(&Var{"z"}) {
		t.Fatal("could not find variable in walk")
	}
}

func TestUnify(t *testing.T) {
	var t1 Term = &Pair{&Var{"x"}, &Pair{&Str{"y"}, &Nil{}}}
	var t2 Term = &Pair{&Str{"sx"}, &Var{"y"}}
	var s = unify(t1, t2, emptySubst())
	if !s.find(&Var{"x"}).Equal(&Str{"sx"}) {
		t.Fatal("can't find var x")
	}
	if !s.find(&Var{"y"}).Equal(&Pair{&Str{"y"}, &Nil{}}) {
		t.Fatal("can't find var y")
	}
}
