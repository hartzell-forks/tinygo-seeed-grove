package main

import (
	"time"

	grove "github.com/soypat/seeed-grove"
)

// Define the hardware we are using.
var (
	shield grove.ShieldXiao
	rotary grove.SensorRotaryAngle
)

func main() {
	rotary.Configure(shield.Conn(0).ADC(), 1)
	buzzer := shield.Conn(1).PinOutput()

	for {
		// Full clockwise travel is ~300 degrees across the ADC range.
		angle := rotary.AngleDegrees()
		fraction := rotary.Fraction()
		halfPeriod := time.Duration(int(((2000-200)*fraction)+200) * int(time.Microsecond))
		println("fraction:", fraction, "angle:", angle, "halfPeriod", halfPeriod)
		beep(buzzer, halfPeriod, 200*time.Millisecond)
		// time.Sleep(200 * time.Millisecond)
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
