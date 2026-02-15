package brand

import (
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/service/helper"
	"github.com/brianvoe/gofakeit/v6"
)

type PositiveTestCase struct {
	CaseName   string
	CategoryID *uint16
	Count      *uint16
	CountryID  *uint16
	EngName    *string
	MarkaID    *uint16
	BrandName  *string
	Slang      *string
	Value      *uint16
}
type InvalidJsonTestCase struct {
	CaseName string
	Error    string
	Brand    map[string]any
}

var PositiveCases = []PositiveTestCase{
	{
		CaseName:   "Valid brand",
		CategoryID: helper.PtrUint16(gofakeit.Uint16()),
		Count:      helper.PtrUint16(gofakeit.Uint16()),
		CountryID:  helper.PtrUint16(gofakeit.Uint16()),
		EngName:    helper.PtrString(gofakeit.CarModel()),
		MarkaID:    helper.PtrUint16(gofakeit.Uint16()),
		BrandName:  helper.PtrString(gofakeit.CarModel()),
		Slang:      helper.PtrString(gofakeit.CarModel()),
		Value:      helper.PtrUint16(gofakeit.Uint16()),
	},
}

func returnBrandMapWithoutOneField(field string) map[string]any {
	brand := map[string]any{
		"brand": map[string]any{
			"category_id": gofakeit.Uint16(),
			"cnt":         gofakeit.Uint16(),
			"country_id":  gofakeit.Uint16(),
			"eng":         gofakeit.CarModel(),
			"marka_id":    gofakeit.Uint16(),
			"name":        gofakeit.CarModel(),
			"slang":       gofakeit.CarModel(),
			"value":       gofakeit.Uint16(),
		},
	}
	inner := brand["brand"].(map[string]any)
	delete(inner, field)

	return brand
}

var InvalidJsonCases = []InvalidJsonTestCase{
	{
		CaseName: "No CategoryID field",
		Brand:    returnBrandMapWithoutOneField("category_id"),
		Error:    "all fields are required",
	},
	{
		CaseName: "No Count field",
		Brand:    returnBrandMapWithoutOneField("cnt"),
		Error:    "all fields are required",
	},
	{
		CaseName: "No CountryID field",
		Brand:    returnBrandMapWithoutOneField("country_id"),
		Error:    "all fields are required",
	},
	{
		CaseName: "No EnglishName field",
		Brand:    returnBrandMapWithoutOneField("eng"),
		Error:    "all fields are required",
	},
	{
		CaseName: "No MarkaID field",
		Brand:    returnBrandMapWithoutOneField("marka_id"),
		Error:    "all fields are required",
	},
	{
		CaseName: "No Name field",
		Brand:    returnBrandMapWithoutOneField("name"),
		Error:    "all fields are required",
	},
	{
		CaseName: "No Slang field",
		Brand:    returnBrandMapWithoutOneField("slang"),
		Error:    "all fields are required",
	},
	{
		CaseName: "No Value field",
		Brand:    returnBrandMapWithoutOneField("value"),
		Error:    "all fields are required",
	},
	{
		CaseName: "Empty body",
		Brand: map[string]any{
			"brand": map[string]any{},
		},
		Error: "all fields are required",
	},
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
		Error: "invalid JSON or wrong field types",
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
		Error: "invalid JSON or wrong field types",
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
		Error: "invalid JSON or wrong field types",
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
		Error: "invalid JSON or wrong field types",
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
		Error: "invalid JSON or wrong field types",
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
		Error: "invalid JSON or wrong field types",
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
		Error: "invalid JSON or wrong field types",
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
		Error: "invalid JSON or wrong field types",
	},
}
