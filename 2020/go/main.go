package main

import (
	"flag"
	"log"
)

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	day := flag.Int("day", 1, "number of the day")
	switch *day {
	case 1:
		day1()
	}
}
