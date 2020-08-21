package channels

import (
	"fmt"
	"time"
)

// MakeArr...
func MakeArr(len int) (arr []int) {
	var a int
	for a = 0; a < len; a++ {
		arr = append(arr, a)
	}
	return
}

// sums items in an array together
func summer(num int, result chan int) {
	var total int

	var i int
	// returns the sum of the array
	for i = 0; i < num; i++ {
		total += i
	}
	// returning value to channel
	result <- total
}

// starts a go routine calculate the sum of items in individual arrays
func looper(number int) (result []int, sumChan int) {
	x := MakeArr(number)

	//making buffered channel
	chanresult := make(chan int, 1)
	var acc int

	// creating channels to calculate the total sum of each array value
	for acc = 0; acc < len(x); acc++ {
		go summer(x[acc], chanresult)
		x := <-chanresult
		result = append(result, x)
		sumChan += x
	}

	return
}

//TestChannel...
func TestChannel(testNumber int) {
	start := time.Now()
	_, sumOfArray := looper(testNumber)
	end := time.Since(start)

	fmt.Println("Sum of Arrays ", sumOfArray)
	fmt.Println("Function Execution took,", end)

}
