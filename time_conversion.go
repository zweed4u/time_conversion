package main

import (
	"fmt"
	"os"
	"strconv"
)

func timeConversion(s string) string {
	hour, err := strconv.Atoi(s[:2])
	if err != nil {
		fmt.Println("unable to convert hour string slice to int:", err)
		os.Exit(1)
	}
	meridiem := s[len(s)-2:]
	if meridiem == "AM" {
		if hour == 12 { // 12AM
			return fmt.Sprintf("00%s", s[2:len(s)-2])
		}
		// ie 3AM
		return s[:len(s)-2]
	}

	if hour == 12 { // 12PM
		return s[:len(s)-2]
	}

	// ie 3PM
	return fmt.Sprintf("%d%s", hour+12, s[2:len(s)-2])
}

func main() {
	timeConversion("12:34:56AM")
}
