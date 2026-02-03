package models

type DriverType struct {
	ID    int    `json:"id" bson:"id"`
	Name  string `json:"name" bson:"name"`
	Value int    `json:"value" bson:"value"`
}
