package cloudid

import "github.com/appscode/cloudid/internal"

func Detect() string {
	done := make(chan string)
	go internal.DetectAWS(done)
	go internal.DetectGCE(done)
	go internal.DetectDigitalOcean(done)
	go internal.DetectAzure(done)
	go internal.DetectVultr(done)
	go internal.DetectLinode(done)
	go internal.DetectSoftlayer(done)
	go internal.DetectScaleway(done)

	n := 8 // total number of go routines
	i := 0
	provider := ""
	for ; i < n; i++ {
		p := <-done
		if p != provider {
			provider = p
			break
		}
	}
	if i < n {
		// run drainer
		go func() {
			for ; i < n; i++ {
				<-done
			}
		}()
	}
	return provider
}
