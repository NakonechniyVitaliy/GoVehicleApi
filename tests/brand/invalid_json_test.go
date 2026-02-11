package brand

import (
	"net/http"
	"testing"

	"github.com/NakonechniyVitaliy/GoVehicleApi/tests/helper"
	"github.com/gavv/httpexpect/v2"
)

func TestInvalidJsonTest(t *testing.T) {
	e := httpexpect.Default(t, helper.TcUrl.String())

	for _, tc := range InvalidJsonCases {
		tc := tc

		t.Run(tc.CaseName, func(t *testing.T) {
			doTestSaveInvalidJSON(e, tc)

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
