package controller

import (
	"encoding/json"
	"net/http"

	"github.com/yogie/go-clean-api/service"
)

type carController struct{}

var (
	carDetailsService service.CarDetailsService
)

type CarDetailsController interface {
	GetCarDetails(response http.ResponseWriter, request *http.Request)
}

func NewCarDetailsController(service service.CarDetailsService) CarDetailsController {
	carDetailsService = service
	return &carController{}
}

func (*carController) GetCarDetails(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	result := carDetailsService.GetDetails()
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(result)
}
