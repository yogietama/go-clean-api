package service

import (
	"fmt"
	"net/http"
)

type OwnerSevice interface {
	FetchData()
}

const (
	ownerServiceUrl = "https://myfakeapi.com/api/users/1"
)

type fetchOwnerDataService struct{}

func NewOwnerService() OwnerSevice {
	return &fetchOwnerDataService{}
}

func (*fetchOwnerDataService) FetchData() {
	client := http.Client{}
	fmt.Println("Fetching the url", ownerServiceUrl)

	// call external API
	resp, _ := client.Get(ownerServiceUrl)

	// write respone to the channel
	ownerDataChannel <- resp
}
