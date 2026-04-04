package brand

// BrandPayload is the request body for POST/PUT /brand
type BrandPayload struct {
	Brand struct {
		CategoryID *uint16 `json:"category_id"`
		Count      *uint16 `json:"cnt"`
		CountryID  *uint16 `json:"country_id"`
		EngName    *string `json:"eng"`
		MarkaID    *uint16 `json:"marka_id"`
		Name       *string `json:"name"`
		Slang      *string `json:"slang"`
		Value      *uint16 `json:"value"`
	} `json:"Brand"`
}
