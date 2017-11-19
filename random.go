package main

import (
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var byteStore []byte

func getRandomByte() byte {
	// use up local random bytes first
	var storedByte byte
	if len(byteStore) > 0 {
		storedByte, byteStore = byteStore[0], byteStore[1:]
		return storedByte
	}

	// Before getting data, check available bits
	for {
		if checkQuota() >= 0 {
			break
		}
		time.Sleep(10 * time.Minute) // as recommended by Random.org
	}

	resp, err := http.Get("https://www.random.org/integers/?num=1000&min=0&max=255&col=1&base=10&format=plain&rnd=new")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	randomNumberBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	randomNumberStringArray := strings.Split(strings.Trim(string(randomNumberBytes), "\n"), "\n")

	for _, randomNumberString := range randomNumberStringArray {
		randomNumber, err := strconv.ParseUint(randomNumberString, 10, 8)
		if err != nil {
			panic(err)
		}
		byteStore = append(byteStore, uint8(randomNumber))
	}

	storedByte, byteStore = byteStore[0], byteStore[1:]
	return storedByte
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
