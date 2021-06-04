package locations_provider

import (
	"github.com/alvaro259818/golang-restclient/rest"
	"github.com/stretchr/testify/assert"
	"net/http"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	rest.StartMockupServer()
	os.Exit(m.Run())
}

func TestGetCountryRestClientError(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          "https://api.mercadolibre.com/countries/CR",
		HTTPMethod:   http.MethodGet,
		RespHTTPCode: -1,
	})
	country, err := GetCountry("CR")

	assert.Nil(t, country)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "invalid restclient error when getting country CR", err.Message)
}

func TestGetCountryNotFound(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          "https://api.mercadolibre.com/countries/CR",
		HTTPMethod:   http.MethodGet,
		RespHTTPCode: http.StatusNotFound,
		RespBody:     `{"message": "Country not found", "error": "not_found", "status": 404, "cause": []}`,
	})
	country, err := GetCountry("CR")

	assert.Nil(t, country)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.Status)
	assert.EqualValues(t, "Country not found", err.Message)
}

func TestGetCountryInvalidJsonInterface(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          "https://api.mercadolibre.com/countries/CR",
		HTTPMethod:   http.MethodGet,
		RespHTTPCode: http.StatusNotFound,
		RespBody:     `{"message": "Country not found", "error": "not_found", "status": "404", "cause": []}`,
	})
	country, err := GetCountry("CR")

	assert.Nil(t, country)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "invalid error response when getting country CR", err.Message)
}

func TestGetCountryInvalidJsonResponse(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          "https://api.mercadolibre.com/countries/CR",
		HTTPMethod:   http.MethodGet,
		RespHTTPCode: http.StatusOK,
		RespBody:     `{"id": 123, "name": "Costa Rica", "time_zone": "GMT-06:00"}`,
	})
	country, err := GetCountry("CR")

	assert.Nil(t, country)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "error when trying to unmarshal country data for CR", err.Message)
}

func TestGetCountryNoError(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          "https://api.mercadolibre.com/countries/CR",
		HTTPMethod:   http.MethodGet,
		RespHTTPCode: http.StatusOK,
		RespBody:     `{"id":"CR","name":"Costa Rica","locale":"es_CR","currency_id":"CRC","decimal_separator":".","thousands_separator":",","time_zone":"GMT-06:00","geo_information":{"location":{"latitude":9.748916,"longitude":-83.753426}},"states":[{"id":"CR-A","name":"Alajuela"},{"id":"CR-C","name":"Cartago"},{"id":"CR-G","name":"Guanacaste"},{"id":"CR-H","name":"Heredia"},{"id":"CR-L","name":"Limón"},{"id":"CR-P","name":"Puntarenas"},{"id":"CR-SJ","name":"San José"}]}`,
	})
	country, err := GetCountry("CR")

	assert.Nil(t, err)
	assert.NotNil(t, country)
	assert.EqualValues(t, "CR", country.Id)
	assert.EqualValues(t, "Costa Rica", country.Name)
	assert.EqualValues(t, "GMT-06:00", country.TimeZone)
	assert.EqualValues(t, 7, len(country.States))
}
