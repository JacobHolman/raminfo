package main

import (
	"fmt"
	"github.com/shirou/gopsutil/mem"
)

const (
	Reset = "\033[0m"
	Red   = "\033[31m"
	Green = "\033[32m"
	Yellow = "\033[33m"
	Cyan  = "\033[36m"
	White = "\033[97m"
)

func progressBar(percent float64) string {
    bars := int(percent / 5)
    if bars > 20 {
        bars = 20
    }
    if bars < 0 {
        bars = 0
    }

    color := Cyan
    if percent > 70 {
        color = Red
    } else if percent > 40 {
        color = Yellow
    }

    bar := ""
    for i := 0; i < 20; i++ {
        if i < bars {
            bar += color + "█"
        } else {
            bar += White + "░"
        }
    }
    bar += Reset
    return bar
}

func main() {
	memInfo, err := mem.VirtualMemory()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	total := float64(memInfo.Total) / 1024 / 1024 / 1024
	used := float64(memInfo.Used) / 1024 / 1024 / 1024
	free := float64(memInfo.Free) / 1024 / 1024 / 1024
	avail := float64(memInfo.Available) / 1024 / 1024 / 1024
	percent := memInfo.UsedPercent

	fmt.Printf("%s┌───────────────────────────────┐%s\n", Cyan, Reset)
	fmt.Printf("%s│        RAM Information        │%s\n", Cyan, Reset)
	fmt.Printf("%s└───────────────────────────────┘%s\n", Cyan, Reset)

	fmt.Printf("  %s%-10s➤ %6.2f GB%s\n", White, "Total", total, Reset)
	fmt.Printf("  %s%-10s➤ %6.2f GB (%s%.1f%%%s)%s\n", White, "Used", used, Cyan, percent, Reset, Reset)
	fmt.Printf("  %s%-10s➤ %6.2f GB%s\n", White, "Free", free, Reset)
	fmt.Printf("  %s%-10s➤ %6.2f GB%s\n", White, "Available", avail, Reset)
	fmt.Printf("\n   %s\n", progressBar(percent))
}