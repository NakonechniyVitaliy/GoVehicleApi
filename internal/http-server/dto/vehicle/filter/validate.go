package filter

import (
	"fmt"
	"strconv"
)

func (f FilterDTO) ValidateAndToModel() (*Filter, error) {
	filter := &Filter{
		Page:  0,
		Limit: 5,
	}

	if f.Page != "" {
		page, err := strconv.Atoi(f.Page)
		if err != nil || page < 1 {
			return nil, fmt.Errorf("invalid page number")
		}
		filter.Page = page
	}

	if f.Limit != "" {
		limit, err := strconv.Atoi(f.Limit)
		if err != nil || limit < 1 {
			return nil, fmt.Errorf("invalid limit number")
		}
		filter.Limit = limit
	}

	return filter, nil
}
