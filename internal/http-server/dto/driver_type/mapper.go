package driver_type

import (
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/models"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/service/helper"
)

func (dto DriverTypeDTO) ToModel() models.DriverType {
	return models.DriverType{
		Name:       helper.DerefString(dto.Name),
		Value:      helper.DerefUint16(dto.Value),
	}
}
