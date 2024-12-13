package main

import (
	"log"
	"math/rand"
	"runtime"
	"time"
)

func send(ch chan<- int) {

	randomizer := rand.New(rand.NewSource(time.Now().Unix()))

	for i := 0; true; i++ {
		ch <- i
		time.Sleep(time.Duration(randomizer.Int()%10+1) * time.Second)
	}

}

func get(ch <-chan int) {
loop:
	for {
		select {
		case data := <-ch:
			log.Printf("%d", data)
		case <-time.After(time.Second * 5):
			log.Println("selesai, 5 detik no activities")
			break loop
		}
	}
}

func main() {

	runtime.GOMAXPROCS(2)

	var msg = make(chan int)

	go send(msg)
	get(msg)

}
