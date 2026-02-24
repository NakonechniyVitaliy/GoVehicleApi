package body_style

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
	req := e.POST("/body-style/").
		WithJSON(tc.BodyStyle).Expect()

	req.Status(http.StatusBadRequest)
	resp := req.JSON().Object()

	resp.Value("error").String().IsEqual(tc.Error)
}

func doTestUpdateInvalidJSON(e *httpexpect.Expect, tc InvalidJsonTestCase) {
	bodyStyleMap := tc.BodyStyle["body_style"].(map[string]any)

	req := e.PUT(fmt.Sprintf("/body-style/1")).
		WithJSON(map[string]any{
			"brand": map[string]any{
				"name":        bodyStyleMap["name"],
				"value":       bodyStyleMap["value"],
			},
		}).Expect()

	req.Status(http.StatusBadRequest)
	resp := req.JSON().Object()

	resp.Value("error").String().IsEqual(tc.Error)

}
