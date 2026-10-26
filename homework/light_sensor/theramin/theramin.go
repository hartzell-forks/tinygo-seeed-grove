package main

import (
	"time"

	grove "github.com/soypat/seeed-grove"
)

const BuzzerConnPosition = 1
const LightSensorConnPosition = 2

// Define the hardware we are using.
var (
	shield grove.ShieldXiao
	light  grove.SensorLight
)

func main() {
	light.Configure(shield.Conn(LightSensorConnPosition).ADC(), 1)
	buzzer := shield.Conn(BuzzerConnPosition).PinOutput()

	var halfPeriod time.Duration

	for {
		fraction := light.Fraction()
		halfPeriod = time.Duration(int(((2000-200)*fraction)+200) * int(time.Microsecond))
		beep(buzzer, halfPeriod, 200*time.Millisecond)
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
