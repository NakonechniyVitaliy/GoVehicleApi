package category

import (
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/models"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/service/helper"
)

func (dto CategoryDTO) ToModel() models.Category {
	return models.Category{
		Name:  helper.DerefString(dto.Name),
		Value: helper.DerefUint16(dto.Value),
	}
}
