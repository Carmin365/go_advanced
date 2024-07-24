package main

import (
	"fmt"
	"time"
)

func to_send(channel chan int) {
	for i := 0; i < 10; i++ {
		channel <- i
	}
	close(channel)
}

func to_receive(channel chan int) {
	for number := range channel {
		fmt.Println(number)
	}
}

func main() {
	channel := make(chan int)

	go to_send(channel)
	go to_receive(channel)

	time.Sleep(time.Second * 2) // Aguarde a execução das goroutines
}
