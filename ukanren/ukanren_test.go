package ukanren

import (
	"testing"
)

func TestFindSubst(t *testing.T) {
	var x Term = &Pair{&Var{"x"}, &Var{"y"}}
	var y Term = &Pair{&Var{"y"}, &Str{"ys"}}
	var v Term = &Pair{x, &Pair{y, &Nil{}}}
	if !findSubst(&Var{"y"}, v).Equal(y) {
		t.Fatal("could not find variable in subst")
	}
}

func TestWalk(t *testing.T) {
	var x Term = &Pair{&Var{"x"}, &Var{"y"}}
	var y Term = &Pair{&Var{"y"}, &Str{"ys"}}
	var v Term = &Pair{x, &Pair{y, &Nil{}}}
	if !walk(&Var{"x"}, v).Equal(&Str{"ys"}) {
		t.Fatal("could not find variable in walk")
	}
}
