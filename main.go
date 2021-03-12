package main

import (
	"fmt"
	"proxy/syscall"
)
var (
	Proxy 		ProxyServer
	Settings 	Config
)

func main() {
	Settings = GetConfig()
	startProxy()
	startSysHooks()
	startRegWatcher()
}

func startProxy() {

}

func startRegWatcher() {
	// testing
	scrWidth, err := syscall.GetWinHeight()
	if err != nil {
		fmt.Printf("Failed to find screen height value: %s\n", err.Error())
	} else {
		fmt.Printf("Found Exalt screen height value: %d\n", scrWidth)
	}
	scrHeight, err := syscall.GetWinWidth()
	if err != nil {
		fmt.Printf("Failed to find screen width value: %s\n", err.Error())
	} else {
		fmt.Printf("Found Exalt screen width value: %d\n", scrHeight)
	}
	guid, err := syscall.GetExaltGUID()
	if err != nil {
		fmt.Printf("Failed to find screen width value: %s\n", err.Error())
	} else {
		fmt.Printf("Found Exalt GUID: %s\n", guid)
	}
	serv, err := syscall.GetLastServer()
	if err != nil {
		fmt.Printf("Failed to find last server: %s\n", err.Error())
	} else {
		fmt.Printf("Found Exalt last server: %s\n\n", serv)
	}
}

func startSysHooks() {
	_, err := syscall.GetProcPIDs()
	if err != nil {
		fmt.Printf("Error getting processes: %s\n", err.Error())
		// todo: make a loop here on failure
		return
	}
	killed, err := syscall.KillCrashHandle()
	if err != nil {
		fmt.Printf("Error killing crash handler: %s\n", err.Error())
	} else if !killed {
		fmt.Printf("Failed to kill crash handler: %v\n", killed)
	} else {
		fmt.Println("Killed Unity crash handler successfully")
	}
}