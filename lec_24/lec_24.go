//JSON, BCRYPT, SORT PACKAGES

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sort"

	"golang.org/x/crypto/bcrypt"
)

type Person struct {
	First string
	Last  string
	Age   int
}

type PersonNew struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

func (p Person) String() string {
	return fmt.Sprintf("%s %s: %d\n", p.First, p.Last, p.Age)
}

// Customized Sort
type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

type ByName []Person

func (bn ByName) Len() int           { return len(bn) }
func (bn ByName) Swap(i, j int)      { bn[i], bn[j] = bn[j], bn[i] }
func (bn ByName) Less(i, j int) bool { return bn[i].First < bn[j].First }

func main() {
	p1 := Person{
		First: "Vaibhavi",
		Last:  "Bhosale",
		Age:   20,
	}
	p2 := Person{
		First: "Sumit",
		Last:  "Bhosale",
		Age:   27,
	}
	siblings := []Person{p1, p2}

	fmt.Println(siblings)

	//Marshal JSON
	bs, err := json.Marshal(siblings) //returns byte slice in json format
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println(string(bs)) //convert byte slice to string
	//to execute Marshal method, the first letters of attributes in struct must be in Capital

	//Unmarshal JSON
	json_str := `[{"First":"Vaibhavi","Last":"Bhosale","Age":20},{"First":"Sumit","Last":"Bhosale","Age":27}]`
	byte_slice := []byte(json_str)
	fmt.Printf("%T\n", json_str)
	fmt.Printf("%T\n", byte_slice)

	//siblings_new := []PersonNew{} -> same as below line
	var siblings_new []PersonNew //for json to string/slice -> this is mre authentic
	// := -> declaration , = -> assigning/reassigning

	err = json.Unmarshal(byte_slice, &siblings_new)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Println("All Data : ")
	for i, v := range siblings_new {
		fmt.Println("Person Number:", i)
		fmt.Println(v.First, v.Last, v.Age)
	}

	//Decode function -> reads next JSOn-encoded value from its input and
	//stores it in the value pointed by v

	fmt.Println(siblings)
	err = json.NewEncoder(os.Stdout).Encode(siblings)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	//WRITER INTERFACE -> wraps basic write method
	//READER INTERFACE -> wraps basic read method

	fmt.Fprintln(os.Stdout, "Hello World!")     //equivalent to fmt.Println("Hello World!")
	io.WriteString(os.Stdout, "Hello World!\n") //equivalent to above

	//Sorting slice
	s := []int{4, 9, -2, 0, 2, 5, 99, -78}
	// s := []string{"Z", "A", "M", "H"}
	fmt.Println("Unsorted:", s)
	sort.Ints(s)
	// sort.Strings(s)
	fmt.Println("Sorted:", s)

	//Customized sort implementation
	fmt.Println(siblings)
	//sort.Sort() -> Sort sorts data. It makes one call to data.Len to determine n,
	//and O(n * log(n)) calls to data.Less and data.Swap. The sort is not guaranteed to be stable.
	sort.Sort(ByAge(siblings))
	fmt.Println(siblings)
	sort.Sort(ByName(siblings))
	fmt.Println(siblings)

	//encrypting and decrypting passwords
	//constants for level of encryption
	//MinCost = 4, MaxCost = 31, DefaultCost = 10
	p := `password1234`
	bs_new, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("Password: ", p)
	fmt.Println("Encrypted password: ", bs_new)

	err = bcrypt.CompareHashAndPassword(bs_new, []byte(p))
	//err = bcrypt.CompareHashAndPassword(bs_new, []byte(`password`))
	if err != nil {
		fmt.Println("YOU CANNOT LOGIN!")
		return
	}
	fmt.Println("Login Successful!")
}
