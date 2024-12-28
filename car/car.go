package car

// Car represents a car with its registration number
type Car struct {
	ID string
}

// AddCar creates a new car instance
func AddCar(ID string) *Car {
	return &Car{ID: ID}
}
