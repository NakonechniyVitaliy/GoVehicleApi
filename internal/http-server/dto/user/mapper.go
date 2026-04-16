package user

import (
	"github.com/NakonechniyVitalii/GoVehicleApi/internal/models"
	"github.com/NakonechniyVitalii/GoVehicleApi/internal/services/helper"
)

func (dto SignUpDTO) ToModel() models.User {
	return models.User{
		Username: helper.DerefString(dto.Username),
		Login:    helper.DerefString(dto.Login),
		Password: helper.DerefString(dto.Password),
	}
}
