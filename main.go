package main

import (
	"net"
	"os"
	"strconv"
)

var (
	Proxy    *ProxyServer
	Logger   *ColorLogger
	Settings Config
)

// main - start of the program / load settings and start goroutines
func main() {
	Settings = GetConfig()
	Logger = createLogger(&Settings)
	startProxy(&Settings)
	//go startSysHooks()
	go startRegWatcher()
}

// startProxy -
func startProxy(conf *Config) {
	Logger.Info("Starting...\n")
	loc := conf.Client.Host + ":" + strconv.Itoa(conf.Client.Port)
	rem := conf.Target.Host + ":" + strconv.Itoa(conf.Target.Port)

	// resolve local host and port
	local, err := net.ResolveTCPAddr("tcp", loc)
	if err != nil {
		Logger.Error("failed to resolve local address %s - %s\n", loc, err)
		os.Exit(1)
	}
	Logger.Trace("Resolved local address\n")
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
		Logger.Error("Could not start the server (port %d in use): %s\n", conf.Client.Port, err)
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

func createLogger(conf *Config) *ColorLogger {
	// defaults to false without checking nullptr
	debug := conf.Log.Debug
	trace := conf.Log.Trace

	return &ColorLogger{
		VeryVerbose: trace,
		Verbose:     debug,
	}
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
