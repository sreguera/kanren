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
