package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Run: ", runtime.GOOS)
	fmt.Println("Arch: ", runtime.GOARCH)
}
