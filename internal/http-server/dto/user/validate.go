package gearbox

import (
	"fmt"

	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/http-server/dto/_errors"
)

type SignUpRequest struct {
	User SignUpDTO
}

func (r SignUpRequest) Validate() error {
	g := r.User

	if g.Username == nil || g.Login == nil || g.Password == nil {
		return fmt.Errorf(_errors.AllFieldsAreRequired)
	}
	return nil
}

type SignInRequest struct {
	SignData SignInDTO
}

func (r SignInRequest) Validate() error {
	g := r.SignData

	if g.Login == nil || g.Password == nil {
		return fmt.Errorf(_errors.AllFieldsAreRequired)
	}
	return nil
}
