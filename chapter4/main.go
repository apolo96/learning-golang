package chapter4

import (
	"fmt"
	"math/rand"
)

func New() {
	fmt.Println("======= Chapter 3")
	ok := num + xnum
	fmt.Println(ok)
	fmt.Println("------- If ")
	// If structure
	if n := rand.Intn(100); n == 0 {
		fmt.Println("That's too low")
	} else if n < 5 {
		fmt.Println("That's too big:", n)
	} else {
		fmt.Println("That's a good number:", n)
	}
	/* For statement to loop */
	fmt.Println("------- For Loop ")
	// Complete for Statement
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	// While For
	for i := 0; i < 10; {
		fmt.Println(i)
		i++
	}
	// Conditional For
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}
	// Infinite For statement
	/* for{
		fmt.Println("Infinite Lopp")
	} */
	i = 0
	// Do While - For
	for {
		fmt.Println("Do While")
		if i > 10 {
			break
		}
		i++
	}
	fmt.Println("------ Continue sentence")
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
			continue
		}
		if i%3 == 0 {
			fmt.Println("Fizz")
			continue
		}
		if i%5 == 0 {
			fmt.Println("Buzz")
			continue
		}
		fmt.Println(i)
	}
	fmt.Println("------- For Range")
	evenVals := []int{2, 4, 6, 8, 10, 12}
	// Get index and value of array
	for i, v := range evenVals {
		fmt.Println(i, v)
	}
	// Avoid index and get only value of array
	for _, v := range evenVals {
		fmt.Println(v)
	}
	// Get only the index
	for i := range evenVals {
		fmt.Println(i)
	}
	// For range maps
	uniqueNames := map[string]bool{"Fred": true, "Raul": true, "Wilma": true}
	fmt.Println(uniqueNames)
	// Get only keys of map
	for k := range uniqueNames {
		fmt.Println(k)
	}
	// For range strings
	samples := []string{"hello", "apple_π!"}
	for _, sample := range samples {
		for i, r := range sample {
			fmt.Println(i, r, string(r))
		}
		fmt.Println()
	}
	// For tag loop
outer:
	for _, sample := range samples {
		for i, r := range sample {
			fmt.Println(i, r, string(r))
			if r == 'l' {
				continue outer
			}
		}
		fmt.Println()
	}
	fmt.Println("------- Switch")
	words := []string{"a", "cow", "smile", "gopher", "octopus", "anthropologist"}
	for _, word := range words {
		switch size := len(word); size {
		case 1, 2, 3, 4:
			fmt.Println(word, "is a short word!")
		case 5:
			fmt.Println(word, "is exactly the right length:", size)
		case 6, 7, 8, 9:
		default:
			fmt.Println(word, "is a long word!")
		}
	}
loop:
	for i := 0; i < 10; i++ {
		switch i {
		case 0, 2, 4, 6:
			fmt.Println(i, "is even")
		case 3:
			fmt.Println(i, "is divisible by 3 but not 2")
		case 7:
			fmt.Println("exit the loop!")
			break loop
		default:
			fmt.Println(i, "is boring")
		}
	}

	for _, word := range words {
		switch wordLen := len(word); {
		case wordLen < 5:
			fmt.Println(word, "is a short word!")
		case wordLen > 10:
			fmt.Println(word, "is a long word!")
		default:
			fmt.Println(word, "is exactly the right length.")
		}
	}
	a := 2
	switch a {
	case 2:
		fmt.Println("a is 2")
	case 3:
		fmt.Println("a is 3")
	case 4:
		fmt.Println("a is 4")
	default:
		fmt.Println("a is ", a)
	}
	fmt.Println("------- resolveChapterTask")
	resolveChapterTask()
}

func resolveChapterTask() {
	// Task One
	numbers := make([]int, 0, 100)
	for i := 1; i <= 100; i++ {
		numbers = append(numbers, rand.Intn(100))
	}
	fmt.Println(numbers, len(numbers))
	// Task Two
	for _, v := range numbers {
		if v%2 == 0 {
			fmt.Println("Two")
			continue
		}
		if v%3 == 0 {
			fmt.Println("Three")
			continue
		}
		if v%2 == 0 && v%3 == 0 {
			fmt.Println("Six")
			continue
		}
		fmt.Println("Never mind")
	}
	// Task Three
	var total int
	for _, v := range numbers {
		total += v
	}
	fmt.Println("Total: ", total)
}
