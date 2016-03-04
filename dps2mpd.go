package main

import "fmt"
import "os"
import "strconv"

func main() {
	if len(os.Args) == 1 {
		help()
	}

	dps, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Println("Could not parse " + os.Args[1] + " as a floating point number")
		os.Exit(1)
	}

	mpd := dps / 1000000 * (24 * 60 * 60)
	fmt.Printf("%.f\n", mpd)
}

func help() {
	fmt.Println("Usage: dps2mpd DPS")
	fmt.Println("Convert DPS (documents per second) to millions of documents per day")
	os.Exit(0)
}
