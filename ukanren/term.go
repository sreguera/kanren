package ukanren

import "fmt"

type Term interface {
	Equal(o Term) bool
}

type Var struct {
	Id int
}

type Pair struct {
	Car Term
	Cdr Term
}

type Str struct {
	Value string
}

type Int struct {
	Value int
}

type Nil struct {
}

func (t *Var) Equal(o Term) bool {
	v, ok := o.(*Var)
	return ok && t.Id == v.Id
}

func (t *Pair) Equal(o Term) bool {
	p, ok := o.(*Pair)
	return ok && t.Car.Equal(p.Car) && t.Cdr.Equal(p.Cdr)
}

func (t *Str) Equal(o Term) bool {
	s, ok := o.(*Str)
	return ok && t.Value == s.Value
}

func (t *Int) Equal(o Term) bool {
	s, ok := o.(*Int)
	return ok && t.Value == s.Value
}

func (t *Nil) Equal(o Term) bool {
	_, ok := o.(*Nil)
	return ok
}

func (t *Var) String() string {
	return fmt.Sprintf("v.%d", t.Id)
}

func (t *Pair) String() string {
	return fmt.Sprintf("p.(%v, %v)", t.Car, t.Cdr)
}

func (t *Str) String() string {
	return fmt.Sprintf("s.%s", t.Value)
}

func (t *Int) String() string {
	return fmt.Sprintf("i.%d", t.Value)
}

func (t *Nil) String() string {
	return "nil"
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
