package driver_type

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
	e.PUT("/driver-type/refresh").Expect().Status(http.StatusOK)
}

func doTestGetAllPositive(e *httpexpect.Expect) {
	obj := e.GET("/driver-type/all").Expect().Status(http.StatusOK).JSON().Object()
	obj.Value("DriverTypes").Array().NotEmpty()
}
