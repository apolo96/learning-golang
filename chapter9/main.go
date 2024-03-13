package chapter9

import (
	"errors"
	"fmt"
	"os"
)

var (
	ErrEmplyeeNotExist = errors.New("the employee no exist")
	ErrUnauthorizer    = errors.New("can not perform the action")
)

var (
	ErrDatabase = ErrT{
		Resource: "Database",
		Code:     123,
	}

	ErrNetwork = ErrT{
		Resource: "Network",
		Code:     456,
	}
)

type ErrT struct {
	Resource string
	Code     int
}

func (e ErrT) Error() string {
	return fmt.Sprintf("%s: %d", e.Resource, e.Code)
}

func (e ErrT) Is(target error) (isErr bool) {
	if t, ok := target.(ErrT); ok {
		matchResource := t.Resource == e.Resource
		matchCode := t.Code == e.Code
		ignoreResource := t.Resource == "" && matchCode
		ignoreCode := t.Code == 0 && matchResource
		isErr = matchResource && matchCode || ignoreCode || ignoreResource
	}
	return isErr
}

func New() {
	fmt.Println("========== Chapter 9")
	fmt.Println("---------- Error handling")
	err := fileChecker("ok.txt")
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println(err)
		}
	}

	if errors.Is(ErrDatabase, ErrT{Resource: "Database"}) {
		fmt.Println("The database is broken:", ErrDatabase)
	}

	if errors.Is(ErrNetwork, ErrT{Resource: "Database"}) {
		fmt.Println("The database is broken:", ErrNetwork)
	}

	if errors.Is(ErrDatabase, ErrT{Code: 123}) {
		fmt.Println("Code 123 triggered:", ErrDatabase)
	}

	if errors.Is(ErrNetwork, ErrT{Code: 123}) {
		fmt.Println("Code 123 triggered:", ErrNetwork)
	}

	if errors.Is(ErrDatabase, ErrT{Resource: "Database", Code: 123}) {
		fmt.Println("Database is broken and Code 123 triggered:", ErrDatabase)
	}

	if errors.Is(ErrDatabase, ErrT{Resource: "Network", Code: 123}) {
		fmt.Println("Network is broken and Code 123 triggered:", ErrDatabase)
	}

	// Panic
	for _, val := range []int{1, 2, 0, 6} {
		div60(val)
	}
	//panic("Failed to bootstrap app")
	/* Error handling Exercise */
	Solve()
}

func div60(i int) {
	defer func() {
		if v := recover(); v != nil {
			fmt.Println(v)
		}
	}()
	fmt.Println(60 / i)
}

func fileChecker(name string) error {
	f, err := os.Open(name)
	if err != nil {
		/* Wrap errors */
		err = fmt.Errorf("in fileChecker: %w", err)
		return fmt.Errorf("new wraper: %w", err)
	}
	defer f.Close()
	return nil
}
