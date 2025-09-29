package main 

import (
	"fmt"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	
	original := x
	reversed := 0
	
	for x != 0 {
		digit := x % 10
		reversed = reversed * 10 + digit
		x /= 10
	}
	
	return original == reversed
}

func main() {
	num := 121
	if isPalindrome(num) {
		fmt.Printf("%d is a palindrome number.\n", num)
	} else {
		fmt.Printf("%d is not a palindrome number.\n", num)
	}
}