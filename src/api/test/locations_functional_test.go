package test

import (
	"encoding/json"
	"fmt"
	"github.com/alvaro259818/golang-restclient/rest"
	"github.com/alvaro259818/golang-testing/src/api/domain/locations"
	"github.com/alvaro259818/golang-testing/src/api/utils/errors"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestGetCountriesNotFound(t *testing.T)  {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          "https://api.mercadolibre.com/countries/CR",
		HTTPMethod:   http.MethodGet,
		RespHTTPCode: http.StatusNotFound,
		RespBody: `{"status": 404, "error": "not_found", "message": "no country with id CR"}`,
	})
	response, err := http.Get("http://localhost:8080/locations/countries/CR")
	assert.Nil(t, err)
	assert.NotNil(t, response)

	bytes, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(bytes))

	var apiErr errors.ApiError
	err = json.Unmarshal(bytes, &apiErr)
	assert.Nil(t, err)

	assert.EqualValues(t, http.StatusNotFound, apiErr.Status)
	assert.EqualValues(t, "not_found", apiErr.Error)
	assert.EqualValues(t, "no country with id CR", apiErr.Message)
}

func TestGetCountriesNoError(t *testing.T)  {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          "https://api.mercadolibre.com/countries/CR",
		HTTPMethod:   http.MethodGet,
		RespHTTPCode: http.StatusOK,
		RespBody:     `{"id":"CR","name":"Costa Rica","locale":"es_CR","currency_id":"CRC","decimal_separator":".","thousands_separator":",","time_zone":"GMT-06:00","geo_information":{"location":{"latitude":9.748916,"longitude":-83.753426}},"states":[{"id":"CR-A","name":"Alajuela"},{"id":"CR-C","name":"Cartago"},{"id":"CR-G","name":"Guanacaste"},{"id":"CR-H","name":"Heredia"},{"id":"CR-L","name":"Limón"},{"id":"CR-P","name":"Puntarenas"},{"id":"CR-SJ","name":"San José"}]}`,
	})
	response, err := http.Get("http://localhost:8080/locations/countries/CR")
	assert.Nil(t, err)
	assert.NotNil(t, response)

	bytes, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(bytes))

	var country locations.Country
	err = json.Unmarshal(bytes, &country)
	assert.Nil(t, err)

	assert.EqualValues(t, "CR", country.Id)
	assert.EqualValues(t, "Costa Rica", country.Name)
}