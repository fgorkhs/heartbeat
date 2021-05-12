package main

import (
	"fmt"
	// "os"
	humanize "github.com/dustin/go-humanize"
	memory_usage "github.com/mackerelio/go-osstat/memory"
	disk_usage "github.com/ricochet2200/go-disk-usage/du"
	linuxproc "github.com/c9s/goprocinfo/linux"
	cpu_use "github.com/shirou/gopsutil/cpu"
)

func main() {
	ram_usage, _ := memory_usage.Get()
	cpu_usage, _ := linuxproc.ReadStat("/proc/stat")
	hd_usage := disk_usage.NewDiskUsage("/home/")
	
	fmt.Println("HD Free:", humanize.Bytes(hd_usage.Free()))
	fmt.Println("HD Used:", humanize.Bytes(hd_usage.Used()))
	fmt.Printf("RAM used: %s\n", humanize.Bytes(ram_usage.Used))
	fmt.Printf("RAM free: %s\n", humanize.Bytes(ram_usage.Free))

	for n, s := range cpu_usage.CPUStats {
		fmt.Printf("CPU User %d: %d\n", n, s.User)
		// s.Nice
		// s.System
		// s.Idle
		// s.IOWait
	}
	


}