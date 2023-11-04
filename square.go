package main

import (
	"fmt"
)

func generate(numbers []int) <-chan int {
	output := make(chan int)

	go func() {
		defer close(output)
		for _, v := range numbers {
			output <- v
		}
	}()

	return output
}

func square(input <-chan int) <-chan int {
	output := make(chan int)

	go func() {
		defer close(output)
		for v := range input {
			output <- v * v
		}

	}()

	return output
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	input := generate(numbers)
	output := square(input)

	for v := range output {
		fmt.Printf("%d ", v)
	}
}
