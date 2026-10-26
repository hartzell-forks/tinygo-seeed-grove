package main

import (
	"time"

	grove "github.com/soypat/seeed-grove"
)

const BuzzerConnPosition = 1
const TouchConnPosition = 0

// Define the hardware we are using.
var shield grove.ShieldXiao

func main() {
	touch := shield.Conn(TouchConnPosition).PinInputPulldown()
	buzzer := shield.Conn(BuzzerConnPosition).PinOutput()

	// Toggle the pin to make a square wave; half-period sets the pitch.
	const halfPeriod = 500 * time.Microsecond // ~1 kHz

	for {
		if touch.GetLevel() {
			beep(buzzer, halfPeriod, 100*time.Millisecond)
		}
	}
}

func beep(pin grove.PinOutput, halfPeriod, dur time.Duration) {
	for elapsed := time.Duration(0); elapsed < dur; elapsed += 2 * halfPeriod {
		pin.High()
		time.Sleep(halfPeriod)
		pin.Low()
		time.Sleep(halfPeriod)
	}
}
