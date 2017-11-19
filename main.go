package main

import (
	"fmt"
)

func main() {
	var output string
	for i := 0; i < 10; i++ {
		output += string(getRandomByte())
	}
	fmt.Println(output)
	fmt.Println(checkQuota())
}
