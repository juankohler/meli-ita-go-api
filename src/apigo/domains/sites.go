package domains

import (
	"../utils"
	"net/http"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"sync"
)

type Site struct {
	ID                 string   `json:"id"`
	Name               string   `json:"name"`
	CountryID          string   `json:"country_id"`
	SaleFeesMode       string   `json:"sale_fees_mode"`
	MercadopagoVersion int      `json:"mercadopago_version"`
	DefaultCurrencyID  string   `json:"default_currency_id"`
	ImmediatePayment   string   `json:"immediate_payment"`
	PaymentMethodIds   []string `json:"payment_method_ids"`
	Settings           struct {
		IdentificationTypes      []string `json:"identification_types"`
		TaxpayerTypes            []string `json:"taxpayer_types"`
		IdentificationTypesRules []struct {
			IdentificationType string `json:"identification_type"`
			Rules              []struct {
				EnabledTaxpayerTypes []interface{} `json:"enabled_taxpayer_types"`
				BeginsWith           string        `json:"begins_with"`
				Type                 string        `json:"type"`
				MinLength            int           `json:"min_length"`
				MaxLength            int           `json:"max_length"`
			} `json:"rules"`
		} `json:"identification_types_rules"`
	} `json:"settings"`
	Currencies []struct {
		ID     string `json:"id"`
		Symbol string `json:"symbol"`
	} `json:"currencies"`
	Categories []struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"categories"`
}


func (site *Site) Get() *utils.ApiError{
	if site.ID == "" {
		return &utils.ApiError{
			Message: "Site ID is empty",
			Status: http.StatusBadRequest,
		}
	}

	url := fmt.Sprintf("%s%s", utils.UrlSite, site.ID)
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

	if err := json.Unmarshal(data, &site); err != nil {
		return &utils.ApiError{
			Message: err.Error(),
			Status: http.StatusInternalServerError,
		}
	}

	return nil
}


func (site *Site) GetWg(wg *sync.WaitGroup) *utils.ApiError{
	if site.ID == "" {
		wg.Done()
		return &utils.ApiError{
			Message: "Site ID is empty",
			Status: http.StatusBadRequest,
		}
	}

	url := fmt.Sprintf("%s%s", utils.UrlSite, site.ID)
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

	if err := json.Unmarshal(data, &site); err != nil {
		wg.Done()
		return &utils.ApiError{
			Message: err.Error(),
			Status: http.StatusInternalServerError,
		}
	}
	wg.Done()
	return nil
}
func (site Site) GetCh(result chan Result) {
	result1 := Result{}
	if site.ID == "" {
		apiErr := &utils.ApiError{
			Message: "Site ID is empty",
			Status: http.StatusBadRequest,
		}
		result1.ApiError = apiErr
		result <- result1
	}

	url := fmt.Sprintf("%s%s", utils.UrlSite, site.ID)
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

	if err := json.Unmarshal(data, &site); err != nil {
		apiErr := &utils.ApiError{
			Message: err.Error(),
			Status: http.StatusInternalServerError,
		}
		result1.ApiError = apiErr
		result <- result1
	}
	result1.Site = &site
	result <- result1
}