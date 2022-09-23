package main

import (
	"fmt"
	"main/vehicle"
)

type ivehicle interface {
	StartEngine()
	StopEngine()
}

type ivehicleStart interface {
	StartEngine()
}

type ivehicleStop interface {
	StopEngine()
}

func main() {
	// ivehicle.
	pVehicle1 := new(vehicle.Aircraft)
	pVehicle2 := new(vehicle.Car)

	fmt.Printf("(%T, %+8v)\n", pVehicle1, pVehicle1)
	fmt.Printf("(%T, %+8v)\n", pVehicle2, pVehicle2)

	sliceVeh := []ivehicle{pVehicle1, pVehicle2}

	for _, pVeh := range sliceVeh {
		pVeh.StartEngine()
		pVeh.StopEngine()
	}

	// ivehicleStart and ivehicleStop.
	var vehicleStart ivehicleStart = new(vehicle.Aircraft)
	vehicleStart.StartEngine()
	var vehicleStop ivehicleStop = new(vehicle.Aircraft)
	vehicleStop.StopEngine()
}
