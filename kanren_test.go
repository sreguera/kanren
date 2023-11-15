package main

import (
	"testing"
)

func TestVarIsPair(t *testing.T) {
	var v Thing = &Var{"x"}
	if v.isPair() {
		t.Fatal("var says it is a pair")
	}
}

func TestPairIsPair(t *testing.T) {
	var v Thing = &Pair{&Var{"x"}, &Var{"y"}}
	if !v.isPair() {
		t.Fatal("pair says it is not a pair")
	}
}

func TestFindSubst(t *testing.T) {
	var x Thing = &Pair{&Var{"x"}, &Str{"xs"}}
	var y Thing = &Pair{&Var{"y"}, &Str{"ys"}}
	var v Thing = &Pair{x, &Pair{y, &Nil{}}}
	if !findSubst(&Var{"y"}, v).equal(y) {
		t.Fatal("could not find variable in subst")
	}
}
