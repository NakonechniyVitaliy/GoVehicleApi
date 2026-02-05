package models

type Brand struct {
	ID         uint16 `json:"id" bson:"id"`
	CategoryID uint16 `json:"category_id" bson:"category_id"`
	Count      uint16 `json:"cnt" bson:"cnt"`
	CountryID  uint16 `json:"country_id" bson:"country_id"`
	EngName    string `json:"eng" bson:"eng"`
	MarkaID    uint16 `json:"marka_id" bson:"marka_id"`
	Name       string `json:"name" bson:"name"`
	Slang      string `json:"slang" bson:"slang"`
	Value      uint16 `json:"value" bson:"value"`
}
