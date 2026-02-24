package vehicle

import (
	"fmt"
	"net/http"
	"testing"

	dto "github.com/NakonechniyVitaliy/GoVehicleApi/internal/http-server/dto/vehicle"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/service/helper"
	testHelper "github.com/NakonechniyVitaliy/GoVehicleApi/tests/helper"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/gavv/httpexpect/v2"
)

func TestPositiveTests(t *testing.T) {

	for _, tc := range PositiveCases {
		tc := tc

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
			Brand:      helper.PtrUint16(gofakeit.Uint16()),
			DriverType: helper.PtrUint16(gofakeit.Uint16()),
			Gearbox:    helper.PtrUint16(gofakeit.Uint16()),
			BodyStyle:  helper.PtrUint16(gofakeit.Uint16()),
			Category:   helper.PtrUint16(gofakeit.Uint16()),
			Mileage:    helper.PtrUint32(gofakeit.Uint32()),
			Model:      helper.PtrString(gofakeit.CarModel()),
			Price:      helper.PtrUint16(gofakeit.Uint16()),
		}

		t.Run(tc.CaseName, func(t *testing.T) {
			e := httpexpect.Default(t, testHelper.TcUrl.String())

			vehicleID := doTestSave(e, tc)
			doTestGetPositive(e, original, vehicleID)
			doTestUpdatePositive(e, updatedVehicleData, vehicleID)
			doTestGetPositive(e, updatedVehicleData, vehicleID)
			doTestDeletePositive(e, vehicleID)
			doTestRefreshPositive(e)
			doTestGetAllPositive(e)
		})
	}
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

func doTestRefreshPositive(e *httpexpect.Expect) {
	e.PUT("/vehicle/refresh").Expect().Status(http.StatusOK)
}

func doTestGetAllPositive(e *httpexpect.Expect) {
	obj := e.GET("/vehicle/all").Expect().Status(http.StatusOK).JSON().Object()
	obj.Value("Vehicles").Array().NotEmpty()
}
