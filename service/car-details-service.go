package service

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/yogie/go-clean-api/entity"
)

type CarDetailsService interface {
	GetDetails() entity.CarDetails
}

type carServiceStruct struct{}

var (
	carService       CarSevice   = NewCarService()
	ownerService     OwnerSevice = NewOwnerService()
	carDataChannel               = make(chan *http.Response)
	ownerDataChannel             = make(chan *http.Response)
)

func NewCarDetailsService() CarDetailsService {
	return &carServiceStruct{}
}

func (*carServiceStruct) GetDetails() entity.CarDetails {
	// goroutine get data from https://myfakeapi.com/api/cars/1 -- endpoint 1
	go carService.FetchData()

	// goroutine get data from https://myfakeapi.com/api/users/1 --  endpoint 2
	go ownerService.FetchData()

	// create carChannel channel to get the data from endpoint 1
	// create ownerChannel channel to get the data from endpoint 2
	car, _ := getCarData()
	owner, _ := getOwnerData()

	return entity.CarDetails{
		ID:             car.CarData.ID,
		Brand:          car.CarData.Brand,
		Model:          car.CarData.Model,
		Year:           car.CarData.Year,
		Vin:            car.CarData.Vin,
		OwnerFirstName: owner.OwnerData.FirstName,
		OwnerLastName:  owner.OwnerData.LastName,
		OwnerEmail:     owner.OwnerData.Email,
		OwnerJobTitle:  owner.OwnerData.JobTitle,
	}
}

func getCarData() (entity.Car, error) {
	r1 := <-carDataChannel
	var car entity.Car

	err := json.NewDecoder(r1.Body).Decode(&car)
	if err != nil {
		fmt.Println(err.Error())
		return car, err
	}

	return car, nil
}

func getOwnerData() (entity.Owner, error) {
	r1 := <-ownerDataChannel
	var owner entity.Owner

	err := json.NewDecoder(r1.Body).Decode(&owner)
	if err != nil {
		fmt.Println(err.Error())
		return owner, err
	}

	return owner, nil
}
