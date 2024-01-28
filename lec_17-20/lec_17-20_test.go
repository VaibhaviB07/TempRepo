package main

import "testing"

//TESTING in Go
//File name : <file_name_want_to_test_on>_test.go
//Function is which testing part is implemented : Test<any_name_you_want>
//Unit Test -> testing a smallest testable part of a software

func TestAdd(t *testing.T) {
	total := adding(6, 9)
	if total != 15 {
		t.Errorf("Sum is Incorrect. Want: %d, Got: %d\n", 15, total)
	}
}
