package main

type ElecConfig struct {
	Devices []Device
	Tarrif  float32
}

func (e ElecConfig) calculate() float32 {
	var totalCost float32
	for _, device := range e.Devices {
		deviceCost := device.calculateCost(e.Tarrif)
		totalCost = totalCost + deviceCost
	}
	return totalCost * 30.5
}

type Device struct {
	Name       string
	MaxWattage int
	AvgWattage int
	AvgOnTime  int
}

func (d Device) calculateCost(tarrif float32) float32 {
	kwh := (d.AvgWattage * d.AvgOnTime) / 1000
	cost := float32(kwh) * tarrif
	return cost / 100 // convert to pounds as opposed to pence
}
