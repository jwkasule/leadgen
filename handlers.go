package main

import (
	"log"
	"strconv"
	"strings"
)

// GetResultHandler takes a filename, query, location, radius, and api key,
// passes this data over to the google APIs to get the data, then passes
// the resulting data from that over to WriteCSV
func GetResultHandler(filename, query, location, radius, key string) {
	// api key check
	if len(key) != 39 {
		PrintHelp()
		log.Fatalln("API_KEY env variable isn't 39 characters, please use a proper key...")
	}
	// location checks
	locationsplit := strings.Split(location, ",")
	if len(locationsplit) != 2 {
		PrintHelp()
		log.Fatalln("Bad location format, please enter a string in format \"lat,long\" as floats including quotations")
	}
	lat, err := strconv.ParseFloat(strings.TrimSpace(locationsplit[0]), 64)
	if err != nil {
		PrintHelp()
		log.Fatalln("Latitude doesn't parse, please enter latitude as a float")
	}
	long, err := strconv.ParseFloat(strings.TrimSpace(locationsplit[1]), 64)
	if err != nil {
		PrintHelp()
		log.Fatalln("Longitude doesn't parse, please enter a longitude as a float")
	}
	// radius check
	radiusInt, err := strconv.Atoi(radius)
	if err != nil {
		PrintHelp()
		log.Fatalln("Radius is not an integer, please enter an integer...")
	}
	if radiusInt > 50000 {
		PrintHelp()
		log.Fatalln("Radius is over 50000, please enter a radius under 50000 meters/50km")
	}
	// Yeah, this takes the uint type specifically. Why? uint itself can't
	// even be shit out with strconv.ParseUInt (Just uint8-64), and the
	// := assignment doesn't work, but we CAN do this fuckery and it just
	// werks. Weird choice on the part of the authors of this lib, but hey
	// that's why we got this handler function in any case -- to make sure we
	// got validated types in line!
	var radiusUInt = uint(radiusInt)
	// We get an array of string arrays
	results := GetResults(query, lat, long, radiusUInt, key)
	WriteCSV(filename, results)
}
