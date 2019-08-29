package services

import (
	"../domains"
	"../utils"
)

// tiene que tener una fncuion que me devuelva un site

func GetCountry(countryId string) (*domains.Country, *utils.ApiError){
	country := domains.Country{
		ID: countryId,
	}
	if err := country.Get(); err != nil {
		return nil, err // este err ya es un puntero
	}

	return &country, nil
}



