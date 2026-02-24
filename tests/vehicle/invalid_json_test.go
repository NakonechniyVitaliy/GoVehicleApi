package vehicle

import (
	"fmt"
	"net/http"
	"testing"

	testHelper "github.com/NakonechniyVitaliy/GoVehicleApi/tests/helper"
	"github.com/gavv/httpexpect/v2"
)

func TestInvalidJsonTest(t *testing.T) {
	e := httpexpect.Default(t, testHelper.TcUrl.String())

	for _, tc := range NoFieldsCases {
		tc := tc

		t.Run(tc.CaseName, func(t *testing.T) {
			doTestSaveInvalidJSON(e, tc)
			doTestUpdateInvalidJSON(e, tc)
		})
	}

	for _, tc := range InvalidJSONCases {
		tc := tc

		t.Run(tc.CaseName, func(t *testing.T) {
			doTestSaveInvalidJSON(e, tc)
			doTestUpdateInvalidJSON(e, tc)
		})
	}

}

func doTestSaveInvalidJSON(e *httpexpect.Expect, tc InvalidJsonTestCase) {
	req := e.POST("/vehicle/").
		WithJSON(tc.Vehicle).Expect()

	req.Status(http.StatusBadRequest)
	resp := req.JSON().Object()

	resp.Value("error").String().IsEqual(tc.Error)
}

func doTestUpdateInvalidJSON(e *httpexpect.Expect, tc InvalidJsonTestCase) {
	vehicleMap := tc.Vehicle["vehicle"].(map[string]any)

	req := e.PUT(fmt.Sprintf("/vehicle/1")).
		WithJSON(map[string]any{
			"vehicle": map[string]any{
				"brand":       vehicleMap["brand"],
				"driver_type": vehicleMap["driver_type"],
				"gearbox":     vehicleMap["gearbox"],
				"body_style":  vehicleMap["body_style"],
				"category":    vehicleMap["category"],
				"mileage":     vehicleMap["mileage"],
				"model":       vehicleMap["model"],
				"price":       vehicleMap["price"],
			},
		}).Expect()

	req.Status(http.StatusBadRequest)
	resp := req.JSON().Object()

	resp.Value("error").String().IsEqual(tc.Error)

}
