package body_style

import (
	"fmt"
	"net/http"
	"testing"

	dto "github.com/NakonechniyVitaliy/GoVehicleApi/internal/http-server/dto/body_style"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/service/helper"
	testHelper "github.com/NakonechniyVitaliy/GoVehicleApi/tests/helper"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/gavv/httpexpect/v2"
)

func TestPositiveTests(t *testing.T) {

	for _, tc := range PositiveCases {
		tc := tc

		original := dto.BodyStyleDTO{
			Name:  tc.Name,
			Value: tc.Value,
		}

		updatedBSData := dto.BodyStyleDTO{
			Name:  helper.PtrString(gofakeit.CarType()),
			Value: helper.PtrUint16(gofakeit.Uint16()),
		}

		t.Run(tc.CaseName, func(t *testing.T) {
			e := httpexpect.Default(t, testHelper.TcUrl.String())

			bodyStyleID := doTestSave(e, tc)
			doTestGetPositive(e, original, bodyStyleID)
			doTestUpdatePositive(e, updatedBSData, bodyStyleID)
			doTestGetPositive(e, updatedBSData, bodyStyleID)
			doTestDeletePositive(e, bodyStyleID)
			doTestRefreshPositive(e)
			doTestGetAllPositive(e)
		})
	}
}

func doTestSave(e *httpexpect.Expect, tc PositiveTestCase) uint16 {
	resp := e.POST("/body-style/").
		WithJSON(dto.SaveRequest{
			BodyStyleDTO: dto.BodyStyleDTO{
				Name:       tc.Name,
				Value:      tc.Value,
			},
		}).Expect().Status(http.StatusOK).
		JSON().Object()

	bodyStyle := resp.Value("BodyStyle").Object()

	bodyStyle.Value("name").String().IsEqual(helper.DerefString(tc.Name))
	bodyStyle.Value("value").Number().IsEqual(helper.DerefUint16(tc.Value))

	return uint16(bodyStyle.Value("id").Number().Raw())
}

func doTestGetPositive(e *httpexpect.Expect, expected dto.BodyStyleDTO, bodyStyleID uint16) {
	resp := e.GET(fmt.Sprintf("/body-style/%d", bodyStyleID)).Expect().
		Status(http.StatusOK).JSON().Object()

	bodyStyle := resp.Value("BodyStyle").Object()

	bodyStyle.Value("name").String().IsEqual(helper.DerefString(expected.Name))
	bodyStyle.Value("value").Number().IsEqual(helper.DerefUint16(expected.Value))

}

func doTestUpdatePositive(e *httpexpect.Expect, updatedBSData dto.BodyStyleDTO, bodyStyleID uint16) {

	resp := e.PUT(fmt.Sprintf("/body-style/%d", bodyStyleID)).
		WithJSON(dto.UpdateRequest{
			BodyStyleDTO: updatedBSData,
		}).Expect().Status(http.StatusOK).
		JSON().Object()

	bodyStyle := resp.Value("BodyStyle").Object()

	bodyStyle.Value("name").String().IsEqual(helper.DerefString(updatedBSData.Name))
	bodyStyle.Value("value").Number().IsEqual(helper.DerefUint16(updatedBSData.Value))

}

func doTestDeletePositive(e *httpexpect.Expect, bodyStyleID uint16) {
	e.DELETE(fmt.Sprintf("/body-style/%d", bodyStyleID)).Expect().Status(http.StatusOK)

	e.GET(fmt.Sprintf("/body-style/%d", bodyStyleID)).Expect().
		Status(http.StatusNotFound)

}

func doTestRefreshPositive(e *httpexpect.Expect) {
	e.PUT("/body-style/refresh").Expect().Status(http.StatusOK)
}

func doTestGetAllPositive(e *httpexpect.Expect) {
	obj := e.GET("/body-style/all").Expect().Status(http.StatusOK).JSON().Object()
	obj.Value("BodyStyles").Array().NotEmpty()
}
