package main

import (
	"encoding/json"
	"image/color"
	"log"
	"net"
	"os"
	"time"

	"github.com/efournival/iot-project/common"
	"github.com/efournival/iot-project/device/humiture"
	"github.com/efournival/iot-project/device/pcf8591"
	"github.com/efournival/iot-project/device/photoresistor"
	"github.com/efournival/iot-project/device/rgbled"
	"github.com/efournival/iot-project/device/thermistor"
	"github.com/efournival/iot-project/udp"
	"periph.io/x/periph"
	"periph.io/x/periph/host/bcm283x"
)

const (
	Port                 = 8122
	PCFAddress           = 0x48
	ThermistorPCFPin     = 0
	PhotoresistorPCFPin  = 1
	Undefined            = -1000
	ColorTransitionSteps = 20
)

var (
	DHT11Pin     *bcm283x.Pin = bcm283x.GPIO22
	LEDRPin      *bcm283x.Pin = bcm283x.GPIO17
	LEDGPin      *bcm283x.Pin = bcm283x.GPIO18
	LEDBPin      *bcm283x.Pin = bcm283x.GPIO27
	currentColor color.RGBA
)

func main() {
	if _, err := periph.Init(); err != nil {
		log.Println(err)
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
	led.SetColor(currentColor.R, currentColor.G, currentColor.B)
	log.Printf("RGB LED initialized on PWM pins %s, %s and %s", LEDRPin, LEDGPin, LEDBPin)

	serverIP := "127.0.0.1"
	if len(os.Args) > 1 {
		serverIP = os.Args[1]
	}

	client, err := udp.NewClient(serverIP, Port)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	client.OnReceive(func(addr *net.UDPAddr, data []byte) {
		newColor := &color.RGBA{}
		if err := json.Unmarshal(data, newColor); err != nil {
			log.Println(err)
			return
		}

		log.Printf("Received LED target color #%x%x%x", newColor.R, newColor.G, newColor.B)

		// Using code from https://stackoverflow.com/a/21835834
		go func() {
			for i := 0; i < ColorTransitionSteps; i++ {
				p := float64(i) / (ColorTransitionSteps - 1)
				r := uint8((1.0-p)*float64(currentColor.R) + p*float64(newColor.R) + 0.5)
				g := uint8((1.0-p)*float64(currentColor.G) + p*float64(newColor.G) + 0.5)
				b := uint8((1.0-p)*float64(currentColor.B) + p*float64(newColor.B) + 0.5)
				led.SetColor(r, g, b)
				time.Sleep(1000 / ColorTransitionSteps * time.Millisecond)
			}
		}()
	})

	log.Println("UDP client is now sending data to", serverIP, "on port", Port)

	for {
		time.Sleep(time.Second)

		temperature := temp.GetTemperature()
		lightIntensity := light.GetLightIntensity()

		if humidity, temperature2, err := humtemp.GetHumidityAndTemperature(); err == nil {
			lastTemperature = temperature2
			lastHumidity = humidity
		}

		if lastTemperature != Undefined {
			temperature = int((temperature + lastTemperature) / 2)
		}

		data := &common.SensorData{
			Temperature:    temperature,
			Humidity:       lastHumidity,
			LightIntensity: lightIntensity,
		}

		jsonData, err := json.Marshal(data)
		if err != nil {
			log.Println(err)
			continue
		}

		if err := client.Write(jsonData); err != nil {
			log.Println(err)
			continue
		}
	}
}
