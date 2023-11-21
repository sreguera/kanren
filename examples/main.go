package main

import (
	"fmt"

	uk "github.com/sreguera/kanren/ukanren"
)

// Paper, section 4.1, see below
func fives(x *uk.Var) uk.Goal {
	return uk.Disj(
		uk.Equiv(x, &uk.Int{5}),
		func(s *uk.State) uk.Stream {
			return uk.Zzz(func() uk.Stream {
				return fives(x)(s)
			})
		})
}

func main() {

	// Paper, section 2, first example
	r1 := uk.CallFresh(func(v *uk.Var) uk.Goal {
		return uk.Equiv(v, &uk.Int{5})
	})(uk.EmptyState())
	fmt.Println("\nFirst example result:")
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
	fmt.Println("\nSecond example result:")
	fmt.Println(r2)

	// Paper, section 4.1, see fives above
	r3 := uk.CallFresh(fives)(uk.EmptyState())
	fmt.Println("\nThird example result:")
	fmt.Println(r3)

	fmt.Println()
}
