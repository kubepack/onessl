package internal

import (
	"net/http"

	"github.com/appscode/go/net/httpclient"
)

// https://azure.microsoft.com/en-us/blog/what-just-happened-to-my-vm-in-vm-metadata-service/
func DetectAzure(done chan<- string) {
	md := struct {
		ID string `json:"ID"`
		UD string `json:"UD"`
		FD string `json:"FD"`
	}{}

	hc := httpclient.Default()
	resp, err := hc.Call(http.MethodGet, "http://169.254.169.254/metadata/v1/InstanceInfo", nil, &md, false)
	if err == nil && resp.StatusCode == http.StatusOK && md.ID != "" {
		done <- "azure"
	}
	done <- ""
}
