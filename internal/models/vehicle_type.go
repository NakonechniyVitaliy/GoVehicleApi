package models

type VehicleType struct {
	ID         int    `json:"id" bson:"id"`
	Ablative   string `json:"ablative" bson:"ablative"`
	CategoryID int    `json:"category_id" bson:"category_id"`
	Name       string `json:"name" bson:"name"`
	Plural     string `json:"plural" bson:"plural"`
	Rewrite    string `json:"rewrite" bson:"rewrite"`
	Singular   string `json:"singular" bson:"singular"`
}
