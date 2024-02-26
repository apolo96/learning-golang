package chapter7

import "fmt"

/* It’s clearer for someone reading your code when a method has
a parameter of type Percentage than of type int, and it’s
harder for it to be invoked with an invalid value. */

type Employee struct {
	Name string
	ID   string
	X    int
}

func (e Employee) Description() string {
	return fmt.Sprintf("Employee %s (%s) \n", e.Name, e.ID)
}

func (e Employee) sX() string {
	return fmt.Sprintln("X employe",e.X)
}

/* Embedding Is Not Inheritance */
type Manager struct {
	Employee // embedded type
	Reports  []Employee
	X        int
}

func (m Manager) Description() string {
	return fmt.Sprintf("Manager X %v \n", m.X)
}

func Composition() {
	m := Manager{
		Employee: Employee{
			Name: "Bob Bobson",
			ID:   "12345",
			X:    10,
		},
		Reports: []Employee{},
		X:       20,
	}
	fmt.Println(m.ID)
	fmt.Println(m.sX())
	fmt.Println(m.Employee.Description())
	fmt.Println(m.X)
	fmt.Println(m.Employee.X)
	fmt.Println(m.Description())
}
