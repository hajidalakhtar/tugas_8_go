package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	runtime.GOMAXPROCS(2)
	var message = make(chan string)

	go kirimdata(message)
	terimadata(message)

}
func kirimdata(ch chan<- string) {
	for i := 0; true; i++ {
		ch <- "Apa kabar teman teman"
		time.Sleep(time.Duration(rand.Int()%10+1) * time.Second)
	}
}

func terimadata(ch <-chan string) {
loop:
	for {
		select {
		case data := <-ch:
			fmt.Print("Terima data ", data, "\n")
		case <-time.After(time.Second * 5):
			fmt.Println("tidak ada aktifitas dalam 5 detik")
			break loop

		}
	}
}
