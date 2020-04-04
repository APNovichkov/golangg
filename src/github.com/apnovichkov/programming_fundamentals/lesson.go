package main

import "fmt"

var x bool

const a = 42
const b = 33.3123
const c = "Hello World"

const (
	// kb = 1024
	// mb = 1024 * kb
	// gb = 1024 * mb

	_  = iota
	kb = 1 << (iota * 10)
	mb = 1 << (iota * 10)
	gb = 1 << (iota * 10)
)

// or
const (
	d = 23
	e = 22.3234
	f = "James Bond"
)

//Iotas
const (
	i = iota
	j = iota
	k = iota
)

func main() {
	// booleans()
	// numericals()
	// strings()
	// iotas()
	bitShifting()
}

func booleans() {
	fmt.Println(x)
	x = true
	fmt.Println(x)

	a := 4
	b := 42

	fmt.Println(a < b)
}

func numericals() {
	var x int
	var y float64

	x = 42
	y = 42.234

	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
}

func strings() {
	x := "Hello World"
	bs := []byte(x)

	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)

	for i := 0; i < len(x); i++ {
		fmt.Printf("%#U ", x[i])
	}

	fmt.Println()

	for i, v := range x {
		fmt.Printf("At index position %d we have hex %#x\n", i, v)
	}
}

func constants() {
	// Look Above for declaration
}

func iotas() {
	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k)
}

func bitShifting() {
	x := 2
	fmt.Printf("%d\t\t%b\n", x, x)

	y := x << 2
	fmt.Printf("%d\t\t%b\n\n", y, y)

	fmt.Printf("%d\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t%b\n", mb, mb)
	fmt.Printf("%d\t%b\n", gb, gb)

	// Shifting over 10 bits every time. Can use iotas to accomplish this!

}
