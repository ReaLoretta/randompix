package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func getRandomByte() byte {
	resp, err := http.Get("https://www.random.org/integers/?num=1&min=0&max=255&col=1&base=10&format=plain&rnd=new")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	randomNumberBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	randomNumberString := strings.Trim(string(randomNumberBytes), "\n")

	randomNumber, err := strconv.ParseUint(randomNumberString, 10, 8)
	if err != nil {
		panic(err)
	}

	return uint8(randomNumber)
}

func main() {
	var output string
	for i := 0; i < 10; i++ {
		output += string(getRandomByte())
	}
	fmt.Println(output)
}
