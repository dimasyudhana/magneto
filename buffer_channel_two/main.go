package main

import (
	"log"
	"runtime"
	"time"
)

func do(arg map[string]any, msg chan map[string]any) {
	msg <- arg
}

func main() {

	runtime.GOMAXPROCS(2)

	start := time.Now()

	var args = map[string]any{
		"field1":  100.0,
		"field2":  "99.0",
		"field3":  true,
		"field4":  100.0,
		"field5":  "99.0",
		"field6":  true,
		"field7":  100.0,
		"field8":  "99.0",
		"field9":  true,
		"field10": 100.0,
	}

	var msg = make(chan map[string]any, len(args))

	for key, val := range args {

		go func(key string, val any, msg chan map[string]any) {

			arg := map[string]any{
				key: val,
			}

			do(arg, msg)

		}(key, val, msg)

	}

	for i := 0; i < len(args); i++ {

		data := <-msg

		for key, val := range data {

			log.Printf("%s : %v", key, val)

		}

	}

	log.Printf("Concurrent execution time: %v", time.Since(start))

}
