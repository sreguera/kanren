package ukanren

type Term interface {
	IsVar() bool
	IsPair() bool
	IsNil() bool
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

func (t *Var) IsVar() bool  { return true }
func (t *Pair) IsVar() bool { return false }
func (t *Str) IsVar() bool  { return false }
func (t *Nil) IsVar() bool  { return false }

func (t *Var) IsPair() bool  { return false }
func (t *Pair) IsPair() bool { return true }
func (t *Str) IsPair() bool  { return false }
func (t *Nil) IsPair() bool  { return false }

func (t *Var) IsNil() bool  { return false }
func (t *Pair) IsNil() bool { return false }
func (t *Str) IsNil() bool  { return false }
func (t *Nil) IsNil() bool  { return true }

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
