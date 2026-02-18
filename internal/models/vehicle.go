package models

type Vehicle struct {
	ID         uint16 `bson:"id"`
	Brand      uint16 `bson:"brand"`
	DriverType uint16 `bson:"driverType"`
	Gearbox    uint16 `bson:"gearbox"`
	BodyStyle  uint16 `bson:"bodyStyle"`
	Category   uint16 `bson:"category"`
	Mileage    uint32 `bson:"mileage"`
	Model      string `bson:"model"`
	Price      uint16 `bson:"price"`
}
