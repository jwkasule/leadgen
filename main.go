package main

import "os"

func main() {
	if len(os.Args) > 2 {
		PrintHelp()
	}
	term := os.Args[1]
	location := os.Args[2]
	radius := os.Args[3]
	key := os.Getenv("API_KEY")
	GetResultHandler(term, location, key, radius)
}
