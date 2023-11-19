package ukanren

type Term interface {
	Equal(o Term) bool
}

type Var struct {
	Name string
}

type Pair struct {
	Car Term
	Cdr Term
}

type Str struct {
	Value string
}

type Nil struct {
}

func (t *Var) Equal(o Term) bool {
	v, ok := o.(*Var)
	return ok && t.Name == v.Name
}

func (t *Pair) Equal(o Term) bool {
	p, ok := o.(*Pair)
	return ok && t.Car.Equal(p.Car) && t.Cdr.Equal(p.Cdr)
}

func (t *Str) Equal(o Term) bool {
	s, ok := o.(*Str)
	return ok && t.Value == s.Value
}

func (t *Nil) Equal(o Term) bool {
	_, ok := o.(*Nil)
	return ok
}

func Cons(t1, t2 Term) Term {
	return &Pair{t1, t2}
}

func List(terms ...Term) Term {
	if terms == nil {
		return &Nil{}
	} else if len(terms) == 1 {
		return Cons(terms[0], &Nil{})
	} else {
		return Cons(terms[0], List(terms[1:]...))
	}
}
