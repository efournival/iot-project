package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/efournival/iot-project/device/humiture"
	"github.com/efournival/iot-project/device/pcf8591"
	"github.com/efournival/iot-project/device/photoresistor"
	"github.com/efournival/iot-project/device/rgbled"
	"github.com/efournival/iot-project/device/thermistor"
	"periph.io/x/periph"
	"periph.io/x/periph/host/bcm283x"
)

const (
	PCFAddress          = 0x48
	ThermistorPCFPin    = 0
	PhotoresistorPCFPin = 1
	Undefined           = -1000
)

var (
	DHT11Pin *bcm283x.Pin = bcm283x.GPIO22
	LEDRPin  *bcm283x.Pin = bcm283x.GPIO17
	LEDGPin  *bcm283x.Pin = bcm283x.GPIO18
	LEDBPin  *bcm283x.Pin = bcm283x.GPIO27
)

func main() {
	if _, err := periph.Init(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	pcf := pcf8591.NewPCF8591(PCFAddress)
	log.Println("PCF8591 initialized")

	temp := thermistor.NewTemperatureSensor(pcf, ThermistorPCFPin)
	log.Println("Thermistor initialized on PCF input pin", ThermistorPCFPin)

	light := photoresistor.NewLightSensor(pcf, PhotoresistorPCFPin)
	log.Println("Photoresistor initialized on PCF input pin", PhotoresistorPCFPin)

	humtemp := humiture.NewHumitureSensor(DHT11Pin)
	lastHumidity := 50
	lastTemperature := Undefined
	log.Println("DHT11 (humiture sensor) initialized on pin", DHT11Pin)

	led := rgbled.NewRGBLED(LEDRPin, LEDGPin, LEDBPin)
	led.SetColor(0, 0, 0)
	log.Printf("RGB LED initialized on PWM pins %s, %s and %s", LEDRPin, LEDGPin, LEDBPin)

	for {
		temperature := temp.GetTemperature()
		lightIntensity := light.GetLightIntensity()

		if humidity, temperature2, err := humtemp.GetHumidityAndTemperature(); err == nil {
			lastTemperature = temperature2
			lastHumidity = humidity
		}

		if lastTemperature != Undefined {
			temperature = int((temperature + lastTemperature) / 2)
		}

		fmt.Println()
		log.Println("Temperature =", temperature, "°C")
		log.Println("Light intensity =", lightIntensity)
		log.Println("Humidity =", lastHumidity, "%")

		time.Sleep(time.Second)
	}
}
