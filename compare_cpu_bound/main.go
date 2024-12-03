package main

import (
	"log"
	"time"
)

func main() {

	start := time.Now()

	for key, val := range map[string]any{
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
	} {
		arg := map[string]any{
			key: val,
		}
		data := arg
		for key, val := range data {
			log.Printf("%s : %v", key, val)
		}
	}

	log.Printf("Sequential execution time: %v", time.Since(start))

}
