package vehicle

import (
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/service/helper"
	"github.com/brianvoe/gofakeit/v6"
)

type PositiveTestCase struct {
	CaseName   string
	Brand      *uint16
	DriverType *uint16
	Gearbox    *uint16
	BodyStyle  *uint16
	Category   *uint16
	Mileage    *uint32
	Model      *string
	Price      *uint16
}
type InvalidJsonTestCase struct {
	CaseName string
	Error    string
	Vehicle  map[string]any
}

var PositiveCases = []PositiveTestCase{
	{
		CaseName:   "Valid vehicle",
		Brand:      helper.PtrUint16(gofakeit.Uint16()),
		DriverType: helper.PtrUint16(gofakeit.Uint16()),
		Gearbox:    helper.PtrUint16(gofakeit.Uint16()),
		BodyStyle:  helper.PtrUint16(gofakeit.Uint16()),
		Category:   helper.PtrUint16(gofakeit.Uint16()),
		Mileage:    helper.PtrUint32(gofakeit.Uint32()),
		Model:      helper.PtrString(gofakeit.CarModel()),
		Price:      helper.PtrUint16(gofakeit.Uint16()),
	},
}

func returnVehicleMapWithoutOneField(field string) map[string]any {
	vehicle := map[string]any{
		"vehicle": map[string]any{
			"brand":       gofakeit.Uint16(),
			"driver_type": gofakeit.Uint16(),
			"gearbox":     gofakeit.Uint16(),
			"body_style":  gofakeit.Uint16(),
			"category":    gofakeit.Uint16(),
			"mileage":     gofakeit.Uint32(),
			"model":       gofakeit.CarModel(),
			"price":       gofakeit.Uint16(),
		},
	}
	inner := vehicle["vehicle"].(map[string]any)
	delete(inner, field)

	return vehicle
}

var InvalidJSONCases = []InvalidJsonTestCase{
	{
		CaseName: "Invalid brand",
		Vehicle: map[string]any{
			"vehicle": map[string]any{
				"brand":       "invalid",
				"driver_type": gofakeit.Uint16(),
				"gearbox":     gofakeit.Uint16(),
				"body_style":  gofakeit.Uint16(),
				"category":    gofakeit.Uint16(),
				"mileage":     gofakeit.Uint32(),
				"model":       gofakeit.CarModel(),
				"price":       gofakeit.Uint16(),
			},
		},
		Error: "invalid JSON or wrong field types",
	},
	{
		CaseName: "Invalid driver type",
		Vehicle: map[string]any{
			"vehicle": map[string]any{
				"brand":       gofakeit.Uint16(),
				"driver_type": "invalid",
				"gearbox":     gofakeit.Uint16(),
				"body_style":  gofakeit.Uint16(),
				"category":    gofakeit.Uint16(),
				"mileage":     gofakeit.Uint32(),
				"model":       gofakeit.CarModel(),
				"price":       gofakeit.Uint16(),
			},
		},
		Error: "invalid JSON or wrong field types",
	},
	{
		CaseName: "Invalid gearbox",
		Vehicle: map[string]any{
			"vehicle": map[string]any{
				"brand":       gofakeit.Uint16(),
				"driver_type": gofakeit.Uint16(),
				"gearbox":     "invalid",
				"body_style":  gofakeit.Uint16(),
				"category":    gofakeit.Uint16(),
				"mileage":     gofakeit.Uint32(),
				"model":       gofakeit.CarModel(),
				"price":       gofakeit.Uint16(),
			},
		},
		Error: "invalid JSON or wrong field types",
	},
	{
		CaseName: "Invalid body style",
		Vehicle: map[string]any{
			"vehicle": map[string]any{
				"brand":       gofakeit.Uint16(),
				"driver_type": gofakeit.Uint16(),
				"gearbox":     gofakeit.Uint16(),
				"body_style":  "invalid",
				"category":    gofakeit.Uint16(),
				"mileage":     gofakeit.Uint32(),
				"model":       gofakeit.CarModel(),
				"price":       gofakeit.Uint16(),
			},
		},
		Error: "invalid JSON or wrong field types",
	},
	{
		CaseName: "Invalid category",
		Vehicle: map[string]any{
			"vehicle": map[string]any{
				"brand":       gofakeit.Uint16(),
				"driver_type": gofakeit.Uint16(),
				"gearbox":     gofakeit.Uint16(),
				"body_style":  gofakeit.Uint16(),
				"category":    "invalid",
				"mileage":     gofakeit.Uint32(),
				"model":       gofakeit.CarModel(),
				"price":       gofakeit.Uint16(),
			},
		},
		Error: "invalid JSON or wrong field types",
	},
	{
		CaseName: "Invalid mileage",
		Vehicle: map[string]any{
			"vehicle": map[string]any{
				"brand":       gofakeit.Uint16(),
				"driver_type": gofakeit.Uint16(),
				"gearbox":     gofakeit.Uint16(),
				"body_style":  gofakeit.Uint16(),
				"category":    gofakeit.Uint16(),
				"mileage":     "invalid",
				"model":       gofakeit.CarModel(),
				"price":       gofakeit.Uint16(),
			},
		},
		Error: "invalid JSON or wrong field types",
	},
	{
		CaseName: "Invalid model",
		Vehicle: map[string]any{
			"vehicle": map[string]any{
				"brand":       gofakeit.Uint16(),
				"driver_type": gofakeit.Uint16(),
				"gearbox":     gofakeit.Uint16(),
				"body_style":  gofakeit.Uint16(),
				"category":    gofakeit.Uint16(),
				"mileage":     gofakeit.Uint16(),
				"model":       123,
				"price":       gofakeit.Uint16(),
			},
		},
		Error: "invalid JSON or wrong field types",
	},
	{
		CaseName: "Invalid price",
		Vehicle: map[string]any{
			"vehicle": map[string]any{
				"brand":       gofakeit.Uint16(),
				"driver_type": gofakeit.Uint16(),
				"gearbox":     gofakeit.Uint16(),
				"body_style":  gofakeit.Uint16(),
				"category":    gofakeit.Uint16(),
				"mileage":     gofakeit.Uint16(),
				"model":       gofakeit.CarModel(),
				"price":       "invalid",
			},
		},
		Error: "invalid JSON or wrong field types",
	},
}

var NoFieldsCases = []InvalidJsonTestCase{
	{
		CaseName: "No CategoryID field",
		Vehicle:  returnVehicleMapWithoutOneField("brand"),
		Error:    "all fields are required",
	},
	{
		CaseName: "No Count field",
		Vehicle:  returnVehicleMapWithoutOneField("driver_type"),
		Error:    "all fields are required",
	},
	{
		CaseName: "No CountryID field",
		Vehicle:  returnVehicleMapWithoutOneField("gearbox"),
		Error:    "all fields are required",
	},
	{
		CaseName: "No EnglishName field",
		Vehicle:  returnVehicleMapWithoutOneField("body_style"),
		Error:    "all fields are required",
	},
	{
		CaseName: "No MarkaID field",
		Vehicle:  returnVehicleMapWithoutOneField("category"),
		Error:    "all fields are required",
	},
	{
		CaseName: "No Name field",
		Vehicle:  returnVehicleMapWithoutOneField("mileage"),
		Error:    "all fields are required",
	},
	{
		CaseName: "No Slang field",
		Vehicle:  returnVehicleMapWithoutOneField("model"),
		Error:    "all fields are required",
	},
	{
		CaseName: "No Value field",
		Vehicle:  returnVehicleMapWithoutOneField("price"),
		Error:    "all fields are required",
	},
	{
		CaseName: "Empty body",
		Vehicle: map[string]any{
			"vehicle": map[string]any{},
		},
		Error: "all fields are required",
	},
}
