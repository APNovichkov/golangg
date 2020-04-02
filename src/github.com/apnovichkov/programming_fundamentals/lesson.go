package main

import "fmt"

var x bool

func main() {
	booleans()
}

func booleans() {
	fmt.Println(x)
	x = true
	fmt.Println(x)

	a := 4
	b := 42

	fmt.Println(a < b)

}
