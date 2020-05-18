package main

import "os"

func main() {
	if len(os.Args) != 4 {
		PrintHelp()
		os.Exit(1)
	}
	query := os.Args[1]
	location := os.Args[2]
	radius := os.Args[3]
	key := os.Getenv("API_KEY")
	GetResultHandler(query, location, radius, key)
}
