package services

import (
	"mediaLibrary_v2/utils"
)

type service struct {
	expectedParameters map[string]*utils.Parameter
}

func (i service) GetExpectedParameters() map[string]*utils.Parameter {
	return i.expectedParameters
}
