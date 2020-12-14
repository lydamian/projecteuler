package main

import (
    "fmt"
	"os"
	"errors"
	"strconv"
)

func isPrime(num int) bool {
	if num < 2 {
		return false
	}
	upperBound := num/2
	for i := 2; i < upperBound; i++ {
		if isFactor, e := isFactor(num, i); isFactor && e == nil {
			return false
		}
	} 

	return true
}

func isFactor(num int, factor int) ( bool, error ) {
	if factor == 0 {
		return false, errors.New("0 is not a factor of any number")
	}

	return num % factor == 0, nil
}

func main() {
	number := 1
	var primeFactorization []int

	if len(os.Args) > 1 {
		number, _ = strconv.Atoi(os.Args[1])
	}

	for i := 1; i < number; i++ {
		if isFactor, e := isFactor(number, i); isPrime(i) && isFactor && e == nil {
			primeFactorization = append(primeFactorization, i)
		}
	}
	fmt.Println(number, "prime factorization =>", primeFactorization)
}