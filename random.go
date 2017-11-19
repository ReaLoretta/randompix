package main

import (
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

func checkQuota() int64 {
	resp, err := http.Get("https://www.random.org/quota/?format=plain")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	quotaBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	quotaString := strings.Trim(string(quotaBytes), "\n")

	quota, err := strconv.ParseInt(quotaString, 10, 64)
	if err != nil {
		panic(err)
	}

	return quota
}
