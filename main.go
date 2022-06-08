package main

import (
	"fmt"
	"runtime/debug"
)

func main() {
	buildInfo, ok := debug.ReadBuildInfo()
	if !ok {
		fmt.Println("Failed to get version.")
		return
	}
	fmt.Printf("Version: %s", buildInfo.Main.Version)
}
