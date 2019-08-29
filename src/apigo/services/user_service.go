package services

import (
	"../domains"
	"../utils"
)

// tiene que tener una fncuion que me devuelva un site

func GetUser(userId int) (*domains.User, *utils.ApiError){
	user := domains.User{
		ID: userId,
	}
	if err := user.Get(); err != nil {
		return nil, err // este err ya es un puntero
	}

	return &user, nil
}


