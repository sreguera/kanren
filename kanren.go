package main

import "fmt"

type Thing interface {
	isVar() bool
	isPair() bool
	isNil() bool
	equal(o Thing) bool
}

type Var struct {
	name string
}

type Pair struct {
	car Thing
	cdr Thing
}

type Str struct {
	value string
}

type Nil struct {
}

func (t *Var) isVar() bool  { return true }
func (t *Pair) isVar() bool { return false }
func (t *Str) isVar() bool  { return false }
func (t *Nil) isVar() bool  { return false }

func (t *Var) isPair() bool  { return false }
func (t *Pair) isPair() bool { return true }
func (t *Str) isPair() bool  { return false }
func (t *Nil) isPair() bool  { return false }

func (t *Var) isNil() bool  { return false }
func (t *Pair) isNil() bool { return false }
func (t *Str) isNil() bool  { return false }
func (t *Nil) isNil() bool  { return true }

func (t *Var) equal(o Thing) bool {
	v, ok := o.(*Var)
	return ok && t.name == v.name
}

func (t *Pair) equal(o Thing) bool {
	p, ok := o.(*Pair)
	return ok && t.car.equal(p.car) && t.cdr.equal(p.cdr)
}

func (t *Str) equal(o Thing) bool {
	s, ok := o.(*Str)
	return ok && t.value == s.value
}

func (t *Nil) equal(o Thing) bool {
	_, ok := o.(*Nil)
	return ok
}

func walk(u, s Thing) Thing {
	if !u.isVar() {
		return u
	} else {
	}
	return u
}

func findSubst(x, s Thing) Thing {
	v := x.(*Var)
	p, ok := s.(*Pair)
	for ok {
		ass, ok1 := p.car.(*Pair)
		if ok1 && v.equal(ass.car) {
			return ass
		}
		p, ok = p.cdr.(*Pair)
	}
	return &Nil{}
}

func extendSubst(x, v, s Thing) Thing {
	return &Pair{&Pair{x, v}, s}
}

func main() {
	var v Thing = &Var{"x"}
	fmt.Println(v.isPair())

	var z Thing = &Pair{&Var{"x"}, &Var{"y"}}
	fmt.Println(z.isPair())

}
