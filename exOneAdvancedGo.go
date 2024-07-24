package main

import (
	"fmt"
	"time"
)

func stamp(n int) {
	fmt.Println(n)
}

func main() {
	for i := 0; i < 10; i++ {
		go stamp(i)
	}
	time.Sleep(time.Second * 2) // Aguarde a execução das goroutines
}
