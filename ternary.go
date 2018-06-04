// Implementation of a ternary counter. Any integer can be made by addition and
// subtraction of powers of 3. So 21 = 27 - 9 + 3.
//
// Test this by generating loads of them

package main

import "fmt"

type TernaryCounter struct {
	states, powers []int
	maximum, order, total, counter int
}

func NewTernaryCounter(n int) *TernaryCounter {
	// Constructor
	c := new(TernaryCounter)
	c.order = n
	c.states = make([]int, n)
	c.powers =  make([]int, n)
	c.total = 0
	c.counter = 0
	v := 1
	for i := 0; i < n; i++ {
		c.powers[i] = v
		v *= 3
	}
	return c
}

func (c TernaryCounter) String() string {
	var s string
	s = fmt.Sprintf("Order: %d. Total: %d. Counter: %d", c.order, c.total, c.counter)
	s += "\nStates: "
	for _, t := range c.states {
		s += fmt.Sprintf("\t%d", t)
	}
	s += "\nPowers: "
	for _, p := range c.powers {
		s += fmt.Sprintf("\t%d", p)
	}

	return s
}

func main() {
	c := NewTernaryCounter(5)
	fmt.Println(c)
}
