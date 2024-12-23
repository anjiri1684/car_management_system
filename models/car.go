package models

import (
	"errors"
	"strconv"
	"time"

	"github.com/google/uuid"
)

type Car struct {
	ID uuid.UUID 	`json:"id"`
	Name string 	`json:"name"`
	Year string 	`json:"year"`
	Brand string 	`json:"brand"`
	FuelType string `json:"fuelType"`
	Engine Engine `json:"engine"`
	Price float64 `json:"price"`
	CreatedAt time.Time `json:"created_At"`
	UpdatedAT time.Time `json:"updated_At"`

}


type CarRequest struct {
	Name string 	`json:"name"`
	Year string 	`json:"year"`
	Brand string 	`json:"brand"`
	FuelType string `json:"fuelType"`
	Engine Engine `json:"engine"`
	Price float64 `json:"price"`
	
}

func ValidateRequest(CarReq CarRequest) error{
	if err := validateName(CarReq.Name);err != nil {
		return err
	}

	if err := validatePrice(CarReq.Price); err != nil {
		return err
	}

	if err := validateBrand(CarReq.Brand); err != nil {
		return err
	}

	if err:= validateYear(CarReq.Year); err != nil {
		return err
	} 

	if err := validateFuelType(CarReq.FuelType); err != nil {
		return err
	}

	if err := validateEngine(CarReq.Engine); err !=nil {
		return err
	}
		return nil	

}

func validateName(name string)error {
	if name == ""{
		return errors.New("Name is required")
	}
	return nil
}

func validateYear(year string) error {
	if year == "" {
		return errors.New("Year is required")
	}
	_, err := strconv.Atoi(year)
	if err != nil {
		return errors.New("Year must be a number")
	}
	currentYear:= time.Now().Year()
	yearInt , _ := strconv.Atoi(year)

if yearInt < 1886 || yearInt > currentYear {
	return errors.New("Year must be between 1886 and current year")
}
return nil
}

func validateBrand(brand string) error{
	if brand == "" {
		return errors.New("Brand is required")
	}
	return nil
}

func validateFuelType(fuelType string) error{
	validateFuelTypes:= []string{"Petrol", "Diesel", "Electric", "Hybrid"}
	for _, validType := range validateFuelTypes{
		if validType == fuelType {
			return nil
	}
}
return errors.New("FuelType Must be one of : Petrol, Diesel, Electric, Hybrid")
}


func validateEngine(engine Engine) error {
	if engine.EngineId == uuid.Nil{
		return errors.New("EngineID is required")
	}

	if engine.Displacement <=0 {
		return errors.New("Displacement must be greater than zero")
	}

	if engine.NoOfCylinders <= 0 {
		return errors.New("Number of cylinders must be greater than zero")
	}

	if engine.CarRange <=0 {
		return errors.New("CarRange must be greater than zero")
	}

	return nil
}

func validatePrice(price float64)error {
	if price <= 0 {
		return errors.New("Price must be greater than zero")
	}
	return nil
}