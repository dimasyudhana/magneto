package main

import (
	"log"
	"runtime"
)

func avg(nums []int, ch chan float64) {

	var sum int

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}

	ch <- float64(sum) / float64(len(nums))

}

func max(nums []int, ch chan int) {

	var max = nums[0]

	for i := 0; i < len(nums); i++ {
		if max < nums[i] {
			max = nums[i]
		}
	}

	ch <- max

}

func min(nums []int, ch chan int) {

	var min = nums[0]

	for i := 0; i < len(nums); i++ {
		if min > nums[i] {
			min = nums[i]
		}
	}

	ch <- min

}

func main() {

	runtime.GOMAXPROCS(2)

	var numbers = []int{3, 4, 3, 5, 6, 3, 2, 2, 6, 3, 4, 6, 3}
	log.Println(numbers)

	ch1 := make(chan float64)
	ch2 := make(chan int)
	ch3 := make(chan int)

	go avg(numbers, ch1)
	go max(numbers, ch2)
	go min(numbers, ch3)

	log.Printf("min:\t%d\n", <-ch3)

	for i := 0; i < 2; i++ {
		select {
		case avg := <-ch1:
			log.Printf("rerata:\t%.2f\n", avg)
		case max := <-ch2:
			log.Printf("maks:\t%d\n", max)
		}
	}
}
