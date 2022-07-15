package service

import (
	"fmt"
	"net/http"
)

type CarSevice interface {
	FetchData()
}

const (
	carServiceUrl = "https://myfakeapi.com/api/cars/1"
)

type fetchCarDataService struct{}

func NewCarService() CarSevice {
	return &fetchCarDataService{}
}

func (*fetchCarDataService) FetchData() {
	client := http.Client{}
	fmt.Println("Fetching the url", carServiceUrl)

	// call external API
	resp, _ := client.Get(carServiceUrl)

	// write respone to the channel (TODO)
	carDataChannel <- resp
}
