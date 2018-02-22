package humiture

import (
	"fmt"
	"testing"
	"time"

	"periph.io/x/periph/host/bcm283x"
)

func TestGetHumidityAndTemperature(t *testing.T) {
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
