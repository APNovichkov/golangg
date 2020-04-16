package main

import "fmt"

// Part of Problem 3
const a = 40
const b int = 40
const (
	c      = true
	d bool = true
)

//Part of Problem 6
const (
	year1 = 2020 - iota
	year2 = 2020 - iota
	year3 = 2020 - iota
	year4 = 2020 - iota
)

func main() {
	// problem1()
	// problem2()
	// problem3()
	// problem4()
	// problem5()
	problem6()
}

func problem1() {
	x := 20

	fmt.Printf("Regular: %v\n", x)
	fmt.Printf("Binary: %b\n", x)
	fmt.Printf("Hex: %#x\n", x)
}

func problem2() {
	x := 20
	y := 30

	r1 := x == y
	r2 := x <= y
	r3 := x >= y
	r4 := x < y
	r5 := x > y
	r6 := x != y

	fmt.Printf("%v\t%v\t%v\t%v\t%v\t%v\n", r1, r2, r3, r4, r5, r6)
}

func problem3() {
	fmt.Println("No type constant:", a)
	fmt.Println("Int type constant:", b)

	fmt.Println("No type constant:", c)
	fmt.Println("Bool type constant:", d)
}

func problem4() {
	x := 42

	fmt.Printf("Regular: %v\n", x)
	fmt.Printf("Binary: %b\n", x)
	fmt.Printf("Hex: %#x\n", x)

	fmt.Println("Shifting one bit to the left!")

	b := x << 1
	fmt.Printf("Regular: %v\n", b)
	fmt.Printf("Binary: %b\n", b)
	fmt.Printf("Hex: %#x\n", b)
}

func problem5() {
	x := `"hello world!"`
	fmt.Println(x)
}

func problem6() {
	fmt.Println(year1, year2, year3, year4)
}
