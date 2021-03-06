package pcf8591

import (
	"fmt"
	"os"

	"golang.org/x/exp/io/i2c"
)

type PCF8591 struct {
	dev *i2c.Device
}

func NewPCF8591(addr int) *PCF8591 {
	dev, err := i2c.Open(&i2c.Devfs{Dev: "/dev/i2c-1"}, addr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return &PCF8591{dev}
}

func (pcf *PCF8591) Read(pin byte) byte {
	buf := make([]byte, 1)

	pcf.Write([]byte{0x40 | (pin & 0x03)})

	for i := 0; i < 2; i++ {
		if err := pcf.dev.Read(buf); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	return buf[0]
}

func (pcf *PCF8591) Write(value []byte) {
	if err := pcf.dev.Write(value); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
