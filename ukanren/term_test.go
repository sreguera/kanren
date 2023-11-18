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
	var v Term = Cons(&Var{"x"}, &Var{"y"})
	if !v.IsPair() {
		t.Fatal("pair says it is not a pair")
	}
}

func TestList(t *testing.T) {
	var v Term = List(&Str{"a"}, &Str{"b"})
	if !v.Equal(&Pair{&Str{"a"}, &Pair{&Str{"b"}, &Nil{}}}) {
		t.Fatal("bad list construction")
	}
}
