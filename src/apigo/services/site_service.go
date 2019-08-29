package services

import (
	"../domains"
	"../utils"
)

// tiene que tener una fncuion que me devuelva un site

func GetSite(siteId string) (*domains.Site, *utils.ApiError){
	site := domains.Site{
		ID: siteId,
	}
	if err := site.Get(); err != nil {
		return nil, err // este err ya es un puntero
	}
	return &site, nil
}


