package main

import (
	"fmt"
	"runtime"
	"sync" //sync is a package used for concurrency and topics like mutex
	"sync/atomic"
)

var wg sync.WaitGroup

func foo() {
	for i := 1; i <= 10; i++ {
		fmt.Println("foo() :", i)
	}
	wg.Done() //Done decrements the WaitGroup counter by one.
}

func bar() {
	for i := 1; i <= 10; i++ {
		fmt.Println("bar() :", i)
	}
}

//func init() -> Functions that runs only once at
//the beginning of program and set things up for main function

func main() {
	fmt.Println("Operating System :", runtime.GOOS)
	fmt.Println("Architecture :", runtime.GOARCH)
	fmt.Println("Number of CPUs :", runtime.NumCPU())
	fmt.Println("Num of GO Routines :", runtime.NumGoroutine())

	//GO Routine -> it is a function executing concurrently with other
	//GO routines in same address space
	//They are multiplexed onto multiple OS threads, so if one should block,
	//such as while waiting for I/O, others continue to run

	//The above print statements are already running one GO routine
	//But, writing go keyword in front of foo() launches a new GO routine

	// foo()
	//Typically this means the calls to Add should execute before the statement
	//creating the goroutine or other event to be waited for.
	wg.Add(1)
	go foo() //launching its own GO routine -> it makes the code concurrent but not running in parallel
	//go foo() -> foo() is not executed, gone directly to bar() function.
	//using wg methods, executes foo() function as well after bar()
	bar()

	fmt.Println("Number of CPUs :", runtime.NumCPU())
	fmt.Println("Num of GO Routines :", runtime.NumGoroutine())
	wg.Wait() //Wait blocks until the WaitGroup counter is zero.

	//Race Condition
	//Gosched() -> Yields the processors, allowing other goroutines to run. It doesn't suspend the
	//current goroutine, so execution resumes automatically.
	//type Duration int64 -> Represents elapsed time between two instants as an int64 nanosecond count
	fmt.Println("Number of CPUs :", runtime.NumCPU())
	fmt.Println("Number of GO Routines :", runtime.NumGoroutine())

	counter := 0

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	var mu sync.Mutex

	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock() //Lock locks m. If the lock is already in use, the calling goroutine blocks until the mutex is available.
			v := counter
			// time.Sleep(time.Second)
			runtime.Gosched() // same as above line of code
			v++
			counter = v
			mu.Unlock() //Unlock unlocks m. It is a run-time error if m is not locked on entry to Unlock.
			//A locked Mutex is not associated with a particular goroutine.
			//It is allowed for one goroutine to lock a Mutex and then arrange for another goroutine to unlock it.
			wg.Done()
		}()
		fmt.Println("Number of GO Routines :", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Number of GO Routines :", runtime.NumGoroutine())
	fmt.Println("Counter :", counter)

	//go help OR go help [command_name] -> documentation for commands used in execution
	//go run -race file_name.go -> gives data race conditions in GO program

	//To fix Race condition, we can use MUTEX
	//Mutex -> changes are done in above code itself using sync.Mutex

	//Atomic (sync.Atomic) -> The package provides low-level atomic memory primitives
	//useful for implementing synchronization algorithms.
	fmt.Println("Number of CPUs :", runtime.NumCPU())
	fmt.Println("Number of GO Routines :", runtime.NumGoroutine())

	var counter2 int64

	const gs2 = 30
	var wg2 sync.WaitGroup
	wg2.Add(gs2)

	for i := 0; i < gs2; i++ {
		go func() {
			atomic.AddInt64(&counter2, 1)
			runtime.Gosched() // same as above line of code
			fmt.Println("Counter :", atomic.LoadInt64(&counter2))
			wg2.Done()
		}()
		fmt.Println("Number of GO Routines :", runtime.NumGoroutine())
	}
	wg2.Wait()
	fmt.Println("Number of GO Routines :", runtime.NumGoroutine())
	fmt.Println("Counter :", counter)
}
