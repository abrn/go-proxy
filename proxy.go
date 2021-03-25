package main

import (
	"io"
	"net"
	"os"
	"os/signal"
	"proxy/config"
	"proxy/log"
	"proxy/network"
	"syscall"
	"time"
)

type ProxyServer struct {
	status        byte
	bytesSent     uint64
	bytesReceived uint64
	hostLocal     *net.TCPAddr
	hostRemote    *net.TCPAddr
	connLocal     io.ReadWriteCloser
	connRemote    io.ReadWriteCloser
	closed        bool
	errsig        chan bool
	showHex       bool
}

const (
	ProxyFailed   = iota
	ProxyStarting
	ProxyRunning
	ProxyStopped
)

// NewProxy - create a new proxy interface with config options
// (passing nil as the config uses the default settings)
func NewProxy(lconn *net.TCPConn, local, remote *net.TCPAddr, conf *config.Config) *ProxyServer {
	return &ProxyServer{
		status:     ProxyStarting,
		connLocal:  lconn,
		hostLocal:  local,
		hostRemote: remote,
		errsig:     make(chan bool),
		showHex:    conf.Log.UseHex,
	}
}

// Start - set up the proxy server, start a loop
// awaiting new connections and pipe all network
// data between the proxy and game server
func (p *ProxyServer) Start() {
	// defer closing the local connection so we don't run in to memory leaks
	defer p.connLocal.Close()
	// hook the interrupt signal so we can shut down gracefully
	intsig := make(chan os.Signal, 1)
	signal.Notify(intsig, os.Interrupt, syscall.SIGTERM)
	// resolve the connection to the target server
	var err error
	p.connRemote, err = net.DialTCP("tcp", nil, p.hostRemote)
	if err != nil {
		Logger.Error("Failed to connect to RotMG server: %s\n", err.Error())
		return
	}
	// defer closing the remote connection so we don't run in to memory leaks
	defer p.connRemote.Close()

	Logger.Info("Connection received - routing %s >>> %s\n", p.hostLocal.String(), p.hostRemote.String())
	// create an outgoing and incoming pipe coroutine
	go p.pipe(p.connLocal, p.connRemote)
	go p.pipe(p.connRemote, p.connLocal)
	// catch the sigterm/stop command
	go func() {
		if x := <-intsig; x != nil {
			Logger.Warn("Caught stop signal..\n")
			p.ShutDown()
		}
	}()
	// wait for connection closing or a fatal error
	<-p.errsig
	Logger.Info("Connection closed (%d bytes sent, %d bytes recieved)\n", p.bytesSent, p.bytesReceived)
}

// ShutDown - shut down gracefully and log a message
func (p *ProxyServer) ShutDown() {
	msg := "Shutting down at [%s] (%d bytes sent, %d bytes received)\n"
	tme := time.Now().Format("01/02/06 15:04:05")
	Logger.Info(msg, tme, p.bytesSent, p.bytesReceived)
	os.Exit(1)
}

// err - catch all errors on the proxyserver object
// an EOF error is thrown when a client disconnects
// (includes changing areas)
func (p *ProxyServer) err(s string, err error) {
	if p.status != ProxyRunning {
		return
	}
	// log every error except buffer EOFs
	if err != io.EOF {
		Logger.Warn(s, err)
	}
	// catch errors, disconnects, and area switches
	p.errsig <- true
	p.status = ProxyStopped
}

// pipe - run as a gouroutine to copy bytes into
// a buffer, process it and write to the target
func (p *ProxyServer) pipe(src, dst io.ReadWriter) {
	// get the direction of the bytes
	isOutgoing := src == p.connLocal
	var direction string
	if isOutgoing {
		direction = ">>> %d bytes sent%s\n"
		// hook client packets here
	} else {
		direction = "<<< %d bytes received%s\n"
		// hook server packets here
	}
	// set the format of the byte Output (hex / string)
	var byteFormat string
	if p.showHex {
		byteFormat = "%x\n"
	} else {
		byteFormat = "%s\n"
	}
	// create a temporary buffer and copy the bytes
	buffer := make([]byte, 0xffff)
	for {
		index, err := src.Read(buffer)
		if err != nil {
			p.err("Read error: %s\n", err)
			return
		}
		data := buffer[:index]
		head := buffer[0:4]
		packet := network.CreateGamePacket(head)
		if packet == nil {
			log.Logger.Warn("Got nil packet when reading header: %x\n", head)
			return
		}


		if index != 0 {
			// debug output the byte count and direction
			Logger.Debug(direction, index, "")
			// trace output the actual bytes
			Logger.Trace(byteFormat, data)
		}

		out, err := dst.Write(data)
		if err != nil {
			p.err("Write error: %s\n", err)
			return
		}
		if isOutgoing {
			p.bytesSent += uint64(out)
		} else {
			p.bytesReceived += uint64(out)
		}
	}
}

func (p *ProxyServer) process([]byte) {

}