package brand

import "github.com/NakonechniyVitaliy/GoVehicleApi/internal/models"

func (dto BrandDTO) ToModel() models.Brand {
	return models.Brand{
		CategoryID: dto.CategoryID,
		Count:      dto.Count,
		CountryID:  dto.CountryID,
		EngName:    dto.EngName,
		MarkaID:    dto.MarkaID,
		Name:       dto.Name,
		Slang:      dto.Slang,
		Value:      dto.Value,
	}
}
