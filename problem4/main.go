package main

import (
    "fmt"
)

func isPalindrome(num int) bool {
    var reversedNum int = 0

    tmpNum:= num

    for {
        remainder := tmpNum % 10

        reversedNum = reversedNum * 10 + remainder

        tmpNum = tmpNum / 10

       if (tmpNum == 0) {
           break
       }

    }

    if reversedNum == num {
		return true
	}
	return false
    
}

func main() {
	max := 0
	for num1 := 999; num1 >= 100; num1-- {
		for num2 := 999; num1 >= 100; num2-- {
			if isPalindrome(num1 * num2) {
				if num1 * num2 > max {
					max = num1 * num2
				}
			}
		}
		
	}

	fmt.Println(max)
}