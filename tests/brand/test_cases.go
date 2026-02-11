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
		CaseName: "Invalid CategoryID",
		Brand: map[string]any{
			"brand": map[string]any{
				"category_id": "invalid",
				"cnt":         gofakeit.Uint16(),
				"country_id":  gofakeit.Uint16(),
				"eng":         gofakeit.CarModel(),
				"marka_id":    gofakeit.Uint16(),
				"name":        gofakeit.CarModel(),
				"slang":       gofakeit.CarModel(),
				"value":       gofakeit.Uint16(),
			},
		},
		Error: "Failed to decode request",
	},
	{
		CaseName: "Invalid Count",
		Brand: map[string]any{
			"brand": map[string]any{
				"category_id": gofakeit.Uint16(),
				"cnt":         "invalid",
				"country_id":  gofakeit.Uint16(),
				"eng":         gofakeit.CarModel(),
				"marka_id":    gofakeit.Uint16(),
				"name":        gofakeit.CarModel(),
				"slang":       gofakeit.CarModel(),
				"value":       gofakeit.Uint16(),
			},
		},
		Error: "Failed to decode request",
	},
	{
		CaseName: "Invalid CountryID",
		Brand: map[string]any{
			"brand": map[string]any{
				"category_id": gofakeit.Uint16(),
				"cnt":         gofakeit.Uint16(),
				"country_id":  "invalid",
				"eng":         gofakeit.CarModel(),
				"marka_id":    gofakeit.Uint16(),
				"name":        gofakeit.CarModel(),
				"slang":       gofakeit.CarModel(),
				"value":       gofakeit.Uint16(),
			},
		},
		Error: "Failed to decode request",
	},
	{
		CaseName: "Invalid EngName",
		Brand: map[string]any{
			"brand": map[string]any{
				"category_id": gofakeit.Uint16(),
				"cnt":         gofakeit.Uint16(),
				"country_id":  gofakeit.Uint16(),
				"eng":         123,
				"marka_id":    gofakeit.Uint16(),
				"name":        gofakeit.CarModel(),
				"slang":       gofakeit.CarModel(),
				"value":       gofakeit.Uint16(),
			},
		},
		Error: "Failed to decode request",
	},
	{
		CaseName: "Invalid MarkaID",
		Brand: map[string]any{
			"brand": map[string]any{
				"category_id": gofakeit.Uint16(),
				"cnt":         gofakeit.Uint16(),
				"country_id":  gofakeit.Uint16(),
				"eng":         gofakeit.CarModel(),
				"marka_id":    "invalid",
				"name":        gofakeit.CarModel(),
				"slang":       gofakeit.CarModel(),
				"value":       gofakeit.Uint16(),
			},
		},
		Error: "Failed to decode request",
	},
	{
		CaseName: "Invalid Name",
		Brand: map[string]any{
			"brand": map[string]any{
				"category_id": gofakeit.Uint16(),
				"cnt":         gofakeit.Uint16(),
				"country_id":  gofakeit.Uint16(),
				"eng":         gofakeit.CarModel(),
				"marka_id":    gofakeit.Uint16(),
				"name":        123,
				"slang":       gofakeit.CarModel(),
				"value":       gofakeit.Uint16(),
			},
		},
		Error: "Failed to decode request",
	},
	{
		CaseName: "Invalid Slang",
		Brand: map[string]any{
			"brand": map[string]any{
				"category_id": gofakeit.Uint16(),
				"cnt":         gofakeit.Uint16(),
				"country_id":  gofakeit.Uint16(),
				"eng":         gofakeit.CarModel(),
				"marka_id":    gofakeit.Uint16(),
				"name":        gofakeit.CarModel(),
				"slang":       123,
				"value":       gofakeit.Uint16(),
			},
		},
		Error: "Failed to decode request",
	},
	{
		CaseName: "Invalid Value",
		Brand: map[string]any{
			"brand": map[string]any{
				"category_id": gofakeit.Uint16(),
				"cnt":         gofakeit.Uint16(),
				"country_id":  gofakeit.Uint16(),
				"eng":         gofakeit.CarModel(),
				"marka_id":    gofakeit.Uint16(),
				"name":        gofakeit.CarModel(),
				"slang":       gofakeit.CarModel(),
				"value":       "invalid",
			},
		},
		Error: "Failed to decode request",
	},
	{
		CaseName: "Empty body",
		Brand: map[string]any{
			"brand": map[string]any{},
		},
		Error: "Failed to decode request",
	},
}
