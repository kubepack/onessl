package detector

import (
	"net/http"

	"github.com/appscode/go/net/httpclient"
)

func Detect() string {
	done := make(chan string)
	go detectAWS(done)
	go detectGCE(done)
	go detectDigitalOcean(done)

	n := 3 // total number of go routines
	i := 0
	provider := ""
	for {
		select {
		case p := <-done:
			i++
			provider = p
		}
		if i >= n {
			break
		}
	}
	return provider
}

// http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-instance-metadata.html
func detectAWS(done chan<- string) {
	hc := httpclient.Default()
	resp, err := hc.Call(http.MethodGet, "http://169.254.169.254/latest/meta-data/ami-id", nil, nil, false)
	if err == nil && resp.StatusCode == http.StatusOK {
		done <- "AWS"
	}
	done <- ""
}

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

// https://developers.digitalocean.com/documentation/metadata/#metadata-in-json
func detectDigitalOcean(done chan<- string) {
	md := struct {
		DropletID  int      `json:"droplet_id"`
		Hostname   string   `json:"hostname"`
		VendorData string   `json:"vendor_data"`
		PublicKeys []string `json:"public_keys"`
		Region     string   `json:"region"`
		Interfaces struct {
			Private []struct {
				Ipv4 struct {
					IPAddress string `json:"ip_address"`
					Netmask   string `json:"netmask"`
					Gateway   string `json:"gateway"`
				} `json:"ipv4"`
				Mac  string `json:"mac"`
				Type string `json:"type"`
			} `json:"private"`
			Public []struct {
				Ipv4 struct {
					IPAddress string `json:"ip_address"`
					Netmask   string `json:"netmask"`
					Gateway   string `json:"gateway"`
				} `json:"ipv4"`
				Ipv6 struct {
					IPAddress string `json:"ip_address"`
					Cidr      int    `json:"cidr"`
					Gateway   string `json:"gateway"`
				} `json:"ipv6"`
				Mac  string `json:"mac"`
				Type string `json:"type"`
			} `json:"public"`
		} `json:"interfaces"`
		FloatingIP struct {
			Ipv4 struct {
				Active bool `json:"active"`
			} `json:"ipv4"`
		} `json:"floating_ip"`
		DNS struct {
			Nameservers []string `json:"nameservers"`
		} `json:"dns"`
	}{}

	hc := httpclient.Default()
	resp, err := hc.Call(http.MethodGet, "http://169.254.169.254/metadata/v1.json", nil, &md, false)
	if err == nil && resp.StatusCode == http.StatusOK && md.DropletID > 0 {
		done <- "DigitalOcean"
	}
	done <- ""
}
