package main

import (
	"time"

	grove "github.com/soypat/seeed-grove"
	"github.com/soypat/seeed-grove/homework/rotary_angle_sensor/selector/binify"
)

// Define the hardware we are using.
var (
	shield grove.ShieldXiao
	rotary grove.SensorRotaryAngle
)

func main() {
	rotary.Configure(shield.Conn(0).ADC(), 1)

	current_bin := 1

	for {
		percent := rotary.Fraction() * 100
		bin, _ := binify.Binify(percent)
		if bin != current_bin {
			println("Bucket: ", bin)
			current_bin = bin
		}
		time.Sleep(200 * time.Millisecond)
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
