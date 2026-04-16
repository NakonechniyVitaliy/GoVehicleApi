package category

import (
	"github.com/NakonechniyVitalii/GoVehicleApi/internal/models"
	"github.com/NakonechniyVitalii/GoVehicleApi/internal/services/helper"
)

func (dto CategoryDTO) ToModel() models.Category {
	return models.Category{
		Name:  helper.DerefString(dto.Name),
		Value: helper.DerefUint16(dto.Value),
	}
}
