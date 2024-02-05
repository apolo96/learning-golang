package chapter2

import "fmt"

func New() {
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
}
