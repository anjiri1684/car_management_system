package models

import (
	"errors"

	"github.com/google/uuid"
)

type Engine struct {
	EngineId uuid.UUID `json:"engineId"`
	Displacement int64 `json:"displacement"`
	NoOfCylinders int64 `json:"noOfCylinders"`
	CarRange int64 `json:"carRange"`

}

type EngineRequest struct {
	Displacement int64 `json:"displacement"`
	NoOfCylinders int64 `json:"noOfCylinders"`
	CarRange int64 `json:"carRange"`
}




func ValidateEngineRequest(engineReq EngineRequest)error {
		if err := validateDisplacement(engineReq.Displacement); err != nil {
				return err
			}

		if err := validateNumberOfCylinders(engineReq.NoOfCylinders); err != nil {
			return err
		}
		if err := validateCarRange(engineReq.CarRange); err != nil {
			return err
		}

	return nil

}

func validateDisplacement(displacement int64) error {
	if displacement <=0 {
		return errors.New("Displacement must be greater than zero")
	}

	return nil
}

func validateNumberOfCylinders(noOfCylinders int64) error {
	if noOfCylinders <= 0 {
		return errors.New("Number of cylinders must be greater than zero")
	}

	return nil
}


func validateCarRange(carRange int64) error {
	if carRange <= 0 {
		return errors.New("Car range must be greater than zero")
	}

	return nil
}