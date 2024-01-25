package main

import (
	"fmt"
)

// function syntax -> func func_name(parameter_name1 data_type, ...) return type {function_body}
// func add(x int, y int) int {
// 	return x + y
// }

func sayHello() {
	fmt.Println("Hello")
}

// func swap(x, y string) (string, string) {
// 	return y, x
// }

// returning values based on the variable name
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

// var can be package level or function level
var C, Python, Java bool //package level declaration
type ByteSize int

func main() {
	// constants cannot be declared using :=
	const name = "Vaibhavi"
	const age int = 20 //writing data type(here, int) is optional
	// num1, num2, num3 := 9, 6, "three" // := is a short declaration operator - go lang can identify its type at compile time
	// := can only be used in function/curly braces

	// var hobby string = "Reading" //var if not initialized at declaration will assign 0 for all data types
	// var temp int //syntax : var var_name data_type
	// temp = 12
	// var i int //function level declaration

	// fmt.Println("Hello World!!")
	fmt.Printf("My name is %s and I am %d years old\n", name, age)
	// fmt.Printf("My hobby is : %s\n", hobby)
	// fmt.Println(num1, num2, num3, temp)
	// fmt.Println(i, C, Python, Java)

	// %v-value, %b-binary, %x/%X-hexadecimal, %#X-hexadecimal in 0x format, %T-gives type of data
	// fmt.Printf("Number : %d\tBinary : %b\tHexadecimal : %X, %#X\n", num1, num1, num1, num1)
	// fmt.Printf("Number : %d\tBinary : %b\tHexadecimal : %X, , %#X\n", num2, num2, num2, num2)
	// fmt.Printf("Number : %d\tBinary : %b\tHexadecimal : %X, , %#X\n", temp, temp, temp, temp)
	// fmt.Printf("Number : %d\tBinary : %b\tHexadecimal : %X, , %#X\n", 42, 42, 42, 42)

	//Type conversion
	// n1 := 36.93
	// fmt.Printf("Type of variable n1 = %v is : %T\n", n1, n1)
	// var n2 float32 = 36.93
	// fmt.Printf("Type of variable n2 = %v is : %T\n", n2, n2)
	// n1 = float64(n2)
	// fmt.Printf("After Type Conversion : \n")
	// fmt.Printf("Type of variable n1 = %v is : %T\n", n1, n1)
	// fmt.Printf("Type of variable n2 = %v is : %T\n", n2, n2)

	//Random function - range [0, number)
	// fmt.Printf("My favourite number is : %d\n", rand.Intn(11))

	//sqrt function
	// fmt.Printf("Square root of %d is : %g\n", 15, math.Sqrt(15))

	//PI value
	// fmt.Println(math.Pi)

	//function calling
	// fmt.Printf("Addition : %d\n", add(12, 34))
	sayHello()
	// a, b := swap("Vaibhavi", "Bhosale")
	// fmt.Println(a, b)
	fmt.Println(split(12))

	//Basic Types
	//rune data type -> alias for int32 -> represents a unicode code point
	// var (
	// 	toBe   bool       = true
	// 	maxInt uint64     = 1<<64 - 1
	// 	z      complex128 = cmplx.Sqrt(-5 + 12i)
	// )
	// fmt.Printf("Type : %T, Value : %v\n", toBe, toBe)
	// fmt.Printf("Type : %T, Value : %v\n", maxInt, maxInt)
	// fmt.Printf("Type : %T, Value : %v\n", z, z)

	//Type Inference
	// var m int //no value assigned
	// n := m    // since i is declared as int, j will also be the type int
	// fmt.Printf("Type of i : %T, Type of j : %T\n", m, n)
	// f := 3.142 //in this case the type depends on value assigned
	// fmt.Printf("Type of f : %T\n", f)

	// _ (blank identifier): used to neglect or throw out the variable value
	// iota
	const (
		_ = iota //c0 = 0
		a        //c1 = 1
		b        //c2 = 2 and so on
		c        //it data type not mentioned it takes it as first data type assigned
		d
		e
		f
	)

	const (
		_           = iota
		KB ByteSize = 1 << (10 * iota)
		MB
		GB
		TB
		PB
		EB
	)

	//Bitwise shift operator
	// fmt.Printf("%d\t%b\n", 1, 1)
	// fmt.Printf("%d\t%b\n", 1<<1, 1<<1)
	// fmt.Printf("%d\t%b\n", 1<<2, 1<<2)
	// fmt.Printf("%d\t%b\n", 1<<3, 1<<3)
	// fmt.Printf("%d\t%b\n", 1<<4, 1<<4)
	// fmt.Printf("%d\t%b\n", 1<<5, 1<<5)
	// fmt.Printf("%d\t%b\n", 1<<6, 1<<6)

	fmt.Printf("%v\t%T\n", a, a)
	fmt.Printf("%d\t%b\n", 1<<a, 1<<a)
	fmt.Printf("%d\t%b\n", 1<<b, 1<<b)
	fmt.Printf("%d\t%b\n", 1<<c, 1<<c)
	fmt.Printf("%d\t%b\n", 1<<d, 1<<d)
	fmt.Printf("%d\t%b\n", 1<<e, 1<<e)
	fmt.Printf("%d\t%b\n", 1<<f, 1<<f)

	fmt.Printf("%v\t\t%T\n", KB, KB)
	fmt.Printf("%d\t\t%b\n", MB, MB)
	fmt.Printf("%d\t\t%b\n", GB, GB)
	fmt.Printf("%d\t\t%b\n", TB, TB)
	fmt.Printf("%d\t\t%b\n", PB, PB)
	fmt.Printf("%d\t\t%b\n", EB, EB)

}
