package gearbox

import (
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/models"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/services/helper"
)

func (dto SignUpDTO) ToModel() models.User {
	return models.User{
		Username: helper.DerefString(dto.Username),
		Login:    helper.DerefString(dto.Login),
		Password: helper.DerefString(dto.Password),
	}
}
