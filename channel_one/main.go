package main

import (
	"fmt"
	"log"
	"runtime"
)

// menggunakan channel berarti konkurensi berjalan secara sinkronus atau blocking untuk mengontrol aliran data
// tidak perlu menggunakan goroutine utama main selesai sebelum child_goroutine selesai
func do(arg string, msg chan string) {

	var data = fmt.Sprintf("text: %s", arg)
	msg <- data

}

func main() {

	runtime.GOMAXPROCS(2)

	msgChan := make(chan string)

	go do("one", msgChan)

	go do("two", msgChan)

	go do("three", msgChan)

	go do("four", msgChan)

	go do("five", msgChan)

	msg1 := <-msgChan

	log.Println(msg1)

	msg2 := <-msgChan

	log.Println(msg2)

	msg3 := <-msgChan

	log.Println(msg3)

	msg4 := <-msgChan

	log.Println(msg4)

	msg5 := <-msgChan

	log.Println(msg5)

	// var enter string
	// fmt.Scanln(&enter)

}
