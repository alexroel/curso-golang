package main

import (
	"fmt"
	"math/cmplx"
)

func main() {
	option, name, a, b, c, d := true, "Alex", -27, 45, 4.5, cmplx.Sqrt(-5+12i)

	fmt.Printf("Tipo %T valor %v\n", option, option)
	fmt.Printf("Tipo %T valor %v\n", name, name)
	fmt.Printf("Tipo %T valor %v\n", a, a)
	fmt.Printf("Tipo %T valor %v\n", b, b)
	fmt.Printf("Tipo %T valor %v\n", c, c)
	fmt.Printf("Tipo %T valor %v\n", d, d)
}
