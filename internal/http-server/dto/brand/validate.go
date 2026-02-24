package brand

import (
	"fmt"

	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/http-server/dto/_errors"
)

type SaveRequest struct {
	Brand BrandDTO
}

func (r SaveRequest) Validate() error {
	b := r.Brand

	if b.CategoryID == nil ||
		b.Count == nil ||
		b.CountryID == nil ||
		b.EngName == nil ||
		b.MarkaID == nil ||
		b.Name == nil ||
		b.Slang == nil ||
		b.Value == nil {
		return fmt.Errorf(_errors.AllFieldsAreRequired)
	}
	return nil
}

type UpdateRequest struct {
	Brand BrandDTO
}

func (r UpdateRequest) Validate() error {
	b := r.Brand

	if b.CategoryID == nil ||
		b.Count == nil ||
		b.CountryID == nil ||
		b.EngName == nil ||
		b.MarkaID == nil ||
		b.Name == nil ||
		b.Slang == nil ||
		b.Value == nil {
		return fmt.Errorf(_errors.AllFieldsAreRequired)
	}
	return nil
}
