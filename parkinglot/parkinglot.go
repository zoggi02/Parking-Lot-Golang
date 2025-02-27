package parkinglot

import (
	"fmt"
	"parkinglot/car"
	"parkinglot/slot"
)

// ParkingLot represents a parking lot with a collection of slots
type ParkingLot struct {
	Capacity int
	Slots    []*slot.Slot
}

// CreateParkingLot creates a new parking lot with the given capacity
func CreateParkingLot(capacity int) *ParkingLot {
	parkingLot := &ParkingLot{
		Capacity: capacity,
		Slots:    make([]*slot.Slot, capacity),
	}

	// Initialize slots
	for i := 0; i < capacity; i++ {
		parkingLot.Slots[i] = slot.CreateSlot(i + 1) // Slot numbers start from 1
	}

	return parkingLot
}

// ParkCarToSlot adds a car to the nearest available slot
func (p *ParkingLot) ParkCarToSlot(ID string) string {
	for _, s := range p.Slots {
		if s.IsAvailable() {
			car := car.AddCar(ID)
			s.AddSlotCar(car)
			return fmt.Sprintf("Allocated slot number: %d", s.SlotNumber)
		}
	}
	return "Sorry, parking lot is full"
}

// Leave removes a car from the parking lot and calculates charges
func (p *ParkingLot) Leave(ID string, hours int) string {
	for _, s := range p.Slots {
		if s.Car != nil && s.Car.ID == ID {
			charge := 10
			if hours > 2 {
				charge += (hours - 2) * 10
			}
			s.RemoveCar()
			return fmt.Sprintf("Registration number %s with Slot Number %d is free with Charge $%d", ID, s.SlotNumber, charge)
		}
	}
	return fmt.Sprintf("Registration number %s not found", ID)
}

// Status prints the current status of the parking lot
func (p *ParkingLot) Status() string {
	var status string
	status += "Slot No. Registration No.\n"
	for _, s := range p.Slots {
		if s.Car != nil {
			status += fmt.Sprintf("%d %s\n", s.SlotNumber, s.Car.ID)
		}
	}
	return status
}
