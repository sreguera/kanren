package main

import (
	"fmt"
	uk "sreguera/kanren/ukanren"
)

func main() {
	var v uk.Term = &uk.Var{Name: "x"}
	fmt.Println(v.IsPair())

	var z uk.Term = uk.Cons(&uk.Var{Name: "x"}, &uk.Var{Name: "y"})
	fmt.Println(z.IsPair())

	r := uk.CallFresh(func(v *uk.Var) uk.Goal {
		return uk.Equiv(v, &uk.Str{"peich"})
	})(uk.EmptyState())
	fmt.Println(r)
}
