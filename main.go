package main

import (
	"fmt"
	"learngolang/chapter2"
	"learngolang/chapter3"
	"learngolang/chapter4"
	"learngolang/chapter5"
)

func main() {
	chapter2.New()
	chapter3.New()
	chapter4.New()
  chapter5.New()
	search()
	fmt.Println(runners)
}
