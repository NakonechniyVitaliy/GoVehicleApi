package brand

import (
	"net/http"
	"testing"

	"github.com/NakonechniyVitaliy/GoVehicleApi/tests/helper"
	"github.com/gavv/httpexpect/v2"
)

func TestInvalidJsonTest(t *testing.T) {

	for _, tc := range InvalidJsonCases {
		tc := tc

		t.Run(tc.CaseName, func(t *testing.T) {
			e := httpexpect.Default(t, helper.TcUrl.String())

			doTestSaveInvalidJSON(e, tc)

		})
	}
}

func doTestSaveInvalidJSON(e *httpexpect.Expect, tc InvalidJsonTestCase) {
	resp := e.POST("/brand/").
		WithJSON(tc.Brand).Expect().Status(http.StatusBadRequest).JSON().Object()

	resp.Value("error").String().IsEqual(tc.Error)
}
