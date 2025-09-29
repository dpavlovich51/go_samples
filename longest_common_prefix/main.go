package main

import (
	"fmt"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	var prefix = ""

	for i := 0; i < len(strs[0]); i++ {
		char := strs[0][i]
		for j := 1; j < len(strs); j++ {
			if i >= len(strs[j]) || strs[j][i] != char {
				return prefix
			}
		}
		prefix += string(char)
	}
	
	return prefix
}

func main() {
	strs := []string{"flower", "flow", "flight"}
	prefix := longestCommonPrefix(strs)
	fmt.Printf("The longest common prefix is: %s\n", prefix)
}
