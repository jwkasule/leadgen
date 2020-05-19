package main

import (
	"fmt"
	"log"
	"strings"
)

const helpmessage = `
Leadgen, a cli client for hitting the google places API for a search
term, lat/long coordinates, and radius, and outputting the results 
with phone numbers and emails in a CSV file for every example google 
maps can find for these inputs.

Usage:

Make sure API_KEY environment variable is set first (source env/api.sh)
and run as follows, making sure [query] and [location] are in quotes:

leadgen [filename] "[query]" "[location]" [radius]

Where "filename" saves the results in data/filename.csv, "[query]" is the 
type of entity you are searching for, "[location]" is in the format of 
coordinates of lat,long and radius is an integer measuring meters from 
the location point to search from (No more than 50000 metres/50km allowed).

Resulting CSV file will be in data/[filename].csv

`

// PrintHelp prints the default help message
func PrintHelp() {
	fmt.Println(helpmessage)
}

// Concat iterates over strings and writes to a strings.Builder
// for faster concat
func Concat(keys ...string) string {
	var b strings.Builder
	for i := 0; i < len(keys); i++ {
		if _, err := b.WriteString(keys[i]); err != nil {
			log.Fatalf("Error in concatenating buffer: %s", err)
		}
	}
	return b.String()
}
