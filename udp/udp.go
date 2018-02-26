package udp

import (
	"fmt"
	"net"
	"time"
)

type ReceiveFunc = func(*net.UDPAddr, []byte)

type udp struct {
	conn      *net.UDPConn
	onReceive ReceiveFunc
}

func newUdp(address string, port int) (*udp, error) {
	var conn *net.UDPConn

	if len(address) == 0 {
		addr, err := net.ResolveUDPAddr("udp", fmt.Sprintf(":%d", port))
		if err != nil {
			return nil, err
		}

		conn, err = net.ListenUDP("udp", addr)
		if err != nil {
			return nil, err
		}
	} else {
		addr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%d", address, port))
		if err != nil {
			return nil, err
		}

		conn, err = net.DialUDP("udp", nil, addr)
		if err != nil {
			return nil, err
		}
	}

	return &udp{conn, func(_ *net.UDPAddr, _ []byte) {}}, nil
}

func (u *udp) loop() {
	for {
		var buffer [1024]byte

		u.conn.SetReadDeadline(time.Now().Add(500 * time.Millisecond))
		n, client, err := u.conn.ReadFromUDP(buffer[0:])
		if err != nil {
			continue
		}

		u.onReceive(client, buffer[0:n])
	}
}

func (u *udp) OnReceive(rf ReceiveFunc) {
	u.onReceive = rf
}

func (u *udp) writeTo(to *net.UDPAddr, data []byte) error {
	u.conn.SetWriteDeadline(time.Now().Add(500 * time.Millisecond))
	_, err := u.conn.WriteToUDP(data, to)
	return err
}

func (u *udp) write(data []byte) error {
	u.conn.SetWriteDeadline(time.Now().Add(500 * time.Millisecond))
	_, err := u.conn.Write(data)
	return err
}
