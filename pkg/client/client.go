package client

import (
	"fmt"
	"net"
)

const DEFAULT_PORT = "2021"

func New(slicerIP string) (net.Conn, error) {
	targetIP := fmt.Sprintf("%s:%s", slicerIP, DEFAULT_PORT)

	return net.Dial("udp", targetIP)
}
