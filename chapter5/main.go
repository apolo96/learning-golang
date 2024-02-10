package chapter5

import (
	"errors"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

type Person struct {
    FirstName string
    LastName  string
    Age       int
}

type MyFuncOpts struct {
	optone string
	opttwo string
}

type opFuncT func(int, int) (int, error)

var opMap = map[string]opFuncT{
	"+": add,
	"-": sub,
	"*": mul,
	"/": div,
}

var myFuncVar func(string) int

func New() {
	fmt.Println("========== Chapter 5")
	// Func with optional and named parameters, using a struct
	MyFunc(MyFuncOpts{optone: "opt1", opttwo: "opt2"})
	// Func variants parameters
	fmt.Println(addTo(3))
	fmt.Println(addTo(3, 2, 4, 6, 8))
	a := []int{4, 3}
	fmt.Println(addTo(3, a...))
	fmt.Println(addTo(3, []int{1, 2, 3, 4, 5}...))
	// Return multi values
	num, d, err := divv(20, 6)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(num, d)
	getUsers(2)
	myFuncVar = f1
	result := myFuncVar("Hello")
	fmt.Println(result)

	myFuncVar = f2
	result = myFuncVar("Hello")
	fmt.Println(result)
	fmt.Println("----------- Func as values")
	// Func as values
	expressions := [][]string{
		{"2", "+", "3"},
		{"2", "-", "3"},
		{"2", "*", "3"},
		{"2", "/", "3"},
		{"2", "%", "3"},
		{"two", "+", "three"},
		{"5"},
	}
	for _, expression := range expressions {
		if len(expression) != 3 {
			fmt.Println("invalid expression:", expression)
			continue
		}
		p1, err := strconv.Atoi(expression[0])
		if err != nil {
			fmt.Println(err)
			continue
		}
		op := expression[1]
		opFunc, ok := opMap[op]
		if !ok {
			fmt.Println("unsupported operator:", op)
			continue
		}
		p2, err := strconv.Atoi(expression[2])
		if err != nil {
			fmt.Println(err)
			continue
		}
		result, err := opFunc(p1, p2)
		if err != nil{
			log.Fatal(err)
		}
		fmt.Println(result)
	}
	fmt.Println("----------- Aninomous Funcs")
	f := func(j int) {
		fmt.Println("Anonimous ", j)
	}
	f(3)
	func(j int) {
		fmt.Println("printing", j, "from inside of an anonymous function")
	}(1)
	// Closures
	fmt.Println("----------- Funcs Closures")
	aa := 20
	fx := func() {
		fmt.Println(aa)
		aa = 30
	}
	fx()
	fmt.Println(aa)
	fmt.Println("----------- Funcs As Parametros")
	people := []Person{
		{"Pat", "Patterson", 37},
		{"Tracy", "Bobdaughter", 23},
		{"Fred", "Fredson", 18},
	}
	sort.Slice(people, func(i, j int) bool{
		return people[i].LastName < people[j].LastName
	})
	fmt.Println(people)
	fmt.Println("----------- Return Funcs")
	twoBase := makeMult(2)
    threeBase := makeMult(3)
    for i := 0; i < 3; i++ {
        fmt.Println(twoBase(i), threeBase(i))
    }
	fmt.Println("----------- Defer")
	defeExample()

	ff, closer, err := getFile("ok")// os.Args[1]
	if err != nil {
		log.Println("Error to getting file")
		//log.Fatal(err)
	}
	fmt.Println(ff)
	defer closer()
	DoSomeData("MyNameIs")
	helloprex := prefixer("Hello")
	fmt.Println(helloprex("Bob"))
	fmt.Println(helloprex("Maria"))
}

func MyFunc(opts MyFuncOpts) error {
	fmt.Println(opts)
	return nil
}

func addTo(base int, vals ...int) []int {
	out := make([]int, 0, len(vals))
	for _, v := range vals {
		out = append(out, base+v)
	}
	return out
}

func divv(num, denum int) (int, int, error) {
	if denum == 0 {
		return 0, 0, errors.New("cannot divvde by zero")
	}
	return num / denum, num % denum, nil
}

func getUsers(count int) (users []string, err error) {
	users = []string{"Jose", "Matias", "Oscar"}
	return users[:count], nil
}

func f1(a string) int {
	return len(a)
}

func f2(a string) (total int) {
	for _, v := range a {
		total += int(v)
	}
	return total
}

func makeMult(base int) func(int) int {
    return func(factor int) int {
        return base * factor
    }
}

func defeExample() {
	a := 10
	defer func(val int){
		fmt.Println("First ",val)
	}(a)	
	b := 20
	defer func(val int)  {
		fmt.Println("Second ",val)
	}(b)
	a  = 30
	fmt.Println("Exiting ", a)	
}

func getFile(name string) (*os.File, func(), error) {
	file, err := os.Open(name)
	closer := func() {		
		file.Close()		        
    }
	if err != nil {
        return nil, closer, err
    }
    return file, closer , nil
}

func DoSomeData(val string) (err error){
	n, err := fmt.Println("Begin SQL Transact", val)
	if err != nil {
        return err
    }
	defer func(){
		if err == nil {
			_, err = fmt.Println("Commit SQL Transact")
		}
		if err != nil {
			fmt.Println("Rollback SQL Transact")
		}
	}()
	_, err = fmt.Println("INSERT INTO FOO (val) values $1" , n)
	return err
}	

func prefixer(val1 string) func(string) string{
	return func(val2 string) string{
		return val1 + val2
	}
}