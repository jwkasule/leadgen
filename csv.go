package main

import (
	"encoding/csv"
	"log"
	"os"
)

// WriteCSV takes a filename and 2dimensional slice of slices of strings
// and writes those to the given filename after creating it
func WriteCSV(filename string, data [][]string) {
	file, err := os.Create(Concat("data/", filename, ".csv"))
	if err != nil {
		log.Fatalf("Error creating CSV file: %s", err)
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()
	for _, value := range data {
		err := writer.Write(value)
		if err != nil {
			log.Fatalf("Error writing CSV file: %s", err)
		}
	}
}
