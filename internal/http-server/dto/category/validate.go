package category

import (
	"fmt"

	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/http-server/dto/_errors"
)

type SaveRequest struct {
	CategoryDTO
}

func (r SaveRequest) Validate() error {
	c := r.CategoryDTO

	if c.Name == nil || c.Value == nil {
		return fmt.Errorf(_errors.AllFieldsAreRequired)
	}
	return nil
}

type UpdateRequest struct {
	CategoryDTO
}

func (r UpdateRequest) Validate() error {
	c := r.CategoryDTO

	if c.Name == nil || c.Value == nil {
		return fmt.Errorf(_errors.AllFieldsAreRequired)
	}
	return nil
}
