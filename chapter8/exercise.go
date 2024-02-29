package chapter8

import "fmt"

func Multi[T int | float64](n T) T{
	return n * 2
}

type Printable interface {
	~int | ~float64
	fmt.Stringer
}

func Printp[T Printable](v T){
	fmt.Println(v.String())
}
