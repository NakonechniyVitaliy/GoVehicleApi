package brand

import (
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/models"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/service/helper"
)

func (dto BrandDTO) ToModel() models.Brand {
	return models.Brand{
		CategoryID: helper.DerefUint16(dto.CategoryID),
		Count:      helper.DerefUint16(dto.Count),
		CountryID:  helper.DerefUint16(dto.CountryID),
		EngName:    helper.DerefString(dto.EngName),
		MarkaID:    helper.DerefUint16(dto.MarkaID),
		Name:       helper.DerefString(dto.Name),
		Slang:      helper.DerefString(dto.Slang),
		Value:      helper.DerefUint16(dto.Value),
	}
}
