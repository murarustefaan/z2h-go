package main

import (
	"fmt"
)

type Combo struct {
	original int
	squared  int
}

func generate(numbers []int) <-chan Combo {
	output := make(chan Combo)

	go func() {
		defer close(output)
		for _, v := range numbers {
			combo := Combo{original: v}
			output <- combo
		}
	}()

	return output
}

func square(input <-chan Combo) <-chan Combo {
	output := make(chan Combo)

	go func() {
		defer close(output)
		for v := range input {
			squared := v.original * v.original
			output <- Combo{original: v.original, squared: squared}
		}

	}()

	return output
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	input := generate(numbers)
	output := square(input)

	for v := range output {
		fmt.Printf("%d squared is %d\n", v.original, v.squared)
	}
}
