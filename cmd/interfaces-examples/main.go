package main

import (
	"fmt"

	"github.com/JJoddyZZ/golang-examples/someinterface"
)

func main() {
	hosts := map[string]someinterface.IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%s: %s\n", name, ip.String())
	}
}
