package brand

import (
	"fmt"
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
		return fmt.Errorf("all fields are required")
	}
	return nil
}

type UpdateRequest struct {
	Brand BrandDTO
}

func (r UpdateRequest) Validate() error {
	b := r.Brand

	if b.CategoryID == nil &&
		b.Count == nil &&
		b.CountryID == nil &&
		b.EngName == nil &&
		b.MarkaID == nil &&
		b.Name == nil &&
		b.Slang == nil &&
		b.Value == nil {
		return fmt.Errorf("at least one field must be provided")
	}
	return nil
}
