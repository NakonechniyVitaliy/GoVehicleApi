package requests

import (
	"fmt"
	"io"
	"net/http"
)

func GetBrands(key string) {

	url := "https://developers.ria.com/auto/new/marks?category_id=1&api_key=" + key

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching", url, err)
	} else {
		fmt.Println(resp.Body)
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	bodyString := string(bodyBytes)

	fmt.Println(bodyString)

}
