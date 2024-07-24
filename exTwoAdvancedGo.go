package main

import (
	"fmt"
	"time"
)

func print_out(n int) {
	fmt.Println(n)
}

func main() {
	for y := 0; y < 10; y++ {
		go print_out(y)
	}
	time.Sleep(time.Second * 2) 
}
