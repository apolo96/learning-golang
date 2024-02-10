package chapter5

import "errors"

func add(i int, j int) (int, error) { return i + j, nil }

func sub(i int, j int) (int, error) { return i - j, nil }

func mul(i int, j int) (int, error) { return i * j, nil }

func div(i int, j int) (int, error) { 
	if j == 0{
		return 0, errors.New("division by zero")
	}
	return i / j, nil
}