package channels

import (
	"fmt"
	"time"
)

//MakeArr...
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

	for i = 0; i < num; i++ {
		total += i
	}

	result <- total
}

// starts a go routine calculate the sum of items in individual arrays
func looper(number int) (result []int, sumChan int) {
	x := MakeArr(number)

	//making buffered channel
	chanresult := make(chan int, 1)
	var acc int

	for acc = 0; acc < len(x); acc++ {
		go summer(x[acc], chanresult)
		x := <-chanresult
		result = append(result, x)
		sumChan += x
	}

	return
}

func TestChannel(testNumber int) {
	start := time.Now()
	_, sumOfArray := looper(testNumber)
	end := time.Since(start)

	fmt.Println("Sum of Arrays ", sumOfArray)
	fmt.Println("Function Execution took,", end)

}
