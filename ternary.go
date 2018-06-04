// Implementation of a ternary counter. Any integer can be made by addition and
// subtraction of powers of 3. So 21 = 27 - 9 + 3.
//
// Test this by generating loads of them

package main

import (
	"fmt"
	"time"
)

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
	c.maximum = 0
	v := 1
	for i := 0; i < n; i++ {
		c.powers[i] = v
		c.maximum += v
		v *= 3
	}
	return c
}

func (c TernaryCounter) String() string {
	var s string
	s = fmt.Sprintf("Order: %d. Total: %d. Counter: %d. Maximum: %d",
		c.order, c.total, c.counter, c.maximum)
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

func (c *TernaryCounter) Increment(n int) {

	c.states[n]++
	if c.states[n] > 1 && n < c.order - 1 {
		c.states[n] = -1
		c.Increment(n + 1)
	}

	if n == 0 {
		c.counter++
		c.total = 0
		for i := 0 ; i < c.order ; i++ {
			c.total += c.powers[i] * c.states[i]
		}
	}

}

func main() {

	start := time.Now()
	c := NewTernaryCounter(25)
	for i := 0 ; i <= c.maximum ; i++ {
		c.Increment(0)
		if c.counter != c.total {
			fmt.Printf("Error. Bad count\n%v", c)
		}
	}
	elapsed := time.Since(start)
	fmt.Printf("%d operations in %v seconds (%d per second)",
		c.maximum, elapsed, int(float64(c.maximum)/elapsed.Seconds()))
}
