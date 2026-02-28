package body_style

import (
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/service/helper"
	"github.com/brianvoe/gofakeit/v6"
)

type PositiveTestCase struct {
	CaseName string
	Name     *string
	Value    *uint16
}
type InvalidJsonTestCase struct {
	CaseName  string
	Error     string
	BodyStyle map[string]any
}

var PositiveCases = []PositiveTestCase{
	{
		CaseName: "Valid body style",
		Name:     helper.PtrString(gofakeit.CarType()),
		Value:    helper.PtrUint16(gofakeit.Uint16()),
	},
}

var InvalidJSONCases = []InvalidJsonTestCase{
	{
		CaseName: "Invalid Name",
		BodyStyle: map[string]any{
			"bodyStyle": map[string]any{
				"name":  123,
				"value": gofakeit.Uint16(),
			},
		},
		Error: "invalid JSON or wrong field types",
	},
	{
		CaseName: "Invalid Value",
		BodyStyle: map[string]any{
			"bodyStyle": map[string]any{
				"name":  gofakeit.CarType(),
				"value": "invalid",
			},
		},
		Error: "invalid JSON or wrong field types",
	},
}

var NoFieldsCases = []InvalidJsonTestCase{
	{
		CaseName: "No Name field",
		BodyStyle: map[string]any{
			"bodyStyle": map[string]any{
				"value": gofakeit.Uint16(),
			},
		},
		Error: "all fields are required",
	},
	{
		CaseName: "No Value field",
		BodyStyle: map[string]any{
			"bodyStyle": map[string]any{
				"name": gofakeit.CarType(),
			},
		},
		Error: "all fields are required",
	},
	{
		CaseName: "Empty body",
		BodyStyle: map[string]any{
			"bodyStyle": map[string]any{},
		},
		Error: "all fields are required",
	},
}
