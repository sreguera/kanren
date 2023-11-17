package ukanren

func walk(u, s Term) Term {
	if !u.IsVar() {
		return u
	} else {
		r := findSubst(u, s)
		if r.IsNil() {
			return u
		} else {
			rr := r.(*Pair)
			return walk(rr.Cdr, s)
		}
	}
}

func findSubst(x, s Term) Term {
	v := x.(*Var)
	p, ok := s.(*Pair)
	for ok {
		ass, ok1 := p.Car.(*Pair)
		if ok1 && v.Equal(ass.Car) {
			return ass
		}
		p, ok = p.Cdr.(*Pair)
	}
	return &Nil{}
}

func extendSubst(x, v, s Term) Term {
	return &Pair{&Pair{x, v}, s}
}
