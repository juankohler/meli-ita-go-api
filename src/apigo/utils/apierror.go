package utils

// error que voy a manejar

type ApiError struct {
	Message string `json:"message"`
	Status int `json:"status"`
}
