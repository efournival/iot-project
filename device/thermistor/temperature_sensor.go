package thermistor

import (
	"math"

	"github.com/efournival/iot-project/device/pcf8591"
)

type TemperatureSensor struct {
	pcf *pcf8591.PCF8591
	pin byte
}

func NewTemperatureSensor(pcf *pcf8591.PCF8591, pin byte) *TemperatureSensor {
	sensor := &TemperatureSensor{pcf, pin}
	sensor.GetTemperature()
	return sensor
}

func (sensor *TemperatureSensor) GetTemperature() int {
	vr := 5 * float64(sensor.pcf.Read(sensor.pin)) / 255
	return int(1/(((math.Log(10000*vr/(5-vr)/10000))/3950)+(1/(273.15+25))) - 273.15)
}
