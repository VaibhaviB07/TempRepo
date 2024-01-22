package main

import (
	"fmt"
)

var x int = 9 // package scope declaration

type person struct {
	name string
}

func sayMyName1(p person) {
	fmt.Println("Hello, ", p.name)
}

func (p person) sayMyName2() {
	fmt.Println("Hello, ", p.name)
}

func main() {
	var y int = 12 // block scope declaration
	fmt.Println(x, y)
	x := 10 //variable shadowing
	fmt.Println(x)

	p1 := person{
		"Vaibhavi Bhosale",
	}
	sayMyName1(p1)  //function call
	p1.sayMyName2() //method call

	//the function/variable which we have to use in whole folder/package can be written as first letter capital
	//syntax => package_name.function_name
	//eg., furtherexplored.Fascinating() or furtherexplored.PlanetSpeed
}
