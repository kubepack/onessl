package cmds

import (
	"crypto/x509"
	"fmt"
	"os"

	"github.com/appscode/go/log"
	"github.com/spf13/cobra"
	"k8s.io/client-go/util/cert"
	kubeadmapi "k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm/v1alpha1"
	kubeadmconsts "k8s.io/kubernetes/cmd/kubeadm/app/constants"
)

func NewCmdCreateClient() *cobra.Command {
	var (
		certDir   = kubeadmapi.DefaultCertificatesDir
		org       string
		overwrite bool
	)
	cmd := &cobra.Command{
		Use:               "client-cert",
		Short:             "Generate client certificate pair",
		DisableAutoGenTag: true,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				log.Fatalln("Missing client name.")
			}
			if len(args) > 1 {
				log.Fatalln("Multiple client name found.")
			}

			cfg := cert.Config{
				CommonName:   args[0],
				Usages:       []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth},
				Organization: []string{org},
			}

			store, err := NewCertStore(certDir)
			if err != nil {
				fmt.Printf("Failed to create certificate store. Reason: %v.", err)
				os.Exit(1)
			}
			if store.IsExists(store.Filename(cfg)) && overwrite {
				fmt.Printf("Client certificate found at %s. Do you want to overwrite?", store.Location())
				os.Exit(1)
			}

			if !store.PairExists(kubeadmconsts.CACertAndKeyBaseName) {
				fmt.Printf("CA certificates not found in %s. Run `guard init ca`", store.Location())
				os.Exit(1)
			}
			caCert, caKey, err := store.Read(kubeadmconsts.CACertAndKeyBaseName)
			if err != nil {
				fmt.Printf("Failed to load ca certificate. Reason: %v.", err)
				os.Exit(1)
			}

			key, err := cert.NewPrivateKey()
			if err != nil {
				fmt.Printf("Failed to generate private key. Reason: %v.", err)
				os.Exit(1)
			}
			cert, err := cert.NewSignedCert(cfg, key, caCert, caKey)
			if err != nil {
				fmt.Printf("Failed to generate server certificate. Reason: %v.", err)
				os.Exit(1)
			}
			err = store.Write(store.Filename(cfg), cert, key)
			if err != nil {
				fmt.Printf("Failed to init client certificate pair. Reason: %v.", err)
				os.Exit(1)
			}
			fmt.Println("Wrote client certificates in ", store.Location())
			os.Exit(0)
		},
	}

	cmd.Flags().StringVar(&certDir, "cert-dir", certDir, "Path to directory where pki files are stored.")
	cmd.Flags().StringVarP(&org, "organization", "o", org, "Name of Organization (Github or Google).")
	cmd.Flags().BoolVar(&overwrite, "overwrite", overwrite, "Overwrite existing cert/key pair")
	return cmd
}
