package game

import (
	"bufio"
	"net"
)

type Connection struct {
	Connected bool
	Destroyed bool
	Handler   net.Conn
	Socket    *bufio.ReadWriter
}
