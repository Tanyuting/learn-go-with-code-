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
	int
}

type C1 struct {
	*A
	*B
	int
}

func main() {

	c := C{A{10, "aaa"}, B{20, "ccc", 50}, 888}
	fmt.Println(c.b)
	fmt.Println(c.d)
	fmt.Println(c.B)
	fmt.Println(c.A.a)
	fmt.Println(c.B.a)
	fmt.Println(c.int)

	c1 := C1{&A{10, "aaa"}, &B{20, "ccc", 50}, 888}
	fmt.Println(c1)
	fmt.Println(*c1.A)
}
