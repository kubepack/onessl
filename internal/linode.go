package internal

import (
	"net"
	"strings"

	n2 "github.com/appscode/go/net"
)

func DetectLinode(done chan<- string) {
	externalIPs, _, err := n2.HostIPs()
	if err != nil {
		done <- ""
		return
	}
	for _, ip := range externalIPs {
		names, err := net.LookupAddr(ip.String())
		if err == nil {
			for _, name := range names {
				if strings.HasSuffix(name, ".members.linode.com.") {
					done <- "linode"
					return
				}
			}
		}
	}
	done <- ""
}
