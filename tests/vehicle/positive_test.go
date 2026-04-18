package vehicle

import (
	"fmt"
	"net/http"
	"testing"

	dto "github.com/NakonechniyVitalii/GoVehicleApi/internal/http-server/dto/vehicle"
	"github.com/NakonechniyVitalii/GoVehicleApi/internal/services/helper"
	testHelper "github.com/NakonechniyVitalii/GoVehicleApi/tests/helper"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/gavv/httpexpect/v2"
)

func TestPositiveTests(t *testing.T) {
	e := testHelper.NewExpect(t)

	e.PUT("/brand/refresh").Expect().Status(http.StatusOK)
	e.PUT("/body-style/refresh").Expect().Status(http.StatusOK)
	e.PUT("/category/refresh").Expect().Status(http.StatusOK)
	e.PUT("/driver-type/refresh").Expect().Status(http.StatusOK)
	e.PUT("/gearbox/refresh").Expect().Status(http.StatusOK)

	brandID := fetchFirstID(e, "/brand/all", "Brands")
	bodyStyleID := fetchFirstID(e, "/body-style/all", "BodyStyles")
	categoryID := fetchFirstID(e, "/category/all", "Categories")
	driverTypeID := fetchFirstID(e, "/driver-type/all", "DriverTypes")
	gearboxID := fetchFirstID(e, "/gearbox/all", "Gearboxes")

	for _, tc := range PositiveCases {
		tc := tc
		tc.Brand = helper.PtrUint16(brandID)
		tc.DriverType = helper.PtrUint16(driverTypeID)
		tc.Gearbox = helper.PtrUint16(gearboxID)
		tc.BodyStyle = helper.PtrUint16(bodyStyleID)
		tc.Category = helper.PtrUint16(categoryID)

		original := dto.VehicleDTO{
			Brand:      tc.Brand,
			DriverType: tc.DriverType,
			Gearbox:    tc.Gearbox,
			BodyStyle:  tc.BodyStyle,
			Category:   tc.Category,
			Mileage:    tc.Mileage,
			Model:      tc.Model,
			Price:      tc.Price,
		}

		updatedVehicleData := dto.VehicleDTO{
			Brand:      helper.PtrUint16(brandID),
			DriverType: helper.PtrUint16(driverTypeID),
			Gearbox:    helper.PtrUint16(gearboxID),
			BodyStyle:  helper.PtrUint16(bodyStyleID),
			Category:   helper.PtrUint16(categoryID),
			Mileage:    helper.PtrUint32(gofakeit.Uint32()),
			Model:      helper.PtrString(gofakeit.CarModel()),
			Price:      helper.PtrUint32(gofakeit.Uint32()),
		}

		t.Run(tc.CaseName, func(t *testing.T) {
			e := testHelper.NewExpect(t)

			vehicleID := doTestSave(e, tc)
			doTestGetPositive(e, original, vehicleID)
			doTestUpdatePositive(e, updatedVehicleData, vehicleID)
			doTestGetPositive(e, updatedVehicleData, vehicleID)
			doTestGetAllPositive(e)
			doTestDeletePositive(e, vehicleID)
		})
	}
}

func fetchFirstID(e *httpexpect.Expect, path string, key string) uint16 {
	obj := e.GET(path).Expect().Status(http.StatusOK).JSON().Object()
	return uint16(obj.Value(key).Array().Value(0).Object().Value("id").Number().Raw())
}

func doTestSave(e *httpexpect.Expect, tc PositiveTestCase) uint16 {
	resp := e.POST("/vehicle/").
		WithJSON(dto.SaveRequest{
			Vehicle: dto.VehicleDTO{
				Brand:      tc.Brand,
				DriverType: tc.DriverType,
				Gearbox:    tc.Gearbox,
				BodyStyle:  tc.BodyStyle,
				Category:   tc.Category,
				Mileage:    tc.Mileage,
				Model:      tc.Model,
				Price:      tc.Price,
			},
		}).Expect().Status(http.StatusOK).
		JSON().Object()

	vehicle := resp.Value("Vehicle").Object()

	vehicle.Value("brand").Number().IsEqual(helper.DerefUint16(tc.Brand))
	vehicle.Value("mileage").Number().IsEqual(helper.DerefUint32(tc.Mileage))
	vehicle.Value("model").String().IsEqual(helper.DerefString(tc.Model))

	return uint16(vehicle.Value("id").Number().Raw())
}

func doTestGetPositive(e *httpexpect.Expect, expected dto.VehicleDTO, vehicleID uint16) {
	resp := e.GET(fmt.Sprintf("/vehicle/%d", vehicleID)).Expect().
		Status(http.StatusOK).JSON().Object()

	vehicle := resp.Value("Vehicle").Object()

	vehicle.Value("brand").Number().IsEqual(helper.DerefUint16(expected.Brand))
	vehicle.Value("mileage").Number().IsEqual(helper.DerefUint32(expected.Mileage))
	vehicle.Value("model").String().IsEqual(helper.DerefString(expected.Model))

}

func doTestUpdatePositive(e *httpexpect.Expect, updatedVehicleData dto.VehicleDTO, vehicleID uint16) {

	resp := e.PUT(fmt.Sprintf("/vehicle/%d", vehicleID)).
		WithJSON(dto.UpdateRequest{
			Vehicle: updatedVehicleData,
		}).Expect().Status(http.StatusOK).
		JSON().Object()

	vehicle := resp.Value("Vehicle").Object()

	vehicle.Value("brand").Number().IsEqual(helper.DerefUint16(updatedVehicleData.Brand))
	vehicle.Value("mileage").Number().IsEqual(helper.DerefUint32(updatedVehicleData.Mileage))
	vehicle.Value("model").String().IsEqual(helper.DerefString(updatedVehicleData.Model))

}

func doTestDeletePositive(e *httpexpect.Expect, vehicleID uint16) {
	e.DELETE(fmt.Sprintf("/vehicle/%d", vehicleID)).Expect().Status(http.StatusOK)

	e.GET(fmt.Sprintf("/vehicle/%d", vehicleID)).Expect().
		Status(http.StatusNotFound)

}

func doTestGetAllPositive(e *httpexpect.Expect) {
	obj := e.GET("/vehicle/").Expect().Status(http.StatusOK).JSON().Object()
	obj.Value("Vehicles").Array().NotEmpty()
}
