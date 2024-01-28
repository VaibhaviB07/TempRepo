package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

type Person struct {
	first   string
	last    string
	age     int
	hobbies []string
}

// Embedded struct
type Employee struct {
	Person
	salary float64
}

type Human struct {
	name string
}

type Girl struct {
	Human
	gender string
}

type Book struct {
	title string
}

// INTERFACE
type People interface {
	speak()
}

func foo() {
	fmt.Println("I am from foo")
}

func bar(s string) {
	fmt.Println("My Name Is : ", s)
}

func aloha(s string) string {
	return fmt.Sprint("Hey, I am ", s)
}

func intro(name string, age int) (string, int) {
	return fmt.Sprint(name, " is ", age, " years old"), age + 5
}

// parameter as unlimited elements of specific data types are converted into slice
func sum(temp ...int) int {
	// fmt.Println(temp)
	// fmt.Printf("%T\n", temp)
	sum := 0
	for _, v := range temp {
		sum += v
	}

	return sum
}

// METHOD
func (h Human) speak() {
	fmt.Println("I am ", h.name)
}

func (g Girl) speak() {
	fmt.Println("I am a ", g.gender, ", My name is ", g.name)
}

// For Interface
func saySomething(pips People) {
	pips.speak()
}

// String method -> method signature for interface
func (b Book) String() string {
	return fmt.Sprint("Tile of the Book : ", b.title)
}

type count int

func (c count) String() string {
	return fmt.Sprint("This is a number : ", strconv.Itoa(int(c))) //convert int to char/string
}

// Wrapper function -> it provides an additional layer of abstraction/functionality around an
// existing method or function
func logInfo(s fmt.Stringer) {
	log.Println("LOG INFO -> ", s.String())
}

// Wrapper function to read file
func readFile(filename string) ([]byte, error) {
	xb, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("error in readFile(): %s", err)
	}
	return xb, nil
}

// Wrapper function example
func doWork() {
	for i := 0; i < 2000; i++ {
		// fmt.Println(i)
	}
}
func timeFunc(f func()) {
	start := time.Now()
	f()
	elapsed := time.Since(start)
	fmt.Println(elapsed)
}

// Use of Write Interface
// func (h Human) writeOut(w io.Writer) error {
// 	_, err := w.Write([]byte(h.name))
// 	return err
// }

// Returning a Function
func temp1() int {
	return 36
}
func temp2() func() int {
	return func() int {
		return 63
	}
}

// CALLBACK
func doMath(a int, b int, f func(int, int) int) int {
	return f(a, b)
}
func adding(a int, b int) int {
	return a + b
}
func subtract(a int, b int) int {
	return a - b
}

// CLOSURE
func incrementor() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

