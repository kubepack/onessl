package main

import (
	"github.com/appscode/go/net/httpclient"
	"net/http"
	"fmt"
)

func main() {
	done := make(chan string)
	go detectAWS(done)
	go detectGCE(done)

	n := 2 // total number of go routines
	i:= 0
	provider := ""
	for {
		select {
			case p := <- done:
				i++
				provider = p
		}
		if i >= n {
			break
		}
	}
	fmt.Println(provider)
}

// http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-instance-metadata.html
func detectAWS(done chan <-string) {
	hc := httpclient.Default()
	resp, err := hc.Call(http.MethodGet, "http://169.254.169.254/latest/meta-data/ami-id", nil, nil, false)
	if err == nil && resp.StatusCode == http.StatusOK {
		done <- "AWS"
	}
	done <- ""
}

// https://cloud.google.com/compute/docs/storing-retrieving-metadata#endpoints
// curl "http://metadata.google.internal/computeMetadata/v1/instance/tags" -H "Metadata-Flavor: Google"
func detectGCE(done chan <-string) {
	hc := httpclient.New(nil, map[string]string{
		"Metadata-Flavor": "Google",
	}, nil)
	resp, err := hc.Call(http.MethodGet, "http://metadata.google.internal/computeMetadata/v1/instance/tags", nil, nil, false)
	if err == nil && resp.StatusCode == http.StatusOK {
		done <- "GCE"
	}
	done <- ""
}
