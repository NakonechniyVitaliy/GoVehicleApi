package models

type BodyStyle struct {
	ID    uint16 `json:"id" bson:"id"`
	Name  string `json:"name" bson:"name"`
	Value uint16 `json:"value" bson:"value"`
}
