package someinterface

import (
	"strconv"
	"strings"
)

type IPAddr [4]byte

func (ip IPAddr) String() string {
	s := make([]string, 0, len(ip))
	for _, i := range ip {
		s = append(s, strconv.Itoa(int(i)))
	}
	return strings.Join(s, ".")
}
