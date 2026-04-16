package driver_type

import (
	"github.com/NakonechniyVitalii/GoVehicleApi/internal/models"
	"github.com/NakonechniyVitalii/GoVehicleApi/internal/services/helper"
)

func (dto DriverTypeDTO) ToModel() models.DriverType {
	return models.DriverType{
		Name:  helper.DerefString(dto.Name),
		Value: helper.DerefUint16(dto.Value),
	}
}
