package detector

func Detect() string {
	done := make(chan string)
	go detectAWS(done)
	go detectGCE(done)
	go detectDigitalOcean(done)
	go detectAzure(done)
	go detectVultr(done)

	n := 5 // total number of go routines
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
