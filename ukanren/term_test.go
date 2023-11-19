package ukanren

import (
	"testing"
)

func TestList(t *testing.T) {
	var v Term = List(&Str{"a"}, &Str{"b"})
	if !v.Equal(&Pair{&Str{"a"}, &Pair{&Str{"b"}, &Nil{}}}) {
		t.Fatal("bad list construction")
	}
}
