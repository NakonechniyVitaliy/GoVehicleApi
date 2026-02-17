package models

type Vehicle struct {
	ID         uint16 `json:"id" bson:"id"`
	Brand      uint16 `json:"brand" bson:"brand"`
	DriverType uint16 `json:"driverType" bson:"driverType"`
	Gearbox    uint16 `json:"gearbox" bson:"gearbox"`
	BodyStyle  uint16 `json:"bodyStyle" bson:"bodyStyle"`
	Category   uint16 `json:"category" bson:"category"`
	Mileage    uint32 `json:"mileage" bson:"mileage"`
	Model      string `json:"model" bson:"model"`
	Price      uint16 `json:"price" bson:"price"`
}
