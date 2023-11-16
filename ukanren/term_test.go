package ukanren

import (
	"testing"
)

func TestVarIsPair(t *testing.T) {
	var v Term = &Var{"x"}
	if v.IsPair() {
		t.Fatal("var says it is a pair")
	}
}

func TestPairIsPair(t *testing.T) {
	var v Term = &Pair{&Var{"x"}, &Var{"y"}}
	if !v.IsPair() {
		t.Fatal("pair says it is not a pair")
	}
}
