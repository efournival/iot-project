package udp

import "net"

type Server struct {
	*udp
}

func NewServer(port int) (*Server, error) {
	udp, err := newUdp("", port)
	if err != nil {
		return nil, err
	}

	return &Server{udp}, nil
}

func (s *Server) Serve() {
	s.loop()
}

func (s *Server) Write(to *net.UDPAddr, data []byte) error {
	return s.writeTo(to, data)
}
