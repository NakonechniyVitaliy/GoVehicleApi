package brand

import "github.com/brianvoe/gofakeit/v6"

type PositiveTestCase struct {
	CaseName   string
	CategoryID uint16
	Count      uint16
	CountryID  uint16
	EngName    string
	MarkaID    uint16
	BrandName  string
	Slang      string
	Value      uint16
}
type InvalidJsonTestCase struct {
	CaseName string
	Error    string
	Brand    map[string]any
}

var PositiveCases = []PositiveTestCase{
	{
		CaseName:   "Valid brand",
		CategoryID: gofakeit.Uint16(),
		Count:      gofakeit.Uint16(),
		CountryID:  gofakeit.Uint16(),
		EngName:    gofakeit.CarModel(),
		MarkaID:    gofakeit.Uint16(),
		BrandName:  gofakeit.CarModel(),
		Slang:      gofakeit.CarModel(),
		Value:      gofakeit.Uint16(),
	},
}

var InvalidJsonCases = []InvalidJsonTestCase{
	{
		CaseName: "Invalid categoryID",
		Brand: map[string]any{
			"brand": map[string]any{
				"category_id": "invalid",
				"count":       gofakeit.Uint16(),
				"country_id":  gofakeit.Uint16(),
				"eng_name":    gofakeit.CarModel(),
				"marka_id":    gofakeit.Uint16(),
				"name":        gofakeit.CarModel(),
				"slang":       gofakeit.CarModel(),
				"value":       gofakeit.Uint16(),
			},
		},
		Error: "Failed to decode request",
	},
	{
		CaseName: "Invalid slang",
		Brand: map[string]any{
			"brand": map[string]any{
				"category_id": "invalid",
				"count":       gofakeit.Uint16(),
				"country_id":  gofakeit.Uint16(),
				"eng_name":    gofakeit.CarModel(),
				"marka_id":    gofakeit.Uint16(),
				"name":        gofakeit.CarModel(),
				"slang":       123,
				"value":       gofakeit.Uint16(),
			},
		},
		Error: "Failed to decode request",
	},
}
