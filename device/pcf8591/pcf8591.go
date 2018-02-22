package pcf8591

import (
	"fmt"
	"os"

	"golang.org/x/exp/io/i2c"
)

type PCF8591 struct {
	dev     *i2c.Device
	pinBase byte
}

func NewPCF8591(addr int, pinBase byte) *PCF8591 {
	dev, err := i2c.Open(&i2c.Devfs{Dev: "/dev/i2c-1"}, addr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return &PCF8591{dev, pinBase}
}

func (pcf *PCF8591) Read(pin byte) byte {
	buf := make([]byte, 1)

	if err := pcf.dev.ReadReg(0x40|((pin-pcf.pinBase)&3), buf); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return buf[0]
}

func (pcf *PCF8591) Write(value byte) {
	if err := pcf.dev.WriteReg(0x40, []byte{value}); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