// Recursion
func factorial(n int) int {
	fmt.Println("n :", n)
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	//STRUCT
	p1 := Person{
		first:   "Vaibhavi",
		last:    "Bhosale",
		age:     20,
		hobbies: []string{"reading", "movie watching"},
	}
	fmt.Println(p1)
	// fmt.Printf("Type: %T, %#v\n", p1, p1)
	// fmt.Printf("First: %v, Last: %v, Age: %v\n", p1.first, p1.first, p1.age)

	e1 := Employee{
		Person: Person{
			first: "Sumit",
			last:  "Bhosale",
			age:   26,
		},
		salary: 50000,
	}
	fmt.Println(e1)
	// fmt.Printf("Type: %T, %#v\n", e1, e1)
	// fmt.Printf("First: %v, Last: %v, Age: %v, Salary: %v\n", e1.first, e1.first, e1.age, e1.salary)

	//Anonymous struct -> for Person
	//They are indeterminate, there type is not yet known
	p2 := struct {
		first string
		last  string
		age   int
	}{
		first: "Taehyung",
		last:  "Kim",
		age:   27,
	}
	fmt.Println(p2)
	// fmt.Printf("Type: %T, %#v\n", p2, p2)

	//FUNCTION
	//Purpose : Modular & Abstract Code, Code Reusability, DRY -> Don't Repeat Yourself
	//syntax : func(receiver) identifier(parameters)(returns){code}
	foo()
	bar("Vaibhavi")
	fmt.Println(aloha("Sumit")) //Sprint() -> prints string
	s, new_age := intro("Vaibhavi", 20)
	fmt.Println("Intro : ", s)
	fmt.Printf("Age after 5 years : %v\n", new_age)

	add := sum(1, 2, 3, 4, 5, 6) //method1
	x := []int{6, 7, 8, 9, 10, 11, 12}
	xi := sum(x...) //method 2
	fmt.Printf("The Sum is : %v\n", add)
	fmt.Printf("The Sum is : %v\n", xi)

	//defer -> when a function holding a function ends then defer function gets executed
	//below, first bar gets executed then foo
	//It executes in Last-In-First-Out(LIFO) manner
	defer foo()
	bar("Sumit")

	h1 := Human{
		name: "Vaibhavi",
	}
	h2 := Human{
		name: "Sumit",
	}
	//approach with methods
	h1.speak()
	h2.speak()

	//INTERFACES & POLYMORPHISM
	//An interface is Go defines set of method signatures
	//Polymorphism is the ability of a VALUE of a certain TYPE to also be of another TYPE
	//In Go, values can be of more than one type
	g1 := Girl{
		Human: Human{
			name: "Rewa",
		},
		gender: "Female",
	}
	g1.speak()

	//approach with Interface
	saySomething(g1)
	// saySomething(h1)
	// saySomething(h2)

	//Stringer Interface -> implemented by any value that has a String method.
	//The String method is used to print the values passed as an operand to any format
	//that accepts a string or to an unformatted printer such as Print
	//syntax -> type Stringer interface {String() string}
	b := Book{
		title: "Mrutyunjay",
	}
	var n count = 36
	fmt.Println(b)
	fmt.Println(n)

	//Log Interface
	//approach using log
	log.Println(b)
	// log.Println(n)
	//approach using wrapper function
	logInfo(b)
	// logInfo(n)

	//Writer Interface -> wraps basic Write method
	//Writes len(p) bytes from p to the underlying data stream
	//Writing to a file
	// f, err := os.Create("output.txt")
	//Create creates or truncates the named file.
	//If the file already exists, it is truncated. If the file does not exist, it is created
	// if err != nil {
	// 	log.Fatalf("error: %s\n", err)
	// }
	// defer f.Close()

	//Writing to file with Write method
	// str := []byte("Hello World!") //converts string into byte slice for Write method
	// _, err = f.Write(str)
	// if err != nil {
	// 	log.Fatalf("error: %s\n", err)
	// }

	//Reading a file with Wrapper function
	xb, error := readFile("output.txt")
	if error != nil {
		log.Fatalf("readFile() in main: %s", error)
	}
	fmt.Println(xb)
	fmt.Println(string(xb))

	// Wrapper function example execution
	timeFunc(doWork)

	//BYTE BUFFER -> it is a region of memory used to temporarily store sequence of bytes
	// byt := bytes.NewBufferString("Hii")
	// fmt.Println(byt.String())
	// byt.WriteString(" there!")
	// fmt.Println(byt.String())
	// byt.Reset() // resets the buffer
	// byt.WriteString("Its's Hello World Now!")
	// fmt.Println(byt.String())
	// byt.Write([]byte(" Yay!!")) //Write method with byte slice as input
	// fmt.Println(byt.String())

	//Writing to a Byte Buffer
	// var bb bytes.Buffer
	// h1.writeOut(f)
	// h1.writeOut(&bb)
	// fmt.Println(bb.String())

	//Anonymous Function -> there is no receivers and identifiers
	// syntax -> func(parameters){code}(arguments)
	func() {
		fmt.Println("Anonymous Function executed")
	}()
	func(s string) {
		fmt.Println("This is anonymous function showing my name:", s)
	}("Vaibhavi")

	//Func Expressions -> assigning a function to expression and execute it
	//Expression -> combination of values, variables, operators & function calls
	//that are evaluated to produce a single value
	// a := func() {
	// 	fmt.Println("Anonymous Function with func expression executed")
	// }
	// c := func(s string) {
	// 	fmt.Println("This is anonymous function with func expression showing name:", s)
	// }
	// a()
	// c("Sumit")

	//Returning a Function
	fmt.Println(temp1())
	o := temp2()
	fmt.Println(o())
	// fmt.Printf("Type of temp1(): %T\n", temp1)
	// fmt.Printf("Type of temp2(): %T\n", temp2)
	// fmt.Printf("Type of o variable: %T\n", o)

	//CALLBACK -> Passing a function as an argument
	addition := doMath(3, 6, adding)
	fmt.Println(addition)
	subtraction := doMath(3, 6, subtract)
	fmt.Println(subtraction)

	//CLOSURE -> one scope enclosing other
	//variables declared in outer scope are accessible in inner scope
	//helps to limit the scope of variables
	f2 := incrementor()
	fmt.Println(f2()) //Call:1
	// fmt.Println(f2()) //Call:2
	// fmt.Println(f2()) //Call:3
	// fmt.Println(f2()) //Call:4
	// fmt.Println(f2()) //Call:5
	// fmt.Println(f2()) //Call:6

	// f3 := incrementor()
	// fmt.Println(f3()) //Call:1
	// fmt.Println(f3()) //Call:2
	// fmt.Println(f3()) //Call:3
	// fmt.Println(f3()) //Call:4
	// fmt.Println(f3()) //Call:5
	// fmt.Println(f3()) //Call:6

	//RECURSION
	fmt.Println(factorial(3))
}
