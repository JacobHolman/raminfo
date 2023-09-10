package main

import (
	"fmt"
	"github.com/shirou/gopsutil/mem"
)

func main() {
	memInfo, err := mem.VirtualMemory()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Total RAM: %v GB\n", float64(memInfo.Total)/1024/1024/1024)
	fmt.Printf("Available RAM: %v GB\n", float64(memInfo.Available)/1024/1024/1024)
	fmt.Printf("Used RAM: %v GB\n", float64(memInfo.Used)/1024/1024/1024)
	fmt.Printf("Free RAM: %v GB\n", float64(memInfo.Free)/1024/1024/1024)
}