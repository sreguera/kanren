package main

import (
	"fmt"

	uk "github.com/sreguera/kanren/ukanren"
)

func main() {

	// Paper, section 2, first example
	r1 := uk.CallFresh(func(v *uk.Var) uk.Goal {
		return uk.Equiv(v, &uk.Int{5})
	})(uk.EmptyState())
	fmt.Println(r1)

	// Paper, section 2, second example
	var a_and_b uk.Goal = uk.Conj(
		uk.CallFresh(func(a *uk.Var) uk.Goal {
			return uk.Equiv(a, &uk.Int{7})
		}),
		uk.CallFresh(func(b *uk.Var) uk.Goal {
			return uk.Disj(uk.Equiv(b, &uk.Int{5}), uk.Equiv(b, &uk.Int{6}))
		}),
	)

	r2 := a_and_b(uk.EmptyState())
	fmt.Println(r2)
}
