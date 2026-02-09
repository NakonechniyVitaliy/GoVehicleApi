package brand

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"

	handler "github.com/NakonechniyVitaliy/GoVehicleApi/internal/http-server/handlers/brand"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/models"
	"github.com/NakonechniyVitaliy/GoVehicleApi/tests"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/gavv/httpexpect/v2"
)

type testCase struct {
	CaseName   string
	CategoryID uint16
	Count      uint16
	CountryID  uint16
	EngName    string
	MarkaID    uint16
	BrandName  string
	Slang      string
	Value      uint16
}

func TestBrand(t *testing.T) {

	testCases := []testCase{
		{
			CaseName:   "Valid brand",
			CategoryID: gofakeit.Uint16(),
			Count:      gofakeit.Uint16(),
			CountryID:  gofakeit.Uint16(),
			EngName:    gofakeit.CarModel(),
			MarkaID:    gofakeit.Uint16(),
			BrandName:  gofakeit.CarModel(),
			Slang:      gofakeit.CarModel(),
			Value:      gofakeit.Uint16(),
		},
	}

	for _, tc := range testCases {
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

		tcUrl := url.URL{
			Scheme: "http",
			Host:   tests.LOCAL_HOST,
		}

		t.Run(tc.CaseName, func(t *testing.T) {
			e := httpexpect.Default(t, tcUrl.String())

			brandID := doTestSave(e, tc)
			doTestGet(e, original, brandID)
			doTestUpdate(e, updatedBrandData, brandID)
			doTestGet(e, updatedBrandData, brandID)
			doTestDelete(e, brandID)
			doTestRefresh(e)
			doTestGetAll(e)
		})
	}

}

func doTestSave(e *httpexpect.Expect, tc testCase) uint16 {
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

func doTestGet(e *httpexpect.Expect, expected models.Brand, brandID uint16) {
	resp := e.GET(fmt.Sprintf("/brand/%d", brandID)).Expect().
		Status(http.StatusOK).JSON().Object()

	brand := resp.Value("Brand").Object()

	brand.Value("name").String().IsEqual(expected.Name)
	brand.Value("eng").String().IsEqual(expected.EngName)
	brand.Value("marka_id").Number().IsEqual(float64(expected.MarkaID))

}

func doTestUpdate(e *httpexpect.Expect, updatedBrandData models.Brand, brandID uint16) {

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

func doTestDelete(e *httpexpect.Expect, brandID uint16) {
	e.DELETE(fmt.Sprintf("/brand/%d", brandID)).Expect().Status(http.StatusOK)

	e.GET(fmt.Sprintf("/brand/%d", brandID)).Expect().
		Status(http.StatusNotFound)

}

func doTestRefresh(e *httpexpect.Expect) {
	e.PUT("/brand/refresh").Expect().Status(http.StatusOK)
}

func doTestGetAll(e *httpexpect.Expect) {
	obj := e.GET("/brand/all").Expect().Status(http.StatusOK).JSON().Object()
	obj.Value("Brands").Array().NotEmpty()
}
