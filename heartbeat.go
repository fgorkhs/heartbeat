package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	humanize "github.com/dustin/go-humanize"
	memory_usage "github.com/mackerelio/go-osstat/memory"
	disk_usage "github.com/ricochet2200/go-disk-usage/du"
	cpu "github.com/shirou/gopsutil/cpu"
)

func usage_stats(print_to_term bool) []byte {
	ram_usage, _ := memory_usage.Get()
	hd_usage := disk_usage.NewDiskUsage("/home/")
	cpu_use, _ := cpu.Percent(time.Second, false)

	if print_to_term {
		fmt.Printf("HD Free: %s\n", humanize.Bytes(hd_usage.Free()))
		fmt.Printf("HD Used: %s\n", humanize.Bytes(hd_usage.Used()))
		fmt.Printf("RAM used: %s\n", humanize.Bytes(ram_usage.Used))
		fmt.Printf("RAM free: %s\n", humanize.Bytes(ram_usage.Free))
		fmt.Printf("CPU pct: %.1f\n", cpu_use[0])
	}

	output_map := map[string]int{
		"Time":        int(time.Now().Unix()),
		"HD_Free_mb":  int(hd_usage.Free() / 1024 / 1024),
		"HD_Used_mb":  int(hd_usage.Used() / 1024 / 1024),
		"RAM_used_mb": int(ram_usage.Used / 1024 / 1024),
		"RAM_free_mb": int(ram_usage.Free / 1024 / 1024),
		"CPU_pct":     int(cpu_use[0])}

	output_str, _ := json.Marshal(output_map)
	return output_str
}

func main() {
	jline := usage_stats(false)
	t := time.Now()
	log_file_name := fmt.Sprintf("%d-%02d-%d_system_usage.log",
		t.Year(), t.Month(), t.Day())

	f, _ := os.OpenFile(log_file_name,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()
	f.WriteString(string(jline) + "\n")
}
