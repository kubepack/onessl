package cmds

import (
	"crypto/x509"
	"fmt"
	"net"
	"os"

	"github.com/spf13/cobra"
	"k8s.io/client-go/util/cert"
	kubeadmapi "k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm/v1alpha1"
	kubeadmconsts "k8s.io/kubernetes/cmd/kubeadm/app/constants"
)

func NewCmdCreateServer() *cobra.Command {
	var (
		certDir = kubeadmapi.DefaultCertificatesDir
		sans    = cert.AltNames{
			IPs: []net.IP{net.ParseIP("127.0.0.1")},
		}
		overwrite bool
	)
	cmd := &cobra.Command{
		Use:               "server-cert",
		Short:             "Generate server certificate pair",
		DisableAutoGenTag: true,
		Run: func(cmd *cobra.Command, args []string) {
			cfg := cert.Config{
				CommonName: "server",
				AltNames:   sans,
				Usages:     []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
			}

			store, err := NewCertStore(certDir)
			if err != nil {
				fmt.Printf("Failed to create certificate store. Reason: %v.", err)
				os.Exit(1)
			}
			if store.IsExists(store.Filename(cfg)) && !overwrite {
				fmt.Printf("Server certificate found at %s. Do you want to overwrite?", store.Location())
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
				fmt.Printf("Failed to init server certificate pair. Reason: %v.", err)
				os.Exit(1)
			}
			fmt.Println("Wrote server certificates in ", store.Location())
			os.Exit(0)
		},
	}

	cmd.Flags().StringVar(&certDir, "cert-dir", certDir, "Path to directory where pki files are stored.")
	cmd.Flags().IPSliceVar(&sans.IPs, "ips", sans.IPs, "Alternative IP addresses")
	cmd.Flags().StringSliceVar(&sans.DNSNames, "domains", sans.DNSNames, "Alternative Domain names")
	cmd.Flags().BoolVar(&overwrite, "overwrite", overwrite, "Overwrite existing cert/key pair")
	return cmd
}
