package detector

import (
	"net/http"
	"strings"
	"time"

	"github.com/appscode/go/net/httpclient"
)

// http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-identity-documents.html
func detectAWS(done chan<- string) {
	md := struct {
		PrivateIP        string    `json:"privateIp"`
		AvailabilityZone string    `json:"availabilityZone"`
		AccountID        string    `json:"accountId"`
		Version          string    `json:"version"`
		InstanceID       string    `json:"instanceId"`
		InstanceType     string    `json:"instanceType"`
		ImageID          string    `json:"imageId"`
		PendingTime      time.Time `json:"pendingTime"`
		Architecture     string    `json:"architecture"`
		Region           string    `json:"region"`
	}{}

	hc := httpclient.Default()
	resp, err := hc.Call(http.MethodGet, "http://169.254.169.254/latest/dynamic/instance-identity/document", nil, &md, false)
	if err == nil &&
		resp.StatusCode == http.StatusOK &&
		strings.HasPrefix(md.ImageID, "ami-") &&
		strings.HasPrefix(md.InstanceID, "i-") {
		done <- "AWS"
	}
	done <- ""
}
