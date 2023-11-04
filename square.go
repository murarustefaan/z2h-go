package main

import (
	"fmt"
	"os"
	"strconv"
)

type Combo struct {
	original int
	squared  int
}

func generate(count int) <-chan int {
	output := make(chan int)

	go func() {
		defer close(output)
		for i := 1; i <= count; i++ {
			output <- i
		}
	}()

	return output
}

func create(input <-chan int) <-chan Combo {
	output := make(chan Combo)

	go func() {
		defer close(output)
		for v := range input {
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
	count := 100

	fmt.Println(os.Args)
	if len(os.Args) >= 2 {
		parsed, err := strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Printf("Invalid argument, defaulting to %d\n", count)
		}
		if parsed > 0 {
			count = parsed
		}
	}

	input := generate(count)
	mapped := create(input)
	output := square(mapped)

	for val := range output {
		fmt.Printf("%d squared is %d\n", val.original, val.squared)
	}
}
