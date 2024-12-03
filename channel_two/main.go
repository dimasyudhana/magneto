package main

import (
	"log"
	"runtime"
)

// go + channer = konkurensi yang sinkronus
func main() {

	runtime.GOMAXPROCS(2)

	msg := make(chan map[string]any)

	go func(arg map[string]any) {

		msg <- arg

	}(map[string]any{
		"field1": true,
		"field2": "false",
		"field3": 0,
		"field4": 100.0,
		"field5": byte(123),
	})

	for key, val := range <-msg {
		log.Printf("%s : %v", key, val)
		log.Printf("%s : %T", key, val)
	}

}
