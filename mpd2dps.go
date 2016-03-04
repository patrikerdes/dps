package main

import "fmt"
import "os"
import "strconv"

func main() {
	if len(os.Args) == 1 {
		help()
	}

	mpd, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Println("Could not parse " + os.Args[1] + " as a floating point number")
		os.Exit(1)
	}

	dps := mpd * 1000000 / (24 * 60 * 60)
	fmt.Printf("%.f\n", dps)
}

func help() {
	fmt.Println("Usage: mpd2dps MPD")
	fmt.Println("Convert MPD (millions of documents per day) to documents per second")
	os.Exit(0)
}
