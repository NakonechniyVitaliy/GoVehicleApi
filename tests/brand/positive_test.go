package brand

import (
	"fmt"
	"net/http"
	"testing"

	handler "github.com/NakonechniyVitaliy/GoVehicleApi/internal/http-server/handlers/brand"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/models"
	"github.com/NakonechniyVitaliy/GoVehicleApi/tests/helper"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/gavv/httpexpect/v2"
)

func TestPositiveTests(t *testing.T) {

	for _, tc := range PositiveCases {
		tc := tc

		original := models.Brand{
			CategoryID: tc.CategoryID,
			Count:      tc.Count,
			CountryID:  tc.CountryID,
			EngName:    tc.EngName,
			MarkaID:    tc.MarkaID,
			Name:       tc.BrandName,
			Slang:      tc.Slang,
			Value:      tc.Value,
		}

		updatedBrandData := models.Brand{
			CategoryID: gofakeit.Uint16(),
			Count:      gofakeit.Uint16(),
			CountryID:  gofakeit.Uint16(),
			EngName:    gofakeit.CarModel(),
			MarkaID:    gofakeit.Uint16(),
			Name:       gofakeit.CarModel(),
			Slang:      gofakeit.CarModel(),
			Value:      gofakeit.Uint16(),
		}

		t.Run(tc.CaseName, func(t *testing.T) {
			e := httpexpect.Default(t, helper.TcUrl.String())

			brandID := doTestSave(e, tc)
			doTestGetPositive(e, original, brandID)
			doTestUpdatePositive(e, updatedBrandData, brandID)
			doTestGetPositive(e, updatedBrandData, brandID)
			doTestDeletePositive(e, brandID)
			doTestRefreshPositive(e)
			doTestGetAllPositive(e)
		})
	}
}

func doTestSave(e *httpexpect.Expect, tc PositiveTestCase) uint16 {
	resp := e.POST("/brand/").
		WithJSON(handler.SaveRequest{
			Brand: models.Brand{
				CategoryID: tc.CategoryID,
				Count:      tc.Count,
				CountryID:  tc.CountryID,
				EngName:    tc.EngName,
				MarkaID:    tc.MarkaID,
				Name:       tc.BrandName,
				Slang:      tc.Slang,
				Value:      tc.Value,
			},
		}).Expect().Status(http.StatusOK).
		JSON().Object()

	brand := resp.Value("Brand").Object()

	brand.Value("name").String().IsEqual(tc.BrandName)
	brand.Value("eng").String().IsEqual(tc.EngName)
	brand.Value("marka_id").Number().IsEqual(float64(tc.MarkaID))

	return uint16(brand.Value("id").Number().Raw())
}

func doTestGetPositive(e *httpexpect.Expect, expected models.Brand, brandID uint16) {
	resp := e.GET(fmt.Sprintf("/brand/%d", brandID)).Expect().
		Status(http.StatusOK).JSON().Object()

	brand := resp.Value("Brand").Object()

	brand.Value("name").String().IsEqual(expected.Name)
	brand.Value("eng").String().IsEqual(expected.EngName)
	brand.Value("marka_id").Number().IsEqual(float64(expected.MarkaID))

}

func doTestUpdatePositive(e *httpexpect.Expect, updatedBrandData models.Brand, brandID uint16) {

	resp := e.PUT(fmt.Sprintf("/brand/%d", brandID)).
		WithJSON(handler.UpdateRequest{
			Brand: updatedBrandData,
		}).Expect().Status(http.StatusOK).
		JSON().Object()

	brand := resp.Value("Brand").Object()

	brand.Value("name").String().IsEqual(updatedBrandData.Name)
	brand.Value("eng").String().IsEqual(updatedBrandData.EngName)
	brand.Value("marka_id").Number().IsEqual(float64(updatedBrandData.MarkaID))

}

func doTestDeletePositive(e *httpexpect.Expect, brandID uint16) {
	e.DELETE(fmt.Sprintf("/brand/%d", brandID)).Expect().Status(http.StatusOK)

	e.GET(fmt.Sprintf("/brand/%d", brandID)).Expect().
		Status(http.StatusNotFound)

}

func doTestRefreshPositive(e *httpexpect.Expect) {
	e.PUT("/brand/refresh").Expect().Status(http.StatusOK)
}

func doTestGetAllPositive(e *httpexpect.Expect) {
	obj := e.GET("/brand/all").Expect().Status(http.StatusOK).JSON().Object()
	obj.Value("Brands").Array().NotEmpty()
}
