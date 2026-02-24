package gearbox

import (
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/models"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/service/helper"
)

func (dto GearboxDTO) ToModel() models.Gearbox {
	return models.Gearbox{
		Name:  helper.DerefString(dto.Name),
		Value: helper.DerefUint16(dto.Value),
	}
}
