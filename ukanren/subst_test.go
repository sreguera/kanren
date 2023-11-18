package ukanren

import (
	"testing"
)

func TestFind(t *testing.T) {
	var s Subst = emptySubst().extend(&Var{"y"}, &Str{"ys"}).extend(&Var{"x"}, &Var{"y"})
	if s.find(&Var{"y"}) == nil {
		t.Fatal("could not find variable in subst")
	}
}
