package detector

import (
	"fmt"
	"net"
	"net/http"
	"time"
)

// https://www.vultr.com/metadata/
func detectScaleway(done chan<- string) {
	for port := 1; port <= 1024; port++ {
		hc := &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyFromEnvironment,
				Dial: (&net.Dialer{
					LocalAddr: &net.TCPAddr{
						Port: port,
					},
					Timeout:   5 * time.Second,
					KeepAlive: 5 * time.Second,
				}).Dial,
				TLSHandshakeTimeout: 5 * time.Second,
			},
		}
		resp, err := hc.Get("http://169.254.42.42/user_data")
		if err != nil {
			fmt.Printf("Bind to local port %d failed: %s\n", port, err.Error())
		} else {
			if resp.StatusCode == http.StatusOK {
				done <- "scaleway"
			} else {
				done <- ""
			}
			return
		}
	}
	done <- ""
}
