package chapter2

import "fmt"

const x int64 = 10

const (
    idKey   = "id"
    nameKey = "name"
)

func New() {
	fmt.Println("Chapter 2")
	// Task one
	var i int = 20
	var f float64 = float64(i)
	fmt.Println(i, f)
	// Task Two
	const value = 10
	i = value
	f = value
	// Task Three
	var b byte = 255
	var smallI int32 = 2147483647
	var bigI uint64 = 18446744073709551615
	b += 1
	smallI += 1
	bigI += 1
	fmt.Println(b, smallI, bigI)
	// Types Conversions
	var x int = 10
	var y float64 = 30.2
	var sum1 float64 = float64(x) + y
	var sum2 int = x + int(y)
	fmt.Println(sum1,sum2)
	// If you are referring to a character, use the rune type
	var myFirstInitial rune = 'J' 
	fmt.Println(myFirstInitial)
}
