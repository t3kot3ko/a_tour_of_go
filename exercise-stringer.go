package main

import "fmt"

type IPAddr [4]byte

func (addr IPAddr) String() string {
	// return strings.Join(addr, ".")
	return fmt.Sprintf("%d.%d.%d.%d", addr[0], addr[1], addr[2], addr[3])
}

func main() {
	addr := IPAddr{123, 23, 34, 45}
	fmt.Println(addr)
}
