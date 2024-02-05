package chapter3

import (
	"fmt"
	"maps"
	"slices"
)

func New() {
	fmt.Println("------ Arrays")
	fmt.Println("Chapter 3")
	var arr = [...]int{10, 3, 2}
	printArray(arr)

	ar := [3]int{}
	printArray(ar)
	fmt.Println("------ Slices")
	/* Using [...] makes an array. Using [] makes a slice. */
	var s = []int{1, 2, 3, 8}
	// Empty Slice not nil
	var ss = []int{}
	// Empty Slice nil
	var sx []int
	printSlice(sx)
	fmt.Println("Is slice equeal to nil?", sx == nil)
	fmt.Println("Are slices equals? ", slices.Equal(ss, s))
	printSlice(s)
	printSlice(ss)
	s = append(s, 11)
	s = append(s, ss...)
	printSlice(s)
	/* 	Right Way for initial slice */
	sr := make([]int, 5) // [0,0,0,0,0]
	sr = append(sr, 10)
	printSlice(sr)
	/* Initial Empty slice, the follow code is the best way */
	sr = make([]int, 0, 10)
	sr = append(sr, 5)
	printSlice(sr)
	/* Clear Slice */
	clear(sr)
	/* Slicing slices */
	s = s[:2]
	printSlice(s)
	/* Convert array to slice */
	s = arr[:]
	printSlice(s)
	// Secure slicing slices
	x := make([]string, 0, 5)
	x = append(x, "a", "b", "c", "d")
	y := x[:2:2]
	z := x[2:4:4]
	fmt.Println(cap(x), cap(y), cap(z))
	/* Copy Slices */
	num := copy(y, x)
	fmt.Println(num)
	fmt.Println(y)
	/* Convert a Slice to Array */
	xSlice := []int{1, 2, 3, 4}
	xArray := [4]int(xSlice) // You can't convert a slice to an array with a greater size than a slice
	fmt.Println(xArray)

	/* Maps */
	// Initial a Map
	fmt.Println("------ Maps")
	var nilMap map[string]int
	totalWins := map[string]int{}
	teams := map[string][]string{
		"Red":   {"David", "Juls"},
		"White": {"Jose", "Pier"},
	}
	fmt.Println(nilMap, nilMap == nil)
	fmt.Println(totalWins, totalWins == nil)
	fmt.Println(teams)
	// Read a item in the Map
	fmt.Println(teams["Red"])
	// Check If exist a item in the Map
	t, ok := teams["Orange"]
	fmt.Println(t, ok)
	// Write a item in the Map
	teams["Gray"] = []string{"Polo", "Pep"}
	fmt.Println(teams)
	nilMap = make(map[string]int)
	fmt.Println(nilMap, len(nilMap))
	nilMap["hello"] = 12
	totalWins["job"] = 90
	// Delete a item in the Map
	delete(teams, "Gray")
	fmt.Println(teams)
	// Clear a Map
	clear(totalWins)
	fmt.Println(totalWins)
	// Comparing Maps
	fmt.Println("Map 1 is equal to Map 2 ", maps.Equal(totalWins, nilMap))
	// Maps as Collects
	collect := map[int]bool{}
	for _, v := range []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10} {
		collect[v] = true
	}
	if collect[4] {
		fmt.Println("There is a number")
	}
	fmt.Println("------ Structs")
	// Declare a struct, YOU CAN DECLARE Structs to func or package level
	type Hello struct {
		title     string
		isComming bool
	}
	// Initial a struct
	var hello Hello
	fmt.Println(hello)
	h := Hello{"Hi There", false}
	fmt.Println(h)
	hh := Hello{
		title:     "Ey My friend",
		isComming: true,
	}
	fmt.Println(hh)
	// Access a struct attribute
	fmt.Println(hh.title)
	// Update struct
	h.title = "New Dear"
	// Anonimous struct
	pet := struct {
		msg  string
		kind string
	}{
		msg:  "OK",
		kind: "All",
	}
	fmt.Println(pet)
	resolveChapterTask()
}

func printArray(arr [3]int) {
	fmt.Printf("Array len: %v value: %v type: %T \n", len(arr), arr, arr)
}

func printSlice(s []int) {
	fmt.Printf("Slice cap: %v len: %v value: %v type: %T \n", cap(s), len(s), s, s)
}

func resolveChapterTask() {
	// Task one
	greetings := []string{"Hello", "Hola", "नमस्कार", "こんにちは", "Привіт"}
	subG1 := greetings[:2:2]
	subG2 := greetings[1:4:4]
	subG3 := greetings[3:5:5]
	fmt.Println(subG1, subG2, subG3, greetings)
	// Task Two
	message := "Hi 👩 and 👨"
	utf8 := []rune(message)	
	fmt.Println(string(utf8[3]))
	// Task Three
	type Employee struct {
		firstName string
		lastName  string
		id        int
	}
	emp := Employee{
		"Andrew",
		"Polo",
		12,
	}
	emp1 := Employee{
		firstName: "Paul",
		lastName:  "Torre",
		id:        3,
	}
	var emp2 Employee
	emp2.firstName = "P1"
	emp2.lastName = "P2"
	emp2.id = 201
	fmt.Println(emp, emp1, emp2)
}
