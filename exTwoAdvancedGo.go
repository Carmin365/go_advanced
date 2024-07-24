package main

import (
	"fmt"
	"time"
)

func print_out(n int) {
	fmt.Println(n)
}

func main() {
	for i := 0; i < 10; i++ {
		go print_out(i)
	}
	time.Sleep(time.Second * 2) // Aguarde a execução das goroutines
}
