package main

import "fmt"

type Calc struct {
	val1, val2 int
}

type Sums []Calc

// Add is simple function
func Add(q Calc) int {
	return q.val1 + q.val2
}

// Add is Calc method
func (p Calc) Add() int {
	return p.val1 + p.val2
}

// Add is Calc Class method
func (r *Calc) setValue(val1, val2 int) {
	r.val1 = val1
	r.val2 = val2
}

func (s Sums) Adds() (answer int) {
	answer = 0
	for _, s := range s {
		answer += s.Add()
	}

	return
}

var calc1, calc2 Calc

func main() {
	calc1.setValue(10, 30)
	fmt.Println(Add(calc1))
	fmt.Println(calc1)

	calc2.setValue(20, 1)
	fmt.Println(calc2.Add())

	sums := Sums{
		Calc{2, 10},
		{10, 23},
		Calc{2, 1},
		Calc{22, 13},
		{10, 1},
	}
	fmt.Println(sums.Adds())
}
