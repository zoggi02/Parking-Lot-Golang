package slot

import "parkinglot/car"

// Slot represents a parking slot that can hold a car
type Slot struct {
	SlotNumber int
	Car        *car.Car
}

// CreateSlot creates a new slot instance
func CreateSlot(slotNumber int) *Slot {
	return &Slot{SlotNumber: slotNumber}
}

// AddSlotCar parks a car in this slot
func (s *Slot) AddSlotCar(c *car.Car) {
	s.Car = c
}

// RemoveCar removes the car from this slot
func (s *Slot) RemoveCar() {
	s.Car = nil
}

// IsAvailable checks if the slot is available
func (s *Slot) IsAvailable() bool {
	return s.Car == nil
}
