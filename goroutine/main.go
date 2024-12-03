package main

import (
	"fmt"
	"log"
	"runtime"
)

func do(idx int, msg string) {

	for i := 0; i < idx; i++ {
		log.Println(msg)
	}

}

func main() {

	runtime.GOMAXPROCS(2)

	go do(5, "est.1996")
	do(5, "qwerty")

	// test hasil supaya tidak mati duluan.
	var print string
	fmt.Scanln(&print)

}
