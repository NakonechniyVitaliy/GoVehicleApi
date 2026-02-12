package brand

type BrandDTO struct {
	CategoryID *uint16 `json:"category_id"`
	Count      *uint16 `json:"cnt"`
	CountryID  *uint16 `json:"country_id" `
	EngName    *string `json:"eng"`
	MarkaID    *uint16 `json:"marka_id"`
	Name       *string `json:"name" `
	Slang      *string `json:"slang" `
	Value      *uint16 `json:"value"`
}
