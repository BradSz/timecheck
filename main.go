package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {

	errorcode := 0

	// Parse command line arguments
	timezones := flag.String("t", time.Local.String(), "Comma-separated list of time zones to display")
	flag.Parse()

	t := time.Now()
	for _, tz := range strings.Split(*timezones, ",") {

		loc, err := time.LoadLocation(tz)
		if err != nil {
			errorcode = 1
			fmt.Println("Time Zone omitted; not found: ")
		} else {
			fmt.Println(t.In(loc), "in", loc.String())
		}
	}

	os.Exit(errorcode)
}
