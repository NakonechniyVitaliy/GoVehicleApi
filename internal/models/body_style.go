package models

type BodyStyle struct {
	ID    uint16 `bson:"id"`
	Name  string `bson:"name"`
	Value uint16 `bson:"value"`
}
