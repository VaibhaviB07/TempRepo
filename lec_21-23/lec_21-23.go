package main

import (
	"fmt"

	"golang.org/x/exp/constraints" //package constraints
)

type Dog struct {
	name string
}

// Method Sets
func (d Dog) walk() {
	fmt.Println("Mt name is ", d.name, " and I'm walking")
}
func (d *Dog) run() {
	d.name = "Bob"
	fmt.Println("Mt name is ", d.name, " and I'm running")
}

type youngIn interface {
	walk()
	run()
}

func youngRun(y youngIn) {
	y.run()
}

// Call by Reference
// Pointer as parameter in function
func intDelta(n *int) {
	*n = 3
}
func sliceDelta(s []int) {
	s[0] = 0
}
func mapDelta(m map[string]int, s string) {
	m[s] = 21
}

// Value semantics
func addOne(n int) int {
	return n + 1
}

// Pointer semantics
func addOnePtr(n *int) {
	*n = *n + 1
}

// Generic function which adds both integers and floats
// instead of writing 2 different functions -> DRY
// Type Constraints
func addT[T int | float64](a, b T) T {
	return a + b
}

// Type Alias & underlying type constraints
// can be modified with ~ in type set interface
type myAlias int

// Generic function -> Type set Interface
type myDataTypes interface {
	// ~int | ~float64
	constraints.Integer | constraints.Float
}

func add2T[T myDataTypes](a, b T) T {
	return a + b
}

func main() {
	//POINTERS
	//&var_name - used to get the address of the var_name
	//*var_name - used to get value of var whose address is assigned to var_name
	x := 6
	fmt.Printf("Value of x: %v\nAddress of x: %v\nType of &x: %T\n", x, &x, &x)

	y := &x
	//*&x - works the same but gives warning to simplify
	fmt.Println("Value of &x stored in y:", y)
	fmt.Println("Value of x using y:", *y)

	*y = 9
	fmt.Println("Changing value of x with pointer variable y:", x)

	fmt.Println("Before Pointer as parameter in function: ", x)
	intDelta(&x)
	fmt.Println("After Pointer as parameter in function: ", x)

	slice := []int{-1, 1, 2, 3}
	fmt.Println(slice)
	sliceDelta(slice)
	fmt.Println(slice)

	m := make(map[string]int)
	m["Vaibhavi"] = 18
	fmt.Println(m["Vaibhavi"])
	mapDelta(m, "Vaibhavi")
	fmt.Println(m["Vaibhavi"])

	//Value Semantics -> passing values to function -> copied values(pass by value)
	num1 := 12
	fmt.Println("Value Semantics : ")
	fmt.Println(num1)         //num1 = 12
	fmt.Println(addOne(num1)) //num1 = 13
	fmt.Println(num1)         //num1 = 12

	//Pointer Semantics -> passing pointers to function -> shared memory
	num2 := 15
	fmt.Println("Pointer Semantics : ")
	fmt.Println(num2) //num2 = 15
	addOnePtr(&num2)  //num2 = 16
	fmt.Println(num2) //num2 = 16

	//When we use Pointer semantics, it is more likely that it
	//will be moved to Heap from Stack

	//Printing escape analysis and inlining decisions
	//go run -gcflags -m main.go
	//go run -gcflags -m=2 main.go

	//METHOD SETS
	//IMPORTANT -> Method sets of a type determines the interfaces that the type implements
	//it is the set of methods attached to a type.
	//associated with both value and pointer types
	//integral to how interfaces are implemented
	//The method set of type T consists of all methods with receiver type T
	//The method set of type *T consists of all methods with receiver type *T OR T

	d1 := Dog{
		name: "Yeontan",
	}
	d1.walk()
	d1.run()
	// youngRun(d1) -> Error : method set run has pointer receiver

	d2 := &Dog{
		name: "Bam",
	}
	d2.walk()
	d2.run()
	youngRun(d2)

	//GENERICS
	//allow us to abstract our code and make it more DRY
	//DRY -> Do Not Repeat Yourself
	//syntax -> func func_name[T data types separated with |](arguments , separated T) T {body}
	fmt.Println(addT(2, 4))
	fmt.Println(addT(2.89, 4.34))

	fmt.Println(add2T(2, 4))
	fmt.Println(add2T(2.89, 4.34))

	var n myAlias = 23
	fmt.Println(add2T(n, 4))
}
