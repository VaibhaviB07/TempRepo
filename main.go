package main

import (
	"fmt"

	"github.com/VaibhaviB07/TempRepo/furtherexplored"
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
	//eg., furtherexplored.Fascinating() or furtherexplored.PlanetSpeed -> capitalized first letter means
	//visible/visited or exported outside the package

	//go commands
	//go module command : go mod init github.com/VaibhaviB07/repo_name
	//go run ./... -> runs all the files in folder
	//go build -> creates executable file -> ./file_name -> runs/executes the file
	//go env GOARCH GOOS -> shows all environment variables
	//run one of the followings at terminal to build to a certain OS :
	//GOOS=darwin go build
	//GOOS=linux go build
	//GOOS=windows go build
	//go install -> adds the go executable file in go/bin folder
	//go mod init module_name -> creates module
	//go mod tidy -> add missing and remove unused modules
	//go get github.com/VaibhaviB07/repo_name@latest -> fetches code from github and can be used in code

	//function/variable call from another package and file inside it where function/variable first name is capitalized
	furtherexplored.Fascinating()
	println(furtherexplored.PlanetSpeed)

	// git tag vN.N.N -> tags a certain commit (N is number)
	//eg., git tag v1.4
	//git show -> see tag data along with commit that was tagged eg., git show vN.N.N

}
