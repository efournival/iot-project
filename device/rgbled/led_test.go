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

func TestSmoothBlue(t *testing.T) {
	for i := 0; i < 10; i++ {
		p := float64(i) / 9
		r := uint8((1.0-p)*255.0 + 0.5)
		b := uint8(p*255.0 + 0.5)
		led.SetColor(r, 0, b)
		time.Sleep(1000 / 10 * time.Millisecond)
	}
}

func TestOff(t *testing.T) {
	led.SetColor(0, 0, 0)
}
