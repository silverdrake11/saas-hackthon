package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {

	city := flag.String("city", "", "Name of the city")
	day := flag.String("day", "", "Day for the weather forecast (e.g., '2024-10-31')")
	flag.Parse()

	if *city == "" || *day == "" {
		log.Fatal("Both city and day must be provided.")
	}

	lat, lon := FindCityLocation(City{Name: *city})
	fmt.Printf("Weather in %s: %s\n", lat, lon)

	// loc := Location{
	// 	Latitude:  lat,
	// 	Longitude: lon,
	// }

	// weather := GetWeather(loc)

	// fmt.Printf("Weather in %s: %s\n", *city, weather)
	fmt.Printf("Weather in %s: %s\n", *city, "laal")
}
