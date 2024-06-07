package main

import (
	"fmt"
	"runtime"
	"os"

	"github.com/shirou/gopsutil/host"
	
)

func main() {
	
	osName := runtime.GOOS


	fmt.Println("Operating System:", osName)

	
	hostInfo, _ := host.Info()
	fmt.Println("Host name:", hostInfo.Hostname)
	fmt.Println("Platform:", hostInfo.Platform)
	fmt.Println("OS version:", hostInfo.OS)
	fmt.Println("Kernel version:", hostInfo.KernelVersion)

	
	cpuInfo, _ := host.Info()
	fmt.Println("CPU model:", cpuInfo.Hostname)
	fmt.Println("Number of CPUs:", cpuInfo.Platform)

	
	pid := os.Getpid()

	parentpid := os.Getppid()

	fmt.Printf("The parent process id of %v is %v\n", pid, parentpid)
}
