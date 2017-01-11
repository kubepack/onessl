package detector

import (
	"net/http"

	"github.com/appscode/go/net/httpclient"
)

// https://www.vultr.com/metadata/
func detectVultr(done chan<- string) {
	md := struct {
		Bgp struct {
			Ipv4 struct {
				MyAddress   string `json:"my-address"`
				MyAsn       string `json:"my-asn"`
				PeerAddress string `json:"peer-address"`
				PeerAsn     string `json:"peer-asn"`
			} `json:"ipv4"`
			Ipv6 struct {
				MyAddress   string `json:"my-address"`
				MyAsn       string `json:"my-asn"`
				PeerAddress string `json:"peer-address"`
				PeerAsn     string `json:"peer-asn"`
			} `json:"ipv6"`
		} `json:"bgp"`
		Hostname   string `json:"hostname"`
		Instanceid string `json:"instanceid"`
		Interfaces []struct {
			Ipv4 struct {
				Additional []struct {
					Address string `json:"address"`
					Netmask string `json:"netmask"`
				} `json:"additional"`
				Address string `json:"address"`
				Gateway string `json:"gateway"`
				Netmask string `json:"netmask"`
			} `json:"ipv4"`
			Ipv6 struct {
				Additional []struct {
					Network string `json:"network"`
					Prefix  string `json:"prefix"`
				} `json:"additional"`
				Network string `json:"network"`
				Prefix  string `json:"prefix"`
			} `json:"ipv6"`
			Mac         string `json:"mac"`
			NetworkType string `json:"network-type"`
		} `json:"interfaces"`
		PublicKeys string `json:"public-keys"`
		Region     struct {
			Regioncode string `json:"regioncode"`
		} `json:"region"`
	}{}

	hc := httpclient.Default()
	resp, err := hc.Call(http.MethodGet, "http://169.254.169.254/v1.json", nil, &md, false)
	if err == nil &&
		resp.StatusCode == http.StatusOK &&
		md.Hostname != "" &&
		md.Instanceid != "" &&
		len(md.Interfaces) > 0 {
		done <- "vultr"
	}
	done <- ""
}
