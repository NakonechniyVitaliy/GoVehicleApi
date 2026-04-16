package body_style

import (
	"github.com/NakonechniyVitalii/GoVehicleApi/internal/models"
	"github.com/NakonechniyVitalii/GoVehicleApi/internal/services/helper"
)

func (dto BodyStyleDTO) ToModel() models.BodyStyle {
	return models.BodyStyle{
		Name:  helper.DerefString(dto.Name),
		Value: helper.DerefUint16(dto.Value),
	}
}
