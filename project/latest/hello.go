package main

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"
)

type X struct {
	x int
	y int
}
type A struct {
	Reach    uint16
	NumPosts uint8
	NumLikes uint8
}

func Main() {
	var wg sync.WaitGroup
	p := new(int) // Allocate memory for an integer
	*p = 0        // Initialize the integer to 0
	rand.Float32()
	// Print the address of p and the process ID
	fmt.Printf("(%d) address pointed to by p: %p\n", os.Getpid(), p)

	// Use a WaitGroup to keep the main function alive
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			// Simulate the Spin function with a sleep
			time.Sleep(1 * time.Second)
			*p = *p + 1
			fmt.Printf("(%d) p: %d\n", os.Getpid(), *p)
		}
	}()

	// Wait for the goroutine to finish (it never will in this case)
	wg.Wait()
}

type B struct {
	X int64 // 8 bytes
	Y int32 // 4 bytes
}

func spread2(strs ...string) {
	for i := 0; i < len(strs); i++ {
		fmt.Println(strs[i])
	}
}
func prac() {
	for i := 0; i < 3; i++ {
		fmt.Println("Hello")
	}
}
func prac2() {
	for i := 0; i < 3; i++ {
		fmt.Println("World")
	}
}

type email struct {
	body string
	date time.Time
}

func checkEmailAge(emails [3]email) [3]bool {
	isOldChan := make(chan bool)

	go sendIsOld(isOldChan, emails)
	isOld := [3]bool{}
	time.Sleep(time.Second)
	isOld[0] = <-isOldChan
	isOld[1] = <-isOldChan
	isOld[2] = <-isOldChan
	return isOld
}

// don't touch below this line

func sendIsOld(isOldChan chan<- bool, emails [3]email) {
	for _, e := range emails {
		if e.date.Before(time.Date(2020, 0, 0, 0, 0, 0, 0, time.UTC)) {
			isOldChan <- true
		} else {
			isOldChan <- false
		}
		fmt.Println("Later printed")
	}
}

func testRoutine(ch chan<- int) {
	fmt.Println("Hi there")
	ch <- 1 // This will block until someone receives from ch
	fmt.Println("Hi there")
}

func concurrentFib(n int) (ans []int) {
	fibChannel := make(chan int)
	go fibonacci(n, fibChannel)
	for value := range fibChannel {
		ans = append(ans, value)
	}
	return
}

// don't touch below this line

func fibonacci(n int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
	}
	close(ch)
}

// func Main() {
// 	os.Getpid()
// ans := concurrentFib(50)
// fmt.Println(ans)
// var c = make(chan int, 100)

// close(c)
// c <- 1 // panic: send on closed channel
// test := [3]email{
// 	{
// 		body: "Words are pale shadows of forgotten names. As names have power, words have power.",
// 		date: time.Date(2019, 2, 0, 0, 0, 0, 0, time.UTC),
// 	},
// 	{
// 		body: "It's like everyone tells a story about themselves inside their own head.",
// 		date: time.Date(2021, 3, 1, 0, 0, 0, 0, time.UTC),
// 	},
// 	{
// 		body: "Bones mend. Regret stays with you forever.",
// 		date: time.Date(2022, 1, 2, 0, 0, 0, 0, time.UTC),
// 	},
// }
// a := checkEmailAge(test)
// fmt.Println(a)
// ch := make(chan int) // unbuffered channel
// go testRoutine(ch)
// fmt.Println("testing")
// value, ok := <-ch
// fmt.Println("Never execute")
// fmt.Println(value, ok)
// }

// func timePass() {
// 	panic("Wrong")
// }
// func Main() {

// 	// schedule, err := cron.ParseStandard("* * * * *")

// 	// fmt.Println(utils.Sum(1, 2))

// 	// cronExpiryTime := time.Unix(1723741300, 0).UTC()
// 	// c := time.Now().UTC().Before(cronExpiryTime)
// 	// fmt.Println("dd", c)
// 	// // x, y  := "hi", "nice"
// 	// // x, y = utils.Swap(x, y)
// 	// // var a, b int = 12, 12
// 	// // fmt.Println(a, b)
// 	// variable := 12
// 	// variable = 13
// 	// for i := 0; i < 5; i++ {
// 	// 	fmt.Println("Hi there")
// 	// }
// 	// b := 0
// 	// for b < 3 {
// 	// 	fmt.Println(b)
// 	// 	b += 1
// 	// }
// 	// os := "chirag"
// 	// switch os {
// 	// case "chirag":
// 	// 	fmt.Println("WTHIT")
// 	// }
// 	// fmt.Println(time.Now().Weekday().String())
// 	// fmt.Print(variable)
// 	// fmt.Println(&X{1, 2})
// 	// d := [5]int{1, 2, 3, 4, 5}
// 	// d = [...]int{100, 3: 400, 500}
// 	// ff := [...]int{100, 12: 12}
// 	// fmt.Println(ff)
// 	// fmt.Println(d)
// 	// s := make([]string, 4, 10)
// 	// s = append(s, "e", "f", "Ff")
// 	// fmt.Println("len:", s)
// 	// newMap := map[int]string{1: "sdfhl", 2: "pw734"}
// 	// fmt.Println(newMap[1])
// 	// for index, element := range d {
// 	// 	fmt.Println(index, element)
// 	// }
// 	// example := make(map[int]interface{})
// 	// fmt.Println(example)
// 	// example[1] = 1
// 	// example[2] = "shdfoh"
// 	// fmt.Println(example)

// 	// nextTime, _ := gronx.NextTickAfter("*/5 * * * *", time.Now().UTC(), false)
// 	// fmt.Println(nextTime.Unix())
// 	// cronExpr := "*/2* * * *" // Example: every day at midnight

// 	// // Parse the cron expression
// 	// schedule, err := cron.ParseStandard(cronExpr)
// 	// if err != nil {
// 	// 	fmt.Println("Error parsing cron expression:", err)
// 	// 	return
// 	// }

// 	// // Get the current time
// 	// now := time.Now()

// 	// // Find the next scheduled time
// 	// nextTime := schedule.Next(now)
// 	// fmt.Println("Next Scheduled Time:", time.Now().UTC())
// 	// fmt.Println("CURrent Scheduled Time:", time.Now())
// 	// fmt.Println("final Scheduled Time:", nextTime)

// 	defer func() {
// 		if r := recover(); r != nil {
// 			fmt.Println("recovered", r)
// 		}
// 	}()
// 	timePass()
// ch <-1
// 	fmt.Println("Size of A:", unsafe.Sizeof(A{})) // Size of A
// 	fmt.Println("Size of B:", unsafe.Sizeof(B{})) // Size of B

// }
