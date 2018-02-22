package photoresistor

import (
	"fmt"
	"testing"
	"time"

	"github.com/efournival/iot-project/device/pcf8591"
)

func TestGetLightIntensity(t *testing.T) {
	pcf := pcf8591.NewPCF8591(0x48, 120)
	sensor := NewLightSensor(pcf, 1)

	for i := 0; i < 10; i++ {
		fmt.Println("Light intensity =", sensor.GetLightIntensity())
		time.Sleep(500 * time.Millisecond)
	}
}
