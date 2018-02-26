package humiture

import (
	"fmt"
	"os"
	"testing"
	"time"

	"periph.io/x/periph"
	"periph.io/x/periph/host/bcm283x"
)

func TestGetHumidityAndTemperature(t *testing.T) {
	if _, err := periph.Init(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	sensor := NewHumitureSensor(bcm283x.GPIO22)

	for {
		humidity, temperature, err := sensor.GetHumidityAndTemperature()

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Humidity =", humidity, "% \t Temperature =", temperature, "°C")
		}

		time.Sleep(time.Second)
	}
}
