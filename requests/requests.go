package requests

import (
	"encoding/json"
	"fmt"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/models"
	"io"
	"net/http"
)

func GetBrands(key string) ([]models.Brand, error) {

	url := "https://developers.ria.com/auto/new/marks?category_id=1&api_key=" + key

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching", url, err)
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
	}

	var brands []models.Brand

	err = json.Unmarshal(bodyBytes, &brands)
	if err != nil {
		return nil, fmt.Errorf("Error unmarshalling brands: %s", string(bodyBytes))
	}

	return brands, nil

}
