package photoresistor

import (
	"github.com/efournival/iot-project/device/pcf8591"
)

type LightSensor struct {
	pcf *pcf8591.PCF8591
	pin byte
}

func NewLightSensor(pcf *pcf8591.PCF8591, pin byte) *LightSensor {
	sensor := &LightSensor{pcf, pin}
	sensor.GetLightIntensity()
	return sensor
}

func (sensor *LightSensor) GetLightIntensity() int {
	return int(sensor.pcf.Read(sensor.pin))
}
