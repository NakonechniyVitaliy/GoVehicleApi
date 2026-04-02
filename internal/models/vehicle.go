package models

type Vehicle struct {
	ID         uint16 `bson:"id" json:"id"`
	Brand      uint16 `bson:"brand" json:"brand"`
	DriverType uint16 `bson:"driver_type" json:"driver_type"`
	Gearbox    uint16 `bson:"gearbox" json:"gearbox"`
	BodyStyle  uint16 `bson:"body_style" json:"body_style"`
	Category   uint16 `bson:"category" json:"category"`
	Mileage    uint32 `bson:"mileage" json:"mileage"`
	Model      string `bson:"model" json:"model"`
	Price      uint32 `bson:"price" json:"price"`
}
