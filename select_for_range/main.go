package main

import (
	"fmt"
	"runtime"
)

/*
ch chan   string	Parameter ch untuk mengirim dan menerima data
ch chan<- string	Parameter ch hanya untuk mengirim data
ch <-chan string	Parameter ch hanya untuk menerima data
*/

func send(ch chan<- string) {

	for i := 0; i < 20; i++ {
		ch <- fmt.Sprintf("data %d", i)
	}

	close(ch)

}

func print(ch <-chan string) {
	for msg := range ch {
		fmt.Println(msg)
	}
}

func main() {

	runtime.GOMAXPROCS(2)

	var msg = make(chan string)
	go send(msg)
	print(msg)

}
