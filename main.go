package main

import (
	"net"
	"os"
	"proxy/config"
	"proxy/log"
	"strconv"
)

var (
	Proxy    *ProxyServer
	Logger   *log.ColorLogger
	Settings config.Config
)

// main - start of the program / load settings and start goroutines
func main() {
	Settings = config.GetConfig()
	Logger = createLogger(&Settings)
	startProxy(&Settings)
	startSysHooks()
	go startRegWatcher()
}

// startProxy - resolve config IP addresses and create a listener
func startProxy(conf *config.Config) {
	Logger.Info("Starting...\n")
	loc := conf.Client.Host + ":" + strconv.Itoa(conf.Client.Port)
	rem := conf.Target.Host + ":" + strconv.Itoa(conf.Target.Port)

	// resolve local host and port
	local, err := net.ResolveTCPAddr("tcp", loc)
	if err != nil {
		Logger.Error("failed to resolve local address %s - %s\n", loc, err)
		os.Exit(1)
	}
	Logger.Debug("Resolved local address\n")
	// resolve the target server
	remote, err := net.ResolveTCPAddr("tcp", rem)
	if err != nil {
		Logger.Error("Failed to connect to RotMG server: %s\n", err)
		os.Exit(1)
	}
	Logger.Debug("Resolved remote address\n")
	// try to create a local listener
	listener, err := net.ListenTCP("tcp", local)
	if err != nil {
		Logger.Error("Could not start the server: %s\n", conf.Client.Port, err)
		Logger.Error("If you have another RotMG proxy open, close it down and try again")
		os.Exit(1)
	}
	Logger.Info("Ready to accept connections!\n")
	// loop awaiting a client connection
	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			Logger.Error("Failed to accept connection from client: %s\n", err.Error())
			continue
		}
		Proxy = NewProxy(conn, local, remote, conf)
		go Proxy.Start()
	}
}

// createLogger - set up a logger with a given config (can be nil for default)
func createLogger(conf *config.Config) *log.ColorLogger {
	// defaults to false without checking nullptr
	debug := conf.Log.Debug
	trace := conf.Log.Trace

	return &log.ColorLogger{
		VeryVerbose: trace,
		Verbose:     debug,
	}
}

func startSysHooks() {

}

func startRegWatcher() {

}

//func startSysHooks() {
//	_, closed := syscall.GetProcPIDs()
//	if closed != nil {
//		fmt.Printf("Error getting processes: %s\n", closed.Error())
//		// todo: make a loop here on failure
//		return
//	}
//	killed, closed := syscall.KillCrashHandle()
//	if closed != nil {
//		fmt.Printf("Error killing crash handler: %s\n", closed.Error())
//	} else if !killed {
//		fmt.Printf("Failed to kill crash handler: %v\n", killed)
//	} else {
//		fmt.Println("Killed Unity crash handler successfully")
//	}
//}
