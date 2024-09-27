package main

import "fmt"

func main() {
	var a string
	a = "foo"

	fmt.Println(a)

	var b int = 99
	fmt.Println(b)

	var c = true

	fmt.Println(c)

	d := 3.1415 // go will assign the default, which is float64, you may want to explicit the type depending on the code.

	fmt.Println(d)

	var e int = int(d) // you need to explicit convert, GO DOESNT support implicit conversion

	fmt.Println(e)

	var f int8 = 6

	var g int16 = int16(f) // see like here, you will need to explicit it. even they are in the "same" datatype.

	fmt.Println(g)

	// const can be definied in group, if you leave one without value, it will copy the above one.
	const (
		te = 42
		ef = "ae"
		copied
	)
	// you can also create a complex const, unless it's able to run on compile time. You cannot allocate memory to create that.

	//pointers

	s := "Hello, world!"

	p := &s
	fmt.Println(p)
	fmt.Println(*p)

	*p = "Hello, Pozzani"

	fmt.Println(s)

	p = new(string)

	fmt.Println(p)

	*p
}
