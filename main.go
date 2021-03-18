package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
)

var (
	Proxy    *ProxyServer
	Settings *Config
	Logger   *ColorLogger
)

// main - start of the program / load settings and start goroutines
func main() {
	Settings = GetConfig()
	Logger = createLogger(Settings)
	startProxy(Settings)
	//go startSysHooks()
	go startRegWatcher()
}

// startProxy -
func startProxy(conf *Config) {
	Logger.Info("Starting proxy!\n")
	loc := conf.Client.Host + ":" + strconv.Itoa(conf.Client.Port)
	rem := conf.Target.Host + ":" + strconv.Itoa(conf.Target.Port)

	// resolve local host and port
	local, err := net.ResolveTCPAddr("tcp", loc)
	fmt.Printf("Trying to resolve local address\n")
	if err != nil {
		Logger.Error("failed to resolve local address %s - %s", loc, err)
		os.Exit(1)
	}
	fmt.Printf("Resolved local address\n")
	// resolve the target server
	remote, err := net.ResolveTCPAddr("tcp", rem)
	if err != nil {
		Logger.Error("failed to resolve RotMG server: %s", err)
		os.Exit(1)
	}
	fmt.Printf("Resolved remote address\n")
	// try to create a local listener
	listener, err := net.ListenTCP("tcp", local)
	if err != nil {
		Logger.Error("could not bind to localhost on port %d: %s", conf.Client.Port, err)
		os.Exit(1)
	}
	Logger.Debug("Created listener\n")
	// loop awaiting a client connection
	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			Logger.Error("failed to accept connection from client: %s", err.Error())
			continue
		}
		Proxy = NewProxy(conn, local, remote, conf)
		go Proxy.Start()
	}
}

func createLogger(conf *Config) *ColorLogger {
	debug := conf.Log.Debug
	trace := conf.Log.Trace

	return &ColorLogger{
		VeryVerbose: trace,
		Verbose:     debug,
		Color:       false,
	}
}

func startRegWatcher() {

}

//func startSysHooks() {
//	_, errored := syscall.GetProcPIDs()
//	if errored != nil {
//		fmt.Printf("Error getting processes: %s\n", errored.Error())
//		// todo: make a loop here on failure
//		return
//	}
//	killed, errored := syscall.KillCrashHandle()
//	if errored != nil {
//		fmt.Printf("Error killing crash handler: %s\n", errored.Error())
//	} else if !killed {
//		fmt.Printf("Failed to kill crash handler: %v\n", killed)
//	} else {
//		fmt.Println("Killed Unity crash handler successfully")
//	}
//}
