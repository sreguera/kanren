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
