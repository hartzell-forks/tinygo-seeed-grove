package main

import (
	"time"

	grove "github.com/soypat/seeed-grove"
)

const ConnPosition = 2

// Define the hardware we are using.
var (
	shield grove.ShieldXiao
	light  grove.SensorLight
)

func main() {
	light.Configure(shield.Conn(ConnPosition).ADC(), 1)
	for {
		fraction := light.Fraction()
		if fraction < 0.5 {
			println("Dark... ", fraction)
		} else {
			println("Light! ", fraction)
		}
		time.Sleep(200 * time.Millisecond)
	}
}
