package models

type VehicleType struct {
	ID         uint16 `json:"id" bson:"id"`
	Ablative   string `json:"ablative" bson:"ablative"`
	CategoryID uint16 `json:"category_id" bson:"category_id"`
	Name       string `json:"name" bson:"name"`
	Plural     string `json:"plural" bson:"plural"`
	Rewrite    string `json:"rewrite" bson:"rewrite"`
	Singular   string `json:"singular" bson:"singular"`
}
