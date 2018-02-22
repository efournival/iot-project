package led

import (
	"fmt"
	"os"

	"periph.io/x/periph/conn/gpio"
)

type RGBLED struct {
	r, g, b gpio.PinPWM
}

func NewRGBLED(rpin, gpin, bpin gpio.PinIO) *RGBLED {
	return &RGBLED{
		rpin.(gpio.PinPWM),
		gpin.(gpio.PinPWM),
		bpin.(gpio.PinPWM),
	}
}

func pwm(pin gpio.PinPWM, value uint8) {
	if err := pin.PWM(gpio.DutyMax-gpio.Duty(float32(value/255)*float32(gpio.DutyMax)), 0); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func (led *RGBLED) SetColor(r, g, b uint8) {
	pwm(led.r, r)
	pwm(led.g, g)
	pwm(led.b, b)
}
