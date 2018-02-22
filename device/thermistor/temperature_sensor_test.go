package thermistor

import (
	"fmt"
	"testing"
	"time"

	"github.com/efournival/iot-project/device/pcf8591"
)

func TestGetTemperature(t *testing.T) {
	pcf := pcf8591.NewPCF8591(0x48)
	sensor := NewTemperatureSensor(pcf, 0)

	for {
		fmt.Println("Temperature =", sensor.GetTemperature(), "°C")
		time.Sleep(time.Second)
	}
}
