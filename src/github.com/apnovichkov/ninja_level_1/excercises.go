package main

import "fmt"

var x int = 42
var y string = "James Bond"
var z bool = true

func main() {
	fmt.Println("Excercise 1:")
	// excerciseOne()
	// excerciseTwo()
	// excerciseThree()
	// excerciseFour()
	excerciseFive()
}

func excerciseOne() {
	x := 42
	y := "James Bond"
	z := true

	fmt.Printf("%v\t%v\t%v\n", x, y, z)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}

func excerciseTwo() {
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}

func excerciseThree() {
	s := fmt.Sprintf("%v\t%v\t%v", x, y, z)
	fmt.Println("b ->", s)
}

func excerciseFour() {
	type hotdog int
	var b hotdog

	fmt.Println("Value of b:", b)
	fmt.Printf("Type of b: %T\n", b)

	b = 42
	fmt.Println("Value of b after asignment:", b)
}

func excerciseFive() {
	type hotdog int
	var x hotdog

	fmt.Println("Value of x:", x)
	fmt.Printf("Type of x: %T\n", x)

	x = 42
	fmt.Println("Value of x after asignment:", x)

	var y int
	y = int(x)

	fmt.Println("Value of y:", y)
	fmt.Printf("Type of y: %T\n", y)

}
