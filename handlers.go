package main

import (
	"log"
	"strconv"
	"strings"
)

func GetResultHandler(query, location, radius, key string) {
	if len(key) != 39 {
		PrintHelp()
		log.Fatalln("API_KEY env variable isn't 39 characters, please use a proper key...")
	}
	locationsplit := strings.Split(location, ",")
	if len(locationsplit) != 2 {
		log.Fatalln("Bad location format, please enter a string in format \"lon,lat\" as floats including quotations")
	}
	lat, err := strconv.ParseFloat(strings.TrimSpace(locationsplit[0]), 10)
	if err != nil {
		log.Fatalln("Latitude doesn't parse, please enter a float")
	}
	long, err := strconv.ParseFloat(strings.TrimSpace(locationsplit[1]), 10)
	if err != nil {
		log.Fatalln("Latitude doesn't parse, please enter a float")
	}
	radiusInt, err := strconv.Atoi(radius)
	if err != nil {
		PrintHelp()
		log.Fatalln("Radius is not an integer, please enter an integer...")
	}
	// Yeah, this takes the uint type specifically. Why? uint itself can't even
	// be shit out with strconv.ParseUInt (Just uint8-64), and the := assignment
	// doesn't work, but we CAN do this fuckery and it just werks. Weird choice
	// on the part of the authors of this lib, but hey that's why we got this
	// handler function in any case -- to make sure we got validated types in line!
	var radiusUInt = uint(radiusInt)
	GetResults(query, lat, long, radiusUInt, key)
}
