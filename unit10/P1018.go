package main

import (
	"fmt"
)

type A struct {
	a int
	b string
}

type B struct {
	c int
	d string
	a int
}

type C struct {
	A
	B
}

func main() {

	c := C{A{10, "aaa"}, B{20, "ccc", 50}}
	fmt.Println(c.b)
	fmt.Println(c.d)
	fmt.Println(c.B)
	fmt.Println(c.A.a)
	fmt.Println(c.B.a)
}
