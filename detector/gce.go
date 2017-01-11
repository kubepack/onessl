package detector

import (
	"net/http"

	"github.com/appscode/go/net/httpclient"
)

// https://cloud.google.com/compute/docs/storing-retrieving-metadata#endpoints
// curl "http://metadata.google.internal/computeMetadata/v1/instance/tags" -H "Metadata-Flavor: Google"
func detectGCE(done chan<- string) {
	hc := httpclient.New(nil, map[string]string{
		"Metadata-Flavor": "Google",
	}, nil)
	resp, err := hc.Call(http.MethodGet, "http://metadata.google.internal/computeMetadata/v1/instance/tags", nil, nil, false)
	if err == nil && resp.StatusCode == http.StatusOK {
		done <- "GCE"
	}
	done <- ""
}
