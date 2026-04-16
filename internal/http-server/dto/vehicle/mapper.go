package vehicle

import (
	"github.com/NakonechniyVitalii/GoVehicleApi/internal/models"
	"github.com/NakonechniyVitalii/GoVehicleApi/internal/services/helper"
)

func (dto VehicleDTO) ToModel() models.Vehicle {
	return models.Vehicle{
		Brand:      helper.DerefUint16(dto.Brand),
		DriverType: helper.DerefUint16(dto.DriverType),
		Gearbox:    helper.DerefUint16(dto.Gearbox),
		BodyStyle:  helper.DerefUint16(dto.BodyStyle),
		Category:   helper.DerefUint16(dto.Category),
		Mileage:    helper.DerefUint32(dto.Mileage),
		Model:      helper.DerefString(dto.Model),
		Price:      helper.DerefUint32(dto.Price),
	}
}
