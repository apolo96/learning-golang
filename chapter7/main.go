package chapter7

import (
	"fmt"
	"time"
)

type Score int
type Convert func(string) Score
type TeamScores map[string]Score

type Counter struct {
	total      int
	lastUpdate time.Time
}

func (c *Counter) Increment() {
	c.total++
	c.lastUpdate = time.Now()
}

func (c Counter) String() string {
	return fmt.Sprintf("total: %d, last update: %v", c.total, c.lastUpdate)
}

type IntTree struct {
	val         int
	left, right *IntTree
}

func (it *IntTree) Insert(val int) *IntTree {
	if it == nil {
		return &IntTree{val: val}
	}
	if val < it.val {
		it.left = it.left.Insert(val)
	} else if val > it.val {
		it.right = it.right.Insert(val)
	}
	return it
}

func (it *IntTree) Contains(val int) bool {
	switch {
	case it == nil:
		return false
	case val < it.val:
		return it.left.Contains(val)
	case val > it.val:
		return it.right.Contains(val)
	default:
		return true
	}
}

type Adder struct {
	start int
}

func (a Adder) AddTo(val int) int {
	return a.start + val
}

func New() {
	fmt.Println("========== Chapter 7")
	fmt.Println("-------- Working with Structs and methods")
	var c Counter           // c := &Counter{}
	fmt.Println(c.String()) // (*c).String()
	c.Increment()           // (&c).Increment()
	fmt.Println(c.String())
	doUpdateWrong(c)
	fmt.Println("in main:", c.String())
	doUpdateRight(&c)
	fmt.Println("in main:", c.String())
	var it *IntTree
	fmt.Println(it)
	it = it.Insert(5)
	fmt.Println(it)
	it = it.Insert(3)
	it = it.Insert(10)
	it = it.Insert(2)
	fmt.Println(it.Contains(2))  // true
	fmt.Println(it.Contains(12)) // false
	fmt.Println("-------- Methods as functions")
	myAdder := Adder{start: 10}
	fmt.Println(myAdder.AddTo(5))
	f1 := myAdder.AddTo
	fmt.Println(f1(10))
	f2 := Adder.AddTo // func(Adder, int) int
	fmt.Println(f2(myAdder, 15))

}

func doUpdateWrong(c Counter) {
	c.Increment()
	fmt.Println("in doUpdateWrong ", c.String())
}

func doUpdateRight(c *Counter) {
	c.Increment()
	fmt.Println("in doUpdateRight:", c.String())
}
