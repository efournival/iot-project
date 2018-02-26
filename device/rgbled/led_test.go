package rgbled

import (
	"fmt"
	"os"
	"testing"
	"time"

	"periph.io/x/periph"
	"periph.io/x/periph/host/bcm283x"
)

var led *RGBLED

func TestStart(t *testing.T) {
	if _, err := periph.Init(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	led = NewRGBLED(bcm283x.GPIO17, bcm283x.GPIO18, bcm283x.GPIO27)
}

func TestRed(t *testing.T) {
	led.SetColor(255, 0, 0)
	time.Sleep(time.Second)
}

func TestGreen(t *testing.T) {
	led.SetColor(0, 255, 0)
	time.Sleep(time.Second)
}

func TestBlue(t *testing.T) {
	led.SetColor(0, 0, 255)
	time.Sleep(time.Second)
}

func TestWhite(t *testing.T) {
	led.SetColor(255, 255, 255)
	time.Sleep(time.Second)
}

func TestOff(t *testing.T) {
	led.SetColor(0, 0, 0)
}
