package main

import (
	"fmt"
)

const helpmessage = `
Leadgen, a cli client for hitting the google places API for a search
term and country, and outputting the results with phone numbers and 
emails in a CSV file for every example google maps can find for these.

Usage:

Making sure API_KEY environment variable is set first (source env/api.sh)
and run as follows, making sure [query] is in quotes:

leadgen "[query]" "[location]" [radius]

Where "[query]" is the type of entity you are searching for, and
"[location]" is in the format of coordinates of lat,long and radius
is an integer measuring meters from the location point to search from.

Resulting CSV file will be in data/[date]_leads.csv
`

func PrintHelp() {
	fmt.Println(helpmessage)
}
