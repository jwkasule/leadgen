package main

import (
	"log"
	"strconv"
)

func GetResultHandler(term, location, key string, radius int) {
	if len(key) != 39 {
		log.Fatalln("API Key isn't 39 characters, exiting...")
	}
	radiusString := strconv.Itoa(radius)
	GetResults(term, lcation, key, radiusString)
}
