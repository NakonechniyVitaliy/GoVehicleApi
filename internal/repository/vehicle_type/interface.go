package brand

type Repository interface {
	NewVehicleType() error
	GetVehicleType() error
	DeleteVehicleType() error
}
