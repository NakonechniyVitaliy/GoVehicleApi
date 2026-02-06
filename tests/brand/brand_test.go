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

	randomBrandName := gofakeit.CarModel()
	randomBrandID := gofakeit.Uint16()

	testCases := []testCase{
		{
			CaseName:   "Valid brand",
			CategoryID: gofakeit.Uint16(),
			Count:      gofakeit.Uint16(),
			CountryID:  gofakeit.Uint16(),
			EngName:    randomBrandName,
			MarkaID:    randomBrandID,
			BrandName:  randomBrandName,
			Slang:      randomBrandName,
			Value:      randomBrandID,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.CaseName, func(t *testing.T) {
			tcUrl := url.URL{
				Scheme: "http",
				Host:   tests.LOCAL_HOST,
			}
			e := httpexpect.Default(t, tcUrl.String())

			doTestSave(e, tc)

		})
	}

}

func doTestSave(e *httpexpect.Expect, tc testCase) {
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

	fmt.Println(resp)

}
