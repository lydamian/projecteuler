package main

import "fmt"

func getFibValues(max int) []int {
	fibValues := []int{1,2}
	fibValue := fibValues[len(fibValues) - 1] + fibValues[len(fibValues) - 2]
	for fibValue < max {
		fibValues = append(fibValues, fibValue)
		fibValue = fibValues[len(fibValues) - 1] + fibValues[len(fibValues) - 2]
	}
	return fibValues
}

func sumEvenValues(mySlice []int) int {
	sum := 0
	for _, v := range mySlice {
		if v % 2 == 0 {
			sum += v
		}
	}
	return sum
}

func main() {
	fibValues := getFibValues(4000000)
	fmt.Println(fibValues)
	fibSum := sumEvenValues(fibValues)
	fmt.Println(fibSum)
}