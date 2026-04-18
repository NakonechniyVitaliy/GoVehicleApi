package gearbox

import (
	"net/http"
	"testing"

	testHelper "github.com/NakonechniyVitalii/GoVehicleApi/tests/helper"
	"github.com/gavv/httpexpect/v2"
)

func TestPositiveTests(t *testing.T) {

	t.Run("RefreshAndGetAll", func(t *testing.T) {
		e := testHelper.NewExpect(t)

		doTestRefreshPositive(e)
		doTestGetAllPositive(e)
	})
}

func doTestRefreshPositive(e *httpexpect.Expect) {
	e.PUT("/gearbox/refresh").Expect().Status(http.StatusOK)
}

func doTestGetAllPositive(e *httpexpect.Expect) {
	obj := e.GET("/gearbox/all").Expect().Status(http.StatusOK).JSON().Object()
	obj.Value("Gearboxes").Array().NotEmpty()
}
