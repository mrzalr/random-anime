package main

import (
	"math/rand"
	"time"
)

var limit = 2000

func Init() {
	rand.NewSource(time.Now().UnixNano())
}

func GetRandomNumber(min, max int) int {
	return min + rand.Intn(max-min+1)
}

func GetRandomPage() int {
	maxNum := limit / 50

	randomOffset := GetRandomNumber(0, maxNum)
	return randomOffset * 50
}
