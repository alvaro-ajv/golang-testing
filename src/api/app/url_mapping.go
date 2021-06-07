package app

import "github.com/alvaro259818/golang-testing/src/api/controllers"

func mapUrls() {
	router.GET("/locations/countries/:country_id", controllers.GetCountry)
}
