package brand

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
	req := e.POST("/brand/").
		WithJSON(tc.Brand).Expect()

	req.Status(http.StatusBadRequest)
	resp := req.JSON().Object()

	resp.Value("error").String().IsEqual(tc.Error)
}

func doTestUpdateInvalidJSON(e *httpexpect.Expect, tc InvalidJsonTestCase) {
	brandMap := tc.Brand["brand"].(map[string]any)

	req := e.PUT(fmt.Sprintf("/brand/1")).
		WithJSON(map[string]any{
			"brand": map[string]any{
				"category_id": brandMap["category_id"],
				"cnt":         brandMap["cnt"],
				"country_id":  brandMap["country_id"],
				"eng":         brandMap["eng"],
				"marka_id":    brandMap["marka_id"],
				"name":        brandMap["name"],
				"slang":       brandMap["slang"],
				"value":       brandMap["value"],
			},
		}).Expect()

	req.Status(http.StatusBadRequest)
	resp := req.JSON().Object()

	resp.Value("error").String().IsEqual(tc.Error)

}
