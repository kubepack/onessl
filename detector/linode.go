package detector

import (
	"net"
	"strings"

	net2 "github.com/appscode/go/net"
)

// https://www.vultr.com/metadata/
func detectLinode(done chan<- string) {
	for _, ip := range net2.GetExternalIPs() {
		names, err := net.LookupAddr(ip)
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
