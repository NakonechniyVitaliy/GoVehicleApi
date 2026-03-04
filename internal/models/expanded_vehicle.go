package models

type ExpandedVehicle struct {
	ID         uint16 `bson:"id" json:"id"`
	Brand      string `bson:"brand" json:"brand"`
	DriverType string `bson:"driver_type" json:"driver_type"`
	Gearbox    string `bson:"gearbox" json:"gearbox"`
	BodyStyle  string `bson:"body_style" json:"body_style"`
	Category   string `bson:"category" json:"category"`
	Mileage    uint32 `bson:"mileage" json:"mileage"`
	Model      string `bson:"model" json:"model"`
	Price      uint16 `bson:"price" json:"price"`
}
