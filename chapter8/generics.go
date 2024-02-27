package chapter8

import "fmt"

type Stack[T comparable] struct {
	vals []T
}

func (s *Stack[T]) Push(val T) {
	s.vals = append(s.vals, val)
}

func (s *Stack[T]) Pop() (top T, ok bool) {
	if len(s.vals) == 0 {
		var zero T
		return zero, false
	}
	cut := len(s.vals) - 1
	top = s.vals[cut]
	s.vals = s.vals[:cut]
	return top, true
}

func (s *Stack[T]) Contains(val T) (ok bool) {
	for _, v := range s.vals {
		if v == val {
			ok = true
		}
	}
	return ok
}

func New() {
	fmt.Println("========== Chapter 8")
	var s Stack[int]
	s.Push(10)
	s.Push(20)
	s.Push(30)
	fmt.Println(s.Contains(10))
	fmt.Println(s.Contains(5))
	fmt.Println("Generic Map func")
	arr := []string{"aa", "a", "wad"}
	// Getting items len
	fmt.Println(Map(arr, func(v string) int {
		return len(v)
	}))
	// Getting sumatory len items
	fmt.Println(Map(arr, func(v string) int {
		return len(v) + 2
	}))
	fmt.Println("Generic Reduce func")
	fmt.Println(Reduce(arr, 2, func(init int, v string) int {
		return len(v) + init
	}))
	fmt.Println("Generic Reduce func")
	words := []string{"One", "Potato", "Two", "Potato"}
	ff := Filter(words, func(s string) bool {
		return s != "Potato"
	})	
	fmt.Printf("cap: %v len: %v value: %v type: %T \n", cap(ff), len(ff), ff, ff)

}

// Map() transforms the items of slice using a function
func Map[T, Y any](s []T, f func(T) Y) []Y {
	m := make([]Y, 0, len(s))
	for _, v := range s {
		m = append(m, f(v))
	}
	return m
}

/*
Reduce() reduces a only value the items of slice
using a operator func wiht initizalizer
*/
func Reduce[T, Y any](s []T, init Y, f func(Y, T) Y) (sy Y) {
	sy = init
	for _, v := range s {
		sy = f(sy, v)
	}
	return sy
}

// Filter() filter the items of slice using comparable operation on a function
func Filter[T any](s []T, f func(T) bool) (sy []T) {
	for _, v := range s {
		if f(v) {
			sy = append(sy, v)
		}
	}
	return sy
}
