package humiture

import (
	"encoding/hex"
	"errors"
	"fmt"
	"os"
	"time"

	"periph.io/x/periph"
	"periph.io/x/periph/conn/gpio"
)

type HumitureSensor struct {
	pin gpio.PinIO
}

func init() {
	if _, err := periph.Init(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func NewHumitureSensor(pin gpio.PinIO) *HumitureSensor {
	return &HumitureSensor{pin}
}

func (sensor *HumitureSensor) checkLevel(level gpio.Level) error {
	i := 20000

	for i > 0 && sensor.pin.Read() == level {
		i--
	}

	if i > 0 {
		return nil
	}

	return errors.New("DHT11 read timeout")
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// Using code found here: https://github.com/adalton/arduino/blob/master/projects/Dht11_Library/Dht11.cpp
func (sensor *HumitureSensor) GetHumidityAndTemperature() (int, int, error) {
	checkError(sensor.pin.Out(gpio.Low))
	time.Sleep(20 * time.Millisecond)
	checkError(sensor.pin.Out(gpio.High))
	time.Sleep(5 * time.Microsecond)
	checkError(sensor.pin.In(gpio.PullNoChange, gpio.NoEdge))

	if err := sensor.checkLevel(gpio.Low); err != nil {
		return 0, 0, err
	}

	if err := sensor.checkLevel(gpio.High); err != nil {
		return 0, 0, err
	}

	data := make([]byte, 5)
	var index uint = 7

	for i := 0; i < len(data)*8; i++ {
		if err := sensor.checkLevel(gpio.Low); err != nil {
			return 0, 0, err
		}

		t := time.Now()

		if err := sensor.checkLevel(gpio.High); err != nil {
			return 0, 0, err
		}

		if time.Now().Sub(t).Nanoseconds()/1000 > 40 {
			data[i/8] |= (1 << index)
		}

		if index > 0 {
			index--
		} else {
			index = 7
		}
	}

	if data[0]+data[2] == data[4] {
		return int(data[0]), int(data[2]), nil
	}

	return 0, 0, errors.New("DHT11 data corrupted, got " + hex.EncodeToString(data))
}
