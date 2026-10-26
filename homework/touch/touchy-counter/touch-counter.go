package main

import (
	"time"

	grove "github.com/soypat/seeed-grove"
)

// Define the hardware we are using.
var shield grove.ShieldXiao

func main() {
	// Capacitive touchSensor drives the data line high while a finger is present.
	touchSensor := shield.Conn(0).PinInputPulldown()
	last := false
	touches := 0
	for {
		beingTouched := touchSensor.GetLevel()
		if last != beingTouched {
			touches += 1
			println("Touches: ", touches)
		}
		last = beingTouched
		time.Sleep(20 * time.Millisecond)
	}
}
