package main

import "fmt"

// var e string
// e = "cowboy"
// var f string = "fino"

func main() {

	a := 10
	b := "golang" // Podria haberlo declarado al inicio (Ver 'e' y 'f')
	c := 4.17
	d := true

	fmt.Printf("%v \n", a) // %v es un placeholder para el valor por defecto: %t, %d, %g, %s, %p, %x...
	fmt.Printf("%v \n", b)
	fmt.Printf("%v \n", c)
	fmt.Printf("%v \n", d)

	fmt.Println()

}
