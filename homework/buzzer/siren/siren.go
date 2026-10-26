package main

import (
	"time"

	grove "github.com/soypat/seeed-grove"
)

const ConnPosition = 1

// Define the hardware we are using.
var shield grove.ShieldXiao

func main() {
	buzzer := shield.Conn(ConnPosition).PinOutput()

	halfPeriod := 300
	dir := 1
	for {
		if halfPeriod < 300 || halfPeriod > 1000 {
			dir = dir * -1
		}
		beep(buzzer, time.Duration(halfPeriod)*time.Microsecond, 5*time.Millisecond)
		halfPeriod += 5 * dir
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
