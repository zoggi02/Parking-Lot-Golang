package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"parkinglot/parkinglot" // Update dengan nama modul
	"parkinglot/utils"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Error			: Invalid arguments")
		fmt.Println("example usage  : go run main.go <input_file>")
		return
	}

	// Read File
	file, err := utils.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	var parkingLot *parkinglot.ParkingLot
	scanner := bufio.NewScanner(file)

	// Loop command from file
	for scanner.Scan() {
		line := scanner.Text()
		commands := strings.Split(line, " ")
		cmd := commands[0]

		// Execute command
		switch cmd {
		case "create_parking_lot":
			if len(commands) != 2 {
				fmt.Println("Invalid command format")
				continue
			}
			capacity, err := strconv.Atoi(commands[1])
			if err != nil {
				fmt.Println("Invalid capacity")
				continue
			}
			parkingLot = parkinglot.CreateParkingLot(capacity)
			fmt.Printf("Created parking lot with %d slots\n", capacity)

		case "park":
			if parkingLot == nil {
				fmt.Println("Parking lot not created")
				continue
			}
			if len(commands) != 2 {
				fmt.Println("Invalid command format")
				continue
			}
			fmt.Println(parkingLot.ParkCarToSlot(commands[1]))

		case "leave":
			if parkingLot == nil {
				fmt.Println("Parking lot not created")
				continue
			}
			if len(commands) != 3 {
				fmt.Println("Invalid command format")
				continue
			}
			registration := commands[1]
			hours, err := strconv.Atoi(commands[2])
			if err != nil {
				fmt.Println("Invalid hours")
				continue
			}
			fmt.Println(parkingLot.Leave(registration, hours))

		case "status":
			if parkingLot == nil {
				fmt.Println("Parking lot not created")
				continue
			}
			fmt.Print(parkingLot.Status())

		default:
			fmt.Println("Unknown command : ", cmd)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %s\n", err)
	}
}
