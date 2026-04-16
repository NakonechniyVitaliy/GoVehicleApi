package user

import (
	"fmt"

	"github.com/NakonechniyVitalii/GoVehicleApi/internal/http-server/dto/_errors"
)

func (dto SignUpDTO) Validate() error {

	if dto.Username == nil || dto.Login == nil || dto.Password == nil {
		return fmt.Errorf(_errors.AllFieldsAreRequired)
	}
	return nil
}

func (dto SignInDTO) Validate() error {
	if dto.Login == nil || dto.Password == nil {
		return fmt.Errorf(_errors.AllFieldsAreRequired)
	}
	return nil
}
