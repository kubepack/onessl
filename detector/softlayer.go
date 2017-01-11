package detector

import (
	"net/http"

	"github.com/appscode/go/net/httpclient"
)

// https://sldn.softlayer.com/blog/jarteche/getting-started-user-data-and-post-provisioning-scripts
// https://github.com/bodenr/cci/wiki/SL-user-metadata
func detectSoftlayer(done chan<- string) {
	hc := httpclient.Default()
	resp, err := hc.Call(http.MethodGet, "https://api.service.softlayer.com/rest/v3/SoftLayer_Resource_Metadata/UserMetadata.txt", nil, nil, false)
	if err == nil &&
		(resp.StatusCode == http.StatusOK || resp.StatusCode == http.StatusNotFound) {
		done <- "softlayer"
	}
	done <- ""
}
