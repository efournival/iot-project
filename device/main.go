package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/efournival/iot-project/device/humiture"
	"github.com/efournival/iot-project/device/pcf8591"
	"github.com/efournival/iot-project/device/photoresistor"
	led "github.com/efournival/iot-project/device/rgb-led"
	"github.com/efournival/iot-project/device/thermistor"
	"periph.io/x/periph"
	"periph.io/x/periph/host/bcm283x"
)

func main() {
	fmt.Println("Initializing sensors...")

	if _, err := periph.Init(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	pcf := pcf8591.NewPCF8591(0x48)
	fmt.Println("PCF8591 (analog-to-digital converter) \t OK")

	temp := thermistor.NewTemperatureSensor(pcf, 0)
	fmt.Println("Thermistor \t\t\t\t OK")

	light := photoresistor.NewLightSensor(pcf, 1)
	fmt.Println("Photoresistor \t\t\t\t OK")

	humtemp := humiture.NewHumitureSensor(bcm283x.GPIO22)
	lastHumidity := 0
	fmt.Println("DHT11 (humiture) \t\t\t OK")

	fmt.Println("\nInitializing RGB LED...")

	led := led.NewRGBLED(bcm283x.GPIO17, bcm283x.GPIO18, bcm283x.GPIO27)
	led.SetColor(0, 0, 0)
	fmt.Println("RGB LED through PWM pins \t\t OK")

	for {
		temperature := temp.GetTemperature()
		lightIntensity := light.GetLightIntensity()

		temperature2, humidity, err := humtemp.GetHumidityAndTemperature()
		if err == nil {
			temperature = (temperature + temperature2) / 2
			lastHumidity = humidity
		}

		fmt.Println()
		log.Println("Temperature =", temperature, "°C")
		log.Println("Light intensity =", lightIntensity)
		log.Println("Humidity =", lastHumidity, "%")

		time.Sleep(time.Second)
	}
}
