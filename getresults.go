/*
Copyright 2022 Dmitri DB

Redistribution and use in source and binary forms, with or without modification, are permitted provided that the following conditions are met:

1. Redistributions of source code must retain the above copyright notice, this list of conditions and the following disclaimer.

2. Redistributions in binary form must reproduce the above copyright notice, this list of conditions and the following disclaimer in the documentation and/or other materials provided with the distribution.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
*/
package main

import (
	"context"
	"fmt"
	"log"

	"googlemaps.github.io/maps"
)

// GetResults takes a given query, lat/long, radius and api key, and returns
// the results as a 2 dimensional array of strings
func GetResults(query string, lat, long float64, radius uint, key string) [][]string {
	// Initialize the slice to return
	var results = [][]string{}
	client, err := maps.NewClient(maps.WithAPIKey(key))
	if err != nil {
		log.Fatalf("Error in initializing maps client: %s", err)
	}
	location := &maps.LatLng{
		Lat: lat,
		Lng: long,
	}
	// First 20 results
	req := &maps.TextSearchRequest{
		Location: location,
		Radius:   radius,
		Query:    query,
	}
	resp, err := client.TextSearch(context.Background(), req)
	if err != nil {
		log.Fatalf("Error in placing text search request 1: %s", err)
	}
	for _, result := range resp.Results {
		fmt.Println(result.PlaceID)
		address, number := getPlaceDetails(result.PlaceID, result.Name, client)
		resultAppend := []string{result.Name, address, number}
		results = append(results, resultAppend)
	}
	// There's some DRY refactoring that could be done here but let's
	// save that optimization for some other time
	if resp.NextPageToken != "" {
		// Next 20 results!
		req := &maps.TextSearchRequest{
			Location:  location,
			Radius:    radius,
			Query:     query,
			PageToken: resp.NextPageToken,
		}
		resp2, err := client.TextSearch(context.Background(), req)
		if err != nil {
			log.Fatalf("Error in placing text search request 2: %s", err)
		}
		for _, result := range resp2.Results {
			fmt.Println(result.PlaceID)
			address, number := getPlaceDetails(result.PlaceID, result.Name, client)
			resultAppend := []string{result.Name, address, number}
			results = append(results, resultAppend)
		}
		if resp2.NextPageToken != "" {
			// Last 20 results!
			req := &maps.TextSearchRequest{
				Location:  location,
				Radius:    radius,
				Query:     query,
				PageToken: resp.NextPageToken,
			}
			resp3, err := client.TextSearch(context.Background(), req)
			if err != nil {
				log.Fatalf("Error in placing text search request 3: %s", err)
			}
			for _, result := range resp3.Results {
				fmt.Println(result.PlaceID)
				address, number := getPlaceDetails(result.PlaceID, result.Name, client)
				resultAppend := []string{result.Name, address, number}
				results = append(results, resultAppend)
			}
		}
	}
	return results
}

// getPlaceDetails hits the places API endpoint for details about a given
// placeID
func getPlaceDetails(placeID, name string, client *maps.Client) (address string, phonenumber string) {
	req := &maps.PlaceDetailsRequest{
		PlaceID: placeID,
	}
	resp, err := client.PlaceDetails(context.Background(), req)
	if err != nil {
		log.Fatalf("Error in placing place details request: %s", err)
	}
	fmt.Println(name)
	fmt.Println(resp.FormattedAddress)
	fmt.Println(resp.FormattedPhoneNumber)
	return resp.FormattedAddress, resp.FormattedPhoneNumber
}
