package main

import (
	"io"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type ProxyServer struct {
	status        ProxyStatus
	bytesSent     uint64
	bytesReceived uint64
	hostLocal     *net.TCPAddr
	hostRemote    *net.TCPAddr
	connLocal     io.ReadWriteCloser
	connRemote    io.ReadWriteCloser
	errored       bool
	errChan       chan bool
	showHex       bool
}

type ProxyStatus byte

const (
	Failed   ProxyStatus = 0
	Starting ProxyStatus = 1
)

// NewProxy - create a new proxy interface with config options
func NewProxy(lconn *net.TCPConn, local, remote *net.TCPAddr, conf *Config) *ProxyServer {
	Logger.Info("Proxy server started!")
	return &ProxyServer{
		connLocal:  lconn,
		hostLocal:  local,
		hostRemote: remote,
		errChan:    make(chan bool),
		showHex:    true,
	}
}

// Start - called when a new connection is made to the proxy
func (p *ProxyServer) Start() {
	defer p.connLocal.Close()
	// catch the cmd interrupt and shut down gracefully
	intsig := make(chan os.Signal, 1)
	signal.Notify(intsig, os.Interrupt, syscall.SIGTERM)
	// resolve the connection to the target server
	var err error
	p.connRemote, err = net.DialTCP("tcp", nil, p.hostRemote)
	if err != nil {
		Logger.Error("Failed connecting to RotMG server: %s", err.Error())
		return
	}
	defer p.connRemote.Close()

	Logger.Info("New connection received - routing %s >>> %s", p.hostLocal.String(), p.hostRemote.String())
	// create an outgoing and incoming pipe coroutine
	go p.pipe(p.connLocal, p.connRemote)
	go p.pipe(p.connRemote, p.connLocal)
	// wait for a fatal error in any routines and shutdown
	go func() {
		<-intsig
		Logger.Info("Caught stop signal..")
		p.errChan <- true
	}()
	<-p.errChan
	p.ShutDown()
}

// ShutDown - gracefully shut down
func (p *ProxyServer) ShutDown() {
	msg := "Shutting down at [%s] (%d bytes sent, %d bytes recieved)"
	tme := time.Now().Format("01/02/06 15:04:05")
	Logger.Info(msg, tme, p.bytesSent, p.bytesReceived)
	os.Exit(1)
}

// err - catch all errors on the proxyserver object
func (p *ProxyServer) err(s string, err error) {
	if p.errored {
		return
	}
	//
	if err != io.EOF {
		Logger.Warn(s, err.Error())
	}
	p.errChan <- true
	p.errored = true
}

// pipe - copy bytes from an incoming/outgoing packet into a rw buffer
func (p *ProxyServer) pipe(src, dst io.ReadWriter) {
	// get the direction of the bytes
	isOutgoing := src == p.connLocal
	var direction string
	if isOutgoing {
		direction = ">>> %d bytes sent%s"
		// hook client packets here
	} else {
		direction = "<<< %d bytes received%s"
		// hook server packets here
	}
	// set the format of the byte Output (hex / string)
	var byteFormat string
	if p.showHex {
		byteFormat = "%x"
	} else {
		byteFormat = "%s"
	}
	// create a 64k temporary buffer and copy the bytes
	buffer := make([]byte, 0xffff)
	for {
		index, err := src.Read(buffer)
		if err != nil {
			// should only throw when the buffer is too small
			p.err("Error reading pipe: %s\n", err)
		}
		temp := buffer[:index]
		// debug output the byte count and direction
		Logger.Debug(direction, index, "")
		// trace output the actual bytes
		Logger.Trace(byteFormat, temp)

		out, err := dst.Write(temp)
		if err != nil {
			p.err("Error writing the pipe: %s\n", err)
		}
		if isOutgoing {
			p.bytesSent += uint64(out)
		} else {
			p.bytesReceived += uint64(out)
		}
	}
}

func (p *ProxyServer) awaitConnection() {

}
