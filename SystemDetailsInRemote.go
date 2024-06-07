package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"

	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/process"
)

func main() {
	// Remote server details
	remoteServer := "T50523@INCHN0117DC"

	// Get host information
	hostInfo, err := host.InfoWithContext(nil, remoteServer)
	if err != nil {
		log.Fatalf("Failed to get host information: %v", err)
	}

	fmt.Println("Host name:", hostInfo.Hostname)
	fmt.Println("Platform:", hostInfo.Platform)
	fmt.Println("OS version:", hostInfo.OS)
	fmt.Println("Kernel version:", hostInfo.KernelVersion)

	// Get CPU information
	cpuInfo, err := host.InfoWithContext(nil, remoteServer)
	if err != nil {
		log.Fatalf("Failed to get CPU information: %v", err)
	}

	fmt.Println("CPU model:", cpuInfo.Hostname)
	fmt.Println("Number of CPUs:", cpuInfo.Platform)

	// Get memory usage
	memory, err := mem.VirtualMemoryWithContext(nil, remoteServer)
	if err != nil {
		log.Fatalf("Failed to get memory information: %v", err)
	}

	fmt.Println("Total memory:", memory.Total)
	fmt.Println("Available memory:", memory.Available)
	fmt.Println("Used memory:", memory.Used)

	// Get process information
	processes, err := process.ProcessesWithContext(nil, remoteServer)
	if err != nil {
		log.Fatalf("Failed to get process information: %v", err)
	}

	for _, p := range processes {
		pid := p.Pid
		ppid, _ := p.Ppid()

		fmt.Printf("Process ID: %d, Parent Process ID: %d\n", pid, ppid)
	}
}
