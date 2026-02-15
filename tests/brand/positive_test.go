package brand

import (
	"fmt"
	"net/http"
	"testing"

	dto "github.com/NakonechniyVitaliy/GoVehicleApi/internal/http-server/dto/brand"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/service/helper"
	testHelper "github.com/NakonechniyVitaliy/GoVehicleApi/tests/helper"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/gavv/httpexpect/v2"
)

func TestPositiveTests(t *testing.T) {

	for _, tc := range PositiveCases {
		tc := tc

		original := dto.BrandDTO{
			CategoryID: tc.CategoryID,
			Count:      tc.Count,
			CountryID:  tc.CountryID,
			EngName:    tc.EngName,
			MarkaID:    tc.MarkaID,
			Name:       tc.BrandName,
			Slang:      tc.Slang,
			Value:      tc.Value,
		}

		updatedBrandData := dto.BrandDTO{
			CategoryID: helper.PtrUint16(gofakeit.Uint16()),
			Count:      helper.PtrUint16(gofakeit.Uint16()),
			CountryID:  helper.PtrUint16(gofakeit.Uint16()),
			EngName:    helper.PtrString(gofakeit.CarModel()),
			MarkaID:    helper.PtrUint16(gofakeit.Uint16()),
			Name:       helper.PtrString(gofakeit.CarModel()),
			Slang:      helper.PtrString(gofakeit.CarModel()),
			Value:      helper.PtrUint16(gofakeit.Uint16()),
		}

		t.Run(tc.CaseName, func(t *testing.T) {
			e := httpexpect.Default(t, testHelper.TcUrl.String())

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
		WithJSON(dto.SaveRequest{
			Brand: dto.BrandDTO{
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

	brand.Value("name").String().IsEqual(helper.DerefString(tc.BrandName))
	brand.Value("eng").String().IsEqual(helper.DerefString(tc.EngName))
	brand.Value("marka_id").Number().IsEqual(float64(*tc.MarkaID))

	return uint16(brand.Value("id").Number().Raw())
}

func doTestGetPositive(e *httpexpect.Expect, expected dto.BrandDTO, brandID uint16) {
	resp := e.GET(fmt.Sprintf("/brand/%d", brandID)).Expect().
		Status(http.StatusOK).JSON().Object()

	brand := resp.Value("Brand").Object()

	brand.Value("name").String().IsEqual(helper.DerefString(expected.Name))
	brand.Value("eng").String().IsEqual(helper.DerefString(expected.EngName))
	brand.Value("marka_id").Number().IsEqual(float64(*expected.MarkaID))

}

func doTestUpdatePositive(e *httpexpect.Expect, updatedBrandData dto.BrandDTO, brandID uint16) {

	resp := e.PUT(fmt.Sprintf("/brand/%d", brandID)).
		WithJSON(dto.UpdateRequest{
			Brand: updatedBrandData,
		}).Expect().Status(http.StatusOK).
		JSON().Object()

	brand := resp.Value("Brand").Object()

	brand.Value("name").String().IsEqual(helper.DerefString(updatedBrandData.Name))
	brand.Value("eng").String().IsEqual(helper.DerefString(updatedBrandData.EngName))
	brand.Value("marka_id").Number().IsEqual(float64(*updatedBrandData.MarkaID))

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
