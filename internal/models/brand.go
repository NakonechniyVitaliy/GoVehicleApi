package models

type Brand struct {
	ID       int    `json:"id" bson:"id"`
	Category int    `json:"category_id" bson:"category_id"`
	Count    int    `json:"cnt" bson:"cnt"`
	Country  int    `json:"country_id" bson:"country_id"`
	EngName  string `json:"eng" bson:"eng"`
	MarkaID  int    `json:"marka_id" bson:"marka_id"`
	Name     string `json:"name" bson:"name"`
	Slang    string `json:"slang" bson:"slang"`
	Value    int    `json:"value" bson:"value"`
}
