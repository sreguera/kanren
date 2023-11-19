package main

import (
	"fmt"
	uk "sreguera/kanren/ukanren"
)

var a_and_b uk.Goal = uk.Conj(
	uk.CallFresh(func(a *uk.Var) uk.Goal {
		return uk.Equiv(a, &uk.Str{"7"})
	}),
	uk.CallFresh(func(b *uk.Var) uk.Goal {
		return uk.Disj(uk.Equiv(b, &uk.Str{"5"}), uk.Equiv(b, &uk.Str{"6"}))
	}),
)

func main() {

	r := uk.CallFresh(func(v *uk.Var) uk.Goal {
		return uk.Equiv(v, &uk.Str{"peich"})
	})(uk.EmptyState())
	fmt.Println(r)

	r2 := a_and_b(uk.EmptyState())
	fmt.Println(r2)
}
