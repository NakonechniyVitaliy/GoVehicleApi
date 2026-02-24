package category

import (
	"net/http"
	"testing"

	testHelper "github.com/NakonechniyVitaliy/GoVehicleApi/tests/helper"
	"github.com/gavv/httpexpect/v2"
)

func TestPositiveTests(t *testing.T) {

	t.Run("RefreshAndGetAll", func(t *testing.T) {
		e := httpexpect.Default(t, testHelper.TcUrl.String())

		doTestRefreshPositive(e)
		doTestGetAllPositive(e)
	})
}

func doTestRefreshPositive(e *httpexpect.Expect) {
	e.PUT("/category/refresh").Expect().Status(http.StatusOK)
}

func doTestGetAllPositive(e *httpexpect.Expect) {
	obj := e.GET("/category/all").Expect().Status(http.StatusOK).JSON().Object()
	obj.Value("BodyStyles").Array().NotEmpty()
}
