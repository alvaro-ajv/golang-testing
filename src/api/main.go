package main

import (
	"fmt"
	"github.com/alvaro259818/golang-testing/src/api/providers/locations_provider"
)

func main() {
	country, err := locations_provider.GetCountry("CR")
	fmt.Println(err)
	fmt.Println(country)
}
