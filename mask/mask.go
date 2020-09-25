package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s dotted-up-ip-addr\n", os.Args[0])
		os.Exit(0)
	}
	dotAddr := os.Args[1]

	addr := net.ParseIP(dotAddr)
	if addr == nil {
		fmt.Println("invlaid address")
		os.Exit(1)
	}

	mask := addr.DefaultMask()
	network := addr.Mask(mask)
	ones, bits := mask.Size()
	fmt.Println("address is ", addr.String(),
		"default mask length is ", bits,
		"leading ones count is ", ones,
		"mask is (hex)", mask.String(),
		"network is ", network.String())
	os.Exit(0)
}
