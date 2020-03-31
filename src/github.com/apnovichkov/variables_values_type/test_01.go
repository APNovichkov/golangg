package main

import "fmt"

var z = 50

// Can assign types to vars
var i int = 400
var f float32 = 2.33
var w string = `I said "hello my name is Andrey"`

func main() {
	// Short Declaration Operator
	// x := 2
	// fmt.Println(x)
	// x = 200
	// fmt.Println(x)

	// Var Keyword, Can put it outside of a function body
	// var y = 122
	// fmt.Println(y)
	// fmt.Println(z)

	// ***Try to limit the scope of the variable as much as possible, and use the
	// short declaration operator as much as possible!
	// foo()

	// fmt.Println("Integer:", i)
	// fmt.Printf("%T\n", i)
	// fmt.Println("Float:", f)
	// fmt.Printf("%T\n", f)
	// fmt.Println("String:", w)
	// fmt.Printf("%T\n", w)

	// fmtTest()

	creatingType()
}

func fmtTest() {
	// Printf lets you format your variables into different types
	fmt.Println("normal:", z)
	fmt.Printf("Type: %T\n", z)
	fmt.Printf("Binary: %b\n", z)
	fmt.Printf("Hexadecimal: %x\n", z)
	fmt.Printf("Hexadecimal w/ Ox: %#x\n", z)
	fmt.Printf("Random: %b\t%x\t%#x\n\n", z, z, z)

	// Sprint functions print to a variable, int this case s
	s := fmt.Sprintf("Random: %b\t%x\t%#x", z, z, z)
	fmt.Println(s)

	// Prints normal line
	fmt.Printf("%v", z)
}

func creatingType() {
	var a int
	// Create our own type hotdog whose underlying type is int
	type hotdog int
	var b hotdog

	a = 42
	b = 43

	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(b)
	fmt.Printf("%T\n", b)

	a = int(b)
	fmt.Println(a)
	fmt.Printf("%T\n", a)

	b = hotdog(a)
	fmt.Println(b)
	fmt.Printf("%T\n", b)
}

func foo() {
	var test int = 4
	fmt.Println("My name is Andrey, here is z:", z)
	fmt.Println(test)
}

func loopy() {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}
