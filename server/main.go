package main

import (
	"encoding/json"
	"fmt"
	"image/color"
	"log"
	"math"
	"net"
	"os"

	"github.com/efournival/iot-project/common"
	"github.com/efournival/iot-project/udp"
)

const (
	Port                   = 8122
	StandardTemperature    = 19
	StandardHumidity       = 50
	StandardLightIntensity = 50
)

var (
	// Sailors blue
	Low = color.RGBA{R: 33, G: 181, B: 223}
	// Sailors red
	High = color.RGBA{R: 243, G: 84, B: 84}
)

func main() {
	server, err := udp.NewServer(Port)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	server.OnReceive(func(addr *net.UDPAddr, data []byte) {
		sensorData := &common.SensorData{}
		if err := json.Unmarshal(data, sensorData); err != nil {
			log.Println(err)
			return
		}

		t := 0.5 + float64(sensorData.Temperature-StandardTemperature)/10
		h := 0.5 + float64(sensorData.Humidity-StandardHumidity)/100
		l := 0.5 - float64(sensorData.LightIntensity-StandardLightIntensity)/150
		p := math.Min(math.Max((t+h+l)/3.0, 0.0), 1.0)
		r := uint8((1.0-p)*float64(Low.R) + p*float64(High.R) + 0.5)
		g := uint8((1.0-p)*float64(Low.G) + p*float64(High.G) + 0.5)
		b := uint8((1.0-p)*float64(Low.B) + p*float64(High.B) + 0.5)

		jsonData, err := json.Marshal(&color.RGBA{R: r, G: g, B: b})
		if err != nil {
			log.Println(err)
			return
		}

		log.Printf("Received t=%d, h=%d and l=%d, computed p=%.2f and target color is #%x%x%x", sensorData.Temperature, sensorData.Humidity, sensorData.LightIntensity, p, r, g, b)
		server.Write(addr, jsonData)
	})

	log.Println("Now listening for sensor data on port", Port)
	server.Serve()
}
