package main

import (
	"fmt"
	"time"
)

func stamp(n int) {
	fmt.Println(n)
}

func main() {
	for y := 0; y < 10; y++ {
		go stamp(y)
	}
	time.Sleep(time.Second * 2) 
}
