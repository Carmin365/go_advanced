package main

import (
	"bufio"
	"os"
	"sync"
)

func main() {

	inputFile, err := os.Open("prohibited.txt")
	if err != nil {
		panic(err)
	}
	defer inputFile.Close()

	channelLines := make(chan string)

	go func() {
		scanner := bufio.NewScanner(inputFile)
		for scanner.Scan() {
			line := scanner.Text()
			channelLines <- line
		}
		if err := scanner.Err(); err != nil {
			panic(err)
		}
		close(channelLines)
	}()

	outputFile, err := os.Create("saida.txt")
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()

		write := bufio.NewWriter(outputFile)
		for line := range channelLines {
			_, err := write.WriteString(line + "\n")
			if err != nil {
				panic(err)
			}
		}
		write.Flush()
	}()

	wg.Wait()
}
