package ukanren

import (
	"testing"
)

func TestFind(t *testing.T) {
	var s Subst = emptySubst().extend(&Var{1}, &Str{"ys"}).extend(&Var{0}, &Var{1})
	if s.find(&Var{1}) == nil {
		t.Fatal("could not find variable in subst")
	}
}
