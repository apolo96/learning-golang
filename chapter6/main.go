package chapter6

import "fmt"

type Money struct {
	amount float64
}

type Person struct {
	firstName string
	lastName  string
	age       int
}

func newPerson(firstName, lastName string, age int) Person {
	return Person{
		firstName: firstName,
		lastName:  lastName,
		age:       age,
	}
}

func newPersonP(firstName, lastName string, age int) *Person {
	return &Person{
		firstName: firstName,
		lastName:  lastName,
		age:       age,
	}
}

func NewMoney(amount float64) Money {
	return Money{amount: amount}
}

func (m Money) Cantidad() float64 {
	return m.amount
}

func New() {
	/* A pointer is a variable that contains the memory address
	where a value is stored */
	// In Golang, all values passed to a function are copies. //
	/* Passing a pointer to a function means you passed a copy
	of the pointer (Memory address of value) to the function. */
	var x int = 10
	var y bool = true
	pointerX := &x
	pointerY := &y
	var pointerZ *string
	fmt.Println(pointerX, pointerY, pointerZ)
	fmt.Println(*pointerX, *pointerY)
	var xx = new(int)
	fmt.Println(xx, *xx)
	xy := &Money{}
	fmt.Println(xy)
	fmt.Printf("%p \n", xy)
	update(pointerX)
	fmt.Println(*pointerX)
	failedUpdate(pointerX)
	fmt.Println(*pointerX)
	/* In Golang, maps are implemented as pointers. */
	m := map[string]string{}
	addStudentToMap(m, "Andrew")
	fmt.Println(m)
	/* Anyone modified to the content of the slice is reflected
	in the original variable, but a change in the len or cap
	of slice no is reflected in the original var */
	s := make([]int, 2, 10)
	fmt.Println(s)
	updateSlice(s)
	fmt.Println(s)
	// Pointers tasks
	p := newPerson("Andrew","Polo",20)
	pp := newPersonP("p-andrew","p-polo",32)
	fmt.Println(p,pp)
}

func update(px *int) {
	*px = 20
}

func failedUpdate(px *int) {
	x2 := 20
	px = &x2
}

func addStudentToMap(students map[string]string, name string) {
	students[name] = "Alias"
}

/*
	I recommend writing documentation about a function

that modifies or does not modify a Slice.
*/
func updateSlice(s []int) {
	s[1] = 10
	s = append(s, 50)
	fmt.Println(s)
}
