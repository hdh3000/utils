package main

import "fmt"

func main() {
	input := []string{"foo","bar","baz", "bing"}
	for i := len(input)-1; i > 0; i-- {
		fmt.Println(input[:i], input[i])
	}

	for i := 0; i < len(input); i++ {
		fmt.Println(input[i:])
	}
}

