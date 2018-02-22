package photoresistor

import (
	"fmt"
	"testing"
	"time"

	"github.com/efournival/iot-project/device/pcf8591"
)

func TestGetLightIntensity(t *testing.T) {
	pcf := pcf8591.NewPCF8591(0x48)
	sensor := NewLightSensor(pcf, 1)

	for {
		fmt.Println("Light intensity =", sensor.GetLightIntensity())
		time.Sleep(time.Second)
	}
}
