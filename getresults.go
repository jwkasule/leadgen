package main

import (
	"log"

	"googlemaps.github.io/maps"
)

func GetResults(term, key, location, radius string) {
	c, err := maps.NewClient(maps.WithAPIKey(key))
	if err != nil {
		log.Fatalf("Error in initializing maps client: %s", err)
	}
}
