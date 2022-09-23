package vehicle

import "fmt"

type Car struct {
	IsEngineEnable bool
}

func (pCar *Car) StartEngine() {
	pCar.IsEngineEnable = true
	fmt.Println("Car engine was started!")
}

func (pCar *Car) StopEngine() {
	pCar.IsEngineEnable = false
	fmt.Println("Car engine was stoped!")
}
