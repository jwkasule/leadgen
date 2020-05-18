package main

import (
	"context"
	"fmt"
	"log"

	"googlemaps.github.io/maps"
)

func GetResults(query string, lat, long float64, radius uint, key string) {
	client, err := maps.NewClient(maps.WithAPIKey(key))
	if err != nil {
		log.Fatalf("Error in initializing maps client: %s", err)
	}
	// Are we really going to use this shitty lib?
	// Read more to find out!
	location := &maps.LatLng{
		Lat: lat,
		Lng: long,
	}
	req := &maps.TextSearchRequest{
		Location: location,
		Radius:   radius,
		Query:    query,
	}
	resp, err := client.TextSearch(context.Background(), req)
	if err != nil {
		log.Fatalf("Error in placing text search request: %s", err)
	}
	// fmt.Printf("%+v\n", resp)
	for _, result := range resp.Results {
		fmt.Println(result.PlaceID)
		getPlaceDetails(result.PlaceID, result.Name, client)
	}
}

func getPlaceDetails(placeID, name string, client *maps.Client) {
	req := &maps.PlaceDetailsRequest{
		PlaceID: placeID,
	}
	resp, err := client.PlaceDetails(context.Background(), req)
	if err != nil {
		log.Fatalf("Error in placing place details request: %s", err)
	}
	log.Println(name)
	log.Println(resp.FormattedAddress)
	log.Println(resp.FormattedPhoneNumber)
}
