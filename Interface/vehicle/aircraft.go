package vehicle

import "fmt"

type Aircraft struct {
	IsEngineEnable bool
}

func (pAircraft *Aircraft) StartEngine() {
	pAircraft.IsEngineEnable = true
	fmt.Println("Aircraft engine was started!")
}

func (pAircraft *Aircraft) StopEngine() {
	pAircraft.IsEngineEnable = false
	fmt.Println("Aircraft engine was stoped!")
}
