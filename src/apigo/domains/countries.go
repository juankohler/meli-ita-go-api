package domains

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"../utils"
	"sync"
)

type Country struct {
	ID                 string `json:"id"`
	Name               string `json:"name"`
	Locale             string `json:"locale"`
	CurrencyID         string `json:"currency_id"`
	DecimalSeparator   string `json:"decimal_separator"`
	ThousandsSeparator string `json:"thousands_separator"`
	TimeZone           string `json:"time_zone"`
	GeoInformation     struct {
		Location struct {
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"location"`
	} `json:"geo_information"`
	States []struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"states"`
}

func (country *Country) Getwg(wg *sync.WaitGroup) *utils.ApiError{
	if country.ID == "" {
		wg.Done()
		return &utils.ApiError{
			Message: "Country ID is empty",
			Status: http.StatusBadRequest,
		}
	}

	url := fmt.Sprintf("%s%s", utils.UrlCountry, country.ID)
	response, err := http.Get(url)
	if err != nil {
		wg.Done()
		return &utils.ApiError{
			Message: err.Error(),
			Status: http.StatusInternalServerError,
		}
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		wg.Done()
		return &utils.ApiError{
			Message: err.Error(),
			Status: http.StatusInternalServerError,
		}
	}

	if err := json.Unmarshal(data, &country); err != nil {
		wg.Done()
		return &utils.ApiError{
			Message: err.Error(),
			Status: http.StatusInternalServerError,
		}
	}
	wg.Done()
	return nil
}

func (country *Country) Get() *utils.ApiError{
	if country.ID == "" {
		return &utils.ApiError{
			Message: "Country ID is empty",
			Status: http.StatusBadRequest,
		}
	}

	url := fmt.Sprintf("%s%s", utils.UrlCountry, country.ID)
	response, err := http.Get(url)
	if err != nil {
		return &utils.ApiError{
			Message: err.Error(),
			Status: http.StatusInternalServerError,
		}
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return &utils.ApiError{
			Message: err.Error(),
			Status: http.StatusInternalServerError,
		}
	}

	if err := json.Unmarshal(data, &country); err != nil {
		return &utils.ApiError{
			Message: err.Error(),
			Status: http.StatusInternalServerError,
		}
	}
	return nil
}


func (country Country) GetCh(result chan Result) {
	result1 := Result{}

	if country.ID == "" {
		apiErr := &utils.ApiError{
			Message: "Country ID is empty",
			Status: http.StatusBadRequest,
		}
		result1.ApiError = apiErr
		result <- result1
	}

	url := fmt.Sprintf("%s%s", utils.UrlCountry, country.ID)
	response, err := http.Get(url)
	if err != nil {
		apiErr := &utils.ApiError{
			Message: err.Error(),
			Status: http.StatusInternalServerError,
		}
		result1.ApiError = apiErr
		result <- result1
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		apiErr := &utils.ApiError{
			Message: err.Error(),
			Status: http.StatusInternalServerError,
		}
		result1.ApiError = apiErr
		result <- result1
	}

	if err := json.Unmarshal(data, &country); err != nil {
		apiErr := &utils.ApiError{
			Message: err.Error(),
			Status: http.StatusInternalServerError,
		}
		result1.ApiError = apiErr
		result <- result1
	}
	result1.Country = &country
	result <- result1
}