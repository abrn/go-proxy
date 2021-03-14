package network

import (
	"bufio"
	"net"
)

type GameConnection struct {
	Connected bool
	Destroyed bool
	Handler   net.Conn
	Socket    *bufio.ReadWriter
}
