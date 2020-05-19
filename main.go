package main

import "os"

func main() {
	if len(os.Args) != 5 {
		PrintHelp()
		os.Exit(1)
	}
	filename := os.Args[1]
	query := os.Args[2]
	location := os.Args[3]
	radius := os.Args[4]
	key := os.Getenv("API_KEY")
	GetResultHandler(filename, query, location, radius, key)
}
