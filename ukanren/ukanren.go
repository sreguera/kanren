package ukanren

func walk(u Term, s Subst) Term {
	v, isVar := u.(*Var)
	if !isVar {
		return u
	}
	r := s.find(v)
	if r == nil {
		return u
	}
	return walk(r, s)
}

func unify(u, v Term, s Subst) Subst {
	uw := walk(u, s)
	vw := walk(v, s)

	uv, uIsVar := uw.(*Var)
	vv, vIsVar := vw.(*Var)
	if uIsVar && vIsVar && uw.Equal(vw) {
		return s
	} else if uIsVar {
		return s.extend(uv, vw)
	} else if vIsVar {
		return s.extend(vv, uw)
	}

	up, uIsPair := uw.(*Pair)
	vp, vIsPair := vw.(*Pair)
	if uIsPair && vIsPair {
		s1 := unify(up.Car, vp.Car, s)
		if s1 == nil {
			return nil
		}
		s2 := unify(up.Cdr, vp.Cdr, s1)
		if s2 == nil {
			return nil
		}
		return s2
	}

	if uw.Equal(vw) {
		return s
	} else {
		return nil
	}
}

func Equiv(u, v Term) Goal {
	return func(s *State) Stream {
		rs := unify(u, v, s.s)
		if rs != nil {
			return unit(&State{rs, s.i})
		} else {
			return mzero()
		}
	}
}

func CallFresh(f func(*Var) Goal) Goal {
	return func(s *State) Stream {
		return f(&Var{s.i})(&State{s.s, s.i + 1})
	}
}

func Disj(g1, g2 Goal) Goal {
	return func(s *State) Stream {
		return g1(s).mplus(g2(s))
	}
}

func Conj(g1, g2 Goal) Goal {
	return func(s *State) Stream {
		return g1(s).bind(g2)
	}
}
