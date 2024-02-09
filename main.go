package main

import (
	"fmt"
	"learngolang/chapter2"
	"learngolang/chapter3"
	"learngolang/chapter4"
)

func main() {
	chapter2.New()
	chapter3.New()
	chapter4.New()
	search()
	fmt.Println(runners)
}
