package main

import "fmt"

var z = 50

// Can assign types to vars
var i int = 400
var f float32 = 2.33
var w string = `I said "hello my name is Andrey"`

func main() {
	// Short Declaration Operator
	x := 2
	fmt.Println(x)
	x = 200
	fmt.Println(x)

	// Var Keyword, Can put it outside of a function body
	var y = 122
	fmt.Println(y)
	fmt.Println(z)

	// ***Try to limit the scope of the variable as much as possible, and use the
	// short declaration operator as much as possible!
	foo()

	fmt.Println("Integer:", i)
	fmt.Printf("%T\n", i)
	fmt.Println("Float:", f)
	fmt.Printf("%T\n", f)
	fmt.Println("String:", w)
	fmt.Printf("%T\n", w)

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
