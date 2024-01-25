package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Niladic function -> function/program having no arguments
// init function is a Niladic function
// function init can run before main
func init() {
	fmt.Println("Before main function, executes init function")
}

// main function is the one primary Go routine that runs every time we launch a Go program
func main() {
	//Sequence -> statements execute in sequential order
	fmt.Println("Hello World!!")
	fmt.Println("Vaibhavi Here!!")
	var age int = 18
	num := 80

	//% -> modulus operator -> returns remainder value of division

	//Conditional statements
	if age >= 18 {
		fmt.Println("You are allowed to enter")
	} else if age < 0 {
		fmt.Println("Not a valid age")
	} else {
		fmt.Println("You are not allowed to enter")
	}

	//Logical Operators -> &&-AND ||-OR !-NOT
	if num%2 == 0 && num > 0 {
		fmt.Printf("The number %v is Positive and Even\n", num)
	}

	//Statement; Statement Idiom -> in an if statement, the expression may be preceded
	//by a simple statement, which executes before the expression is evaluated
	//here, scope of z variable is only limited to this if-else block
	if z := 3 * rand.Intn(num); z >= num {
		fmt.Printf("z is %v that is GREATER THAN OR EQUAL to x which is %v\n", z, num)
	} else {
		fmt.Printf("z is %v that is LESS THAN x which is %v\n", z, num)
	}

	//Switch Statements
	const day = 4
	switch {
	case day < 1:
		fmt.Println("Can't be a day of week")
	case day > 7:
		fmt.Println("Can't be a day of week")
	case day == 7:
		fmt.Println("Its a Holiday")
	default:
		fmt.Println("Its a normal day of week")
	}

	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Not a valid day")
	}

	//fallthrough keyword -> follows to the next case and executes it as well with the matching case
	x := 40
	switch x {
	case 40:
		fmt.Println("x = 40")
		fallthrough
	case 41:
		fmt.Println("Printed/Executed because of fallthrough : x=41")
		fallthrough
	case 42:
		fmt.Println("Printed/Executed because of fallthrough : x=42")
	}

	//Concurrency -> select a channel
	ch1 := make(chan int)
	ch2 := make(chan int)

	d1 := time.Duration(rand.Int63n(250))
	d2 := time.Duration(rand.Int63n(250))
	// fmt.Printf("%v\t%T\n", d1, d2)
	// time.Sleep(d1 * time.Millisecond)
	// fmt.Println("Done")

	//Go Routines
	//below go routines push values such as 41, 42 on channel
	go func() {
		time.Sleep(d1 * time.Millisecond)
		ch1 <- 41
	}()
	go func() {
		time.Sleep(d2 * time.Millisecond)
		ch2 <- 42
	}()

	//Select Statements
	//chooses a which of a set of of possible send or receive operations will proceed
	//it looks similar to switch statement but with cases all referring to communication operations
	select {
	case v1 := <-ch1:
		fmt.Println("Value from Channel 1 ", v1)
	case v2 := <-ch2:
		fmt.Println("Value from Channel 2 ", v2)
	}

	//Loop statements
	for i := 0; i < 5; i++ {
		fmt.Printf("Counting 1 : %v\n", i)
	}

	//Like While loop in C
	// y := 1
	// for y < 5 {
	// 	fmt.Printf("Counting 2 : %v\n", y)
	// 	y++
	// }

	//break keyword
	// for {
	// 	fmt.Printf("Counting 3 : %v\n", y)
	// 	if y > 10 {
	// 		break
	// 	}
	// 	y++
	// }

	//continue keyword
	// for i := 10; i >= 1; i-- {
	// 	if i%2 == 0 {
	// 		continue
	// 	}
	// 	fmt.Println("Odd numbers descending order :", i)
	// }

	//Nested Loop
	// for i := 0; i < 5; i++ {
	// 	fmt.Println("--")
	// 	for j := 0; j < 5; j++ {
	// 		fmt.Printf("Outer Loop %v, Inner Loop %v\n", i, j)
	// 	}
	// }

	//for range loops -> can also used on string as a slicing tool
	//Slice data structure
	// arr := []int{1, 2, 3, 4, 5}
	// for i, v := range arr {
	// 	fmt.Println("Ranging over a Slice", i, v)
	// }
	//Map data structure
	// m := map[string]int{
	// 	"Vaibhavi": 20,
	// 	"Sumit":    26,
	// }
	// for k, v := range m {
	// 	fmt.Println("Ranging over a Map", k, v)
	// }

	//ARRAY
	//Generally, instead of array, slice is used in Go since it is dynamic approach of array
	var arr [3]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	// for i, v := range arr {
	// 	fmt.Printf("arr[%v] = %v\n", i, v)
	// }

	//{} -> composite literal
	//array literal
	// arr_new := [6]int{2, 4, 6, 8, 10, 12}
	// fmt.Printf("arr_new : %#v\n", arr_new)
	// %#v -> prints whole arr with data type and size as per the declaration syntax

	//len(arr_name) -> length of array/slice
	//cap(arr_name) -> capacity of slice

	// [...] -> compiler figure outs the size of array
	// str_arr := [...]string{"Vaibhavi", "Rewa", "Dhanashree"}
	// fmt.Printf("str_arr : %#v, Type : %T, Size : %v\n", str_arr, str_arr, len(str_arr))

	//SLICE
	slice := []string{"Hello", "World"}
	fmt.Println(slice)
	// fmt.Printf("str_arr : %#v, Type : %T, Size : %v\n", slice, slice, len(slice))
	//blank identifier not to use the returned value -> used when we don't want the index
	// for _, v := range slice {
	// 	fmt.Printf("Value - %v\n", v)
	// }
	//access by index position
	// for i := 0; i < len(slice); i++ {
	// 	fmt.Printf("Value : %v\n", slice[i])
	// }

	//Append to a slice -> arr_name = append(slice_name, elements_to_append_separated_by_comma)
	//Variadic parameter
	slice = append(slice, "From", "Vaibhavi")
	fmt.Println(slice)

	//Slicing a Slice
	//[inclusive:exclusive]
	// fmt.Println(slice[0:4])
	//[:exclusive]
	// fmt.Println(slice[:3])
	//[inclusive:]
	// fmt.Println(slice[3:])
	//[:] - complete slice
	// fmt.Println(slice[:])

	//Deleting from slice -> using append and slicing
	//slice_name[:]... -> open ups the slice as separate elements
	//x := append(x, y...) -> append all elements of slice y to slice x
	// slice = append(slice[:2], slice[3:]...)
	// slice = append(slice[:1], slice[2:]...)
	// fmt.Println("Deleted \"World\" from slice : ", slice)

	//make function -> to specify how large our slice should be
	//and also how large the underlying array should be
	//syntax -> slice_name = make([]data_type, length, capacity)
	//if in make() function, length is not mentioned it creates slice of capacity and assign all values to 0
	si := make([]int, 0, 6) //length is not mentioned -> empty slice with length/size = 0
	// fmt.Println(si)
	// fmt.Println("Length of si : ", len(si))
	// fmt.Println("Capacity of si : ", cap(si))
	si = append(si, -1, -2, -3, -4, -5)
	fmt.Println(si)
	// fmt.Println("Length of si after adding elements : ", len(si))
	// fmt.Println("Capacity of si after adding elements : ", cap(si))

	//Multidimensional Slice
	//syntax -> slice_name = [][]data_type[slice_name1, slice_name2, ...]
	me := []string{"Vaibhavi", "Bhosale", "20"}
	bro := []string{"Sumit", "Bhosale", "26"}
	// fmt.Println(me)
	// fmt.Println(bro)
	siblings := [][]string{me, bro}
	fmt.Println("Multidimensional Slice : ", siblings)
	// fmt.Println(siblings[0][1])

	//Copying a Slice
	//Changes made in old array will also get updated in copied array
	a := []int{1, 2, 3}
	// b := a
	// fmt.Println(a, b)
	// a[0] = 4
	// fmt.Println(a, b)

	//Copying a Slice -> with copy() function
	b := make([]int, 3)
	copy(b, a) //here, changes in a won't affect the copy b
	a[0] = 4
	fmt.Println(a, b)

	//Functions having slices are parameter/arguments acts as call by reference functions
	//and works on real slice that is passed

	//MAP
	roll := map[string]int{
		"Vaibhavi": 40,
		"Manasi":   34,
		"Rishita":  49,
		"Gauri":    62,
	}
	// fmt.Printf("Roll No. of Vaibhavi is : %v\n", roll["Vaibhavi"])
	// fmt.Println(roll)
	fmt.Printf("%#v\n", roll)

	//make function on map -> Adding elements to map
	//map can  print data in unordered manner
	ag := make(map[string]int)
	ag["Vaibhavi"] = 20
	ag["Sumit"] = 26
	// fmt.Println(len(ag))
	fmt.Printf("%#v\n", ag)

	for k, v := range roll {
		fmt.Println(k, v)
	}

	//Deleting element from map -> delete(map_name, key_name_to_delete) function
	delete(ag, "Sumit")
	fmt.Printf("%#v\n", ag)
	fmt.Println(ag["Sumit"])

	//Map - comma ok idiom
	//If you look up a non-existent key, the zero value will be returned as the value
	//associated with that non-existent key
	//You can check to see if you have looked up an existing key, or a non-existent
	//key using comma ok idiom
	if v, ok := ag["Vaibhavi"]; !ok {
		fmt.Println("Key does not exist")
	} else {
		fmt.Printf("The value prints as, %v\n", v)
	}

	//map[string][]string -> map with Keys as string and Values as Slice of string
}
