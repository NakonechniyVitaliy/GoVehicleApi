package cache_key

import "fmt"

func VehicleByID(id uint16) string {
	return fmt.Sprintf("vehicle:%d", id)
}
