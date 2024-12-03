package main

import (
	"log"
	"runtime"
)

func do(arg map[string]any, msg chan map[string]any) {
	msg <- arg
}

func logging(arg chan map[string]any) {
	data := <-arg
	for key, val := range data {
		log.Printf("%s : %v", key, val)
	}
}

// channel lewat argument bersifat pass by reference, yang di kirim adalah pointer datanya, bukan nilainya

func main() {

	runtime.GOMAXPROCS(2)

	var args = map[string]any{
		"field1": 100.0,
		"field2": "99.0",
		"field3": true,
	}

	var msg = make(chan map[string]any)

	for key, val := range args {

		go func(key string, val any, msg chan map[string]any) {

			row := map[string]any{
				key: val,
			}

			do(row, msg)

		}(key, val, msg)

	}

	for i := 0; i < len(args); i++ {

		logging(msg)

	}

}
