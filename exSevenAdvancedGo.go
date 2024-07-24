package main

import (
	"fmt"
)

func main() {

	channelNumbers := make(chan int)

	go func(channel chan int) {
		for number := range channel {
			if isCousin(number) {
				fmt.Println(number)
			}
		}
	}(channelNumbers)

	for i := 1; i <= 100; i++ {
		channelNumbers <- i
	}

	close(channelNumbers)
}

func isCousin(number int) bool {
	if number <= 1 {
		return false
	}
	for i := 2; i*i <= number; i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}
