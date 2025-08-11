package models

type Brand struct {
	Category int    `json:"category_id"`
	Count    int    `json:"cnt"`
	Country  int    `json:"country_id"`
	EngName  string `json:"eng"`
	Marka    int    `json:"marka_id"`
	Name     string `json:"name"`
	Slang    string `json:"slang"`
	Value    int    `json:"value"`
}
