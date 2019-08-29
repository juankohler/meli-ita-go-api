package services

import(
	"../utils"
	"../domains"
	"sync"
)

func GetResult(userId int) (*domains.Result, *utils.ApiError){

	user := domains.User{
		ID: userId,
	}

	if err := user.Get(); err != nil {
		return nil, err // este err ya es un puntero
	}

	country := domains.Country{
		ID: user.CountryID,
	}

	if err := country.Get(); err != nil {
		return nil, err // este err ya es un puntero
	}

	site := domains.Site{
		ID: user.SiteID,
	}

	if err := site.Get(); err != nil {
		return nil, err // este err ya es un puntero
	}

	resp := domains.Result{
		User: &user,
		Site: &site,
		Country: &country,
	}

	/*
		resp := &domains.Result{
		User: user,
		Site: site,
		Country: country,
	}
	*/

	return &resp, nil

}

func GetResultWg(userId int) (*domains.Result, *utils.ApiError){
	var wg sync.WaitGroup
	user := domains.User{
		ID: userId,
	}

	user.Get()

	country := domains.Country{
		ID: user.CountryID,
	}

	site := domains.Site{
		ID: user.SiteID,
	}
	wg.Add(2)
	go country.Getwg(&wg)
	go site.GetWg(&wg)
	wg.Wait()
	resp := domains.Result{
		User: &user,
		Site: &site,
		Country: &country,
	}

	/*
		resp := &domains.Result{
		User: user,
		Site: site,
		Country: country,
	}
	*/
	return &resp, nil

}

func GetResultCh(userId int) (*domains.Result, *utils.ApiError){
	user := domains.User{
		ID: userId,
	}

	user.Get()

	country := domains.Country{
		ID: user.CountryID,
	}

	site := domains.Site{
		ID: user.SiteID,
	}

	valores := make(chan domains.Result, 2)
	defer close(valores)
	go country.GetCh(valores)
	go site.GetCh(valores)
	resultFinal := domains.Result{
		User: &user,
	}


	for i:=0;i<2;i++ {
		result := <- valores
		if result.Site != nil {
			resultFinal.Site = result.Site
		} else if result.Country != nil {
			resultFinal.Country = result.Country
		} else if result.ApiError != nil {
			resultFinal.ApiError = result.ApiError
		}
	}

	/*
		resp := &domains.Result{
		User: user,
		Site: site,
		Country: country,
	}
	*/
	return &resultFinal, nil

}