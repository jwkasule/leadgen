package main

import (
	"fmt"
	"os"
)

const helpmessage = `
Leadgen, a cli client for hitting the google places API for a search
term and outputting the results with phone numbers and emails in a CSV 
file

Usage:

Making sure API_KEY environment variable is set first (check env/api.key)
and run as follows:

leadgen "[search term]" "[location]" [radius]

Where "[search term]" is the type of entity you are searching for,
"[location]" is a location in coordinates specified as latitude,longitude,
and [radius] is radius in amount of meters.

Resulting CSV file will be in data/[date]_leads.csv
`

func PrintHelp() {
	fmt.Println(helpmessage)
	os.Exit(1)
}
