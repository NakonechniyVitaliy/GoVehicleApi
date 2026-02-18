package vehicle

import (
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/models"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/service/helper"
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
		Price:      helper.DerefUint16(dto.Price),
	}
}
