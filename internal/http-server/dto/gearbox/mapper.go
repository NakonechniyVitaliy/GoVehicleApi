package gearbox

import (
	"github.com/NakonechniyVitalii/GoVehicleApi/internal/models"
	"github.com/NakonechniyVitalii/GoVehicleApi/internal/services/helper"
)

func (dto GearboxDTO) ToModel() models.Gearbox {
	return models.Gearbox{
		Name:  helper.DerefString(dto.Name),
		Value: helper.DerefUint16(dto.Value),
	}
}
