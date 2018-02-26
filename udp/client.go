package udp

type Client struct {
	*udp
}

func NewClient(serverIP string, port int) (*Client, error) {
	udp, err := newUdp(serverIP, port)
	if err != nil {
		return nil, err
	}

	go udp.loop()

	return &Client{udp}, nil
}

func (c *Client) Write(data []byte) error {
	return c.write(data)
}
