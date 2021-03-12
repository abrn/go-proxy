package main

import (
	"io"
	"net"
)

type ProxyServer struct {
	sentBytes     	uint64
	recdBytes 		uint64
	laddr, raddr  	*net.TCPAddr
	lconn, rconn  	io.ReadWriteCloser
	erred         	bool
	errsig        	chan bool

	Log      	  	Logger
	OutputHex     	bool
}

// NewProxy - create a new proxy interface with config options
func NewProxy(lconn *net.TCPConn, laddr, raddr *net.TCPAddr) *ProxyServer {
	// todo: implement config vars here
	logger := ColorLogger{
		VeryVerbose: true,
		Verbose:     true,
		Prefix:      "[proxy]",
		Color:       true,
	}
	logger.Info("Proxy server started")
	return &ProxyServer{
		lconn: lconn,
		laddr: laddr,
		raddr: raddr,
		erred: false,
		errsig: make(chan bool),
		// todo: implement config vars here
		Log: logger,
		OutputHex: false,
	}
}

// Start - called when a new connection is made to the proxy
func (p *ProxyServer) Start() {
	defer p.lconn.Close()

	var err error
	p.rconn, err = net.DialTCP("tcp", nil, p.raddr)
	if err != nil {
		p.Log.Error("Failed connecting to target host: %s", err)
		return
	}
	defer p.rconn.Close()

	p.Log.Info("New connection received - route: %s >>> %s", p.laddr.String(), p.raddr.String())
}