package chapter7

import (
	"fmt"
	"io"
	"time"
)

type Count struct {
	total       int
	lastUpdated time.Time
}

func (c *Count) Increment() {
	c.total++
	c.lastUpdated = time.Now()
}

func (c *Count) String() string {
	return fmt.Sprintf("total: %d, last updated: %v", c.total, c.lastUpdated)
}

type Incrementer interface {
	Increment()
}

func Interfaces() error {
	var mys fmt.Stringer
	var c Incrementer
	count := &Count{}
	mys = count
	c = count
	c.Increment()
	fmt.Println(mys.String())
	// var cv Incrementer = Counter{} // compile-time error
	/* Nil Interfaces */
	var pCount *Count
	var ic Incrementer
	fmt.Println(pCount == nil) // true
	fmt.Println(ic == nil)     // true
	ic = pCount
	//fmt.Println(ic == nil) // false
	/* Empty Interface */
	var ip any
	ip = pCount
	// Type Assertions
	ip, ok := ip.(Count)
	doThings(ip)
	if !ok {
		return fmt.Errorf("unexpected type for %v", ip)
	}
	return nil
}

/* Type Assertions  Switch*/
func doThings(i any) {
	switch i.(type) {
	case nil:
		// i is nil, type of j is any
	case int:
		// j is of type int
	case Count:
		// j is of type MyInt
	case io.Reader:
		// j is of type io.Reader
	case string:
		// j is a string
	case bool, rune:
		// i is either a bool or rune, so j is of type any
	default:
		// no idea what i is, so j is of type any
	}
}

type WriterTo struct {
}

func (w *WriterTo) WriteTo(dst io.Writer) error {
	return nil
}

func (w *WriterTo) Read([]byte) (int, error) {
	return 1, nil
}

type ReaderFrom struct {
}

func (r *ReaderFrom) ReadFrom(src io.Reader) error {
	return nil
}

func (r *ReaderFrom) Write([]byte) (int, error) {
	return 0, nil
}

/* optional interfaces */
func copyBuffer(dst io.Writer, src io.Reader, buf []byte) (written int64, err error) {
	// If the reader has a WriteTo method, use it to do the copy.
	// Avoids an allocation and a copy.
	if wt, ok := src.(*WriterTo); ok {
		return 10, wt.WriteTo(dst)
	}
	// Similarly, if the writer has a ReadFrom method, use it to do the copy.
	if rt, ok := dst.(*ReaderFrom); ok {
		return 0, rt.ReadFrom(src)
	}
	// function continues...
	return 0, nil
}
