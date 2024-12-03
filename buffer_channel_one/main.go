package main

import (
	"fmt"
	"runtime"
)

// goroutine + channel + buffer channel = konkurensi secara asinkronus
// buffer channel menjadi penentu jumlah data yang bisa dikirim bersamaan
// jika jumlah data melebihi buffer channel maka proses akan menunggu hingga ada channel yang kosong
func main() {

	runtime.GOMAXPROCS(2)

	msg := make(chan int, 2)

	go func() {
		for {
			fmt.Println("terima data: ", <-msg)
		}
	}()

	for i := 0; i < 5; i++ {
		fmt.Println("kirim data: ", i)
		msg <- i
	}

	var print string
	fmt.Scanln(&print)
	// atau
	// time.sleep(1 * time.Second)

}
