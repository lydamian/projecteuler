package main

import "fmt"

func main() {
	multiple1, multiple2 := 3, 5
	sumMultiples := 0
	max := 1000

	
	for i := 1; i < max; i++ {
		if i % multiple1 == 0 || i % multiple2 == 0 {
			sumMultiples += i
		}
	}

	fmt.Println(sumMultiples)
}