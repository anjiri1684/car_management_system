package models

import "github.com/google/uuid"

type Engine struct {
	EngineId uuid.UUID `json:"engineId"`
	Displacement int64 `json:"displacement"`
	NoOfCylinders int `json:"noOfCylinders"`
	CarRange int64 `json:"carRange"`

}