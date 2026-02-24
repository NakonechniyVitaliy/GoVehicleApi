package body_style

import (
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/models"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/service/helper"
)

func (dto BodyStyleDTO) ToModel() models.BodyStyle {
	return models.BodyStyle{
		Name:       helper.DerefString(dto.Name),
		Value:      helper.DerefUint16(dto.Value),
	}
}
