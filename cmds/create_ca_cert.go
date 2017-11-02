package cmds

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"k8s.io/client-go/util/cert"
	kubeadmapi "k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm/v1alpha1"
	kubeadmconsts "k8s.io/kubernetes/cmd/kubeadm/app/constants"
)

func NewCmdCreateCA() *cobra.Command {
	var (
		certDir   = kubeadmapi.DefaultCertificatesDir
		overwrite bool
	)
	cmd := &cobra.Command{
		Use:               "ca-cert",
		Short:             "Create CA cert/key pair",
		DisableAutoGenTag: true,
		Run: func(cmd *cobra.Command, args []string) {
			cfg := cert.Config{
				CommonName: kubeadmconsts.CACertAndKeyBaseName,
			}

			store, err := NewCertStore(certDir)
			if err != nil {
				fmt.Printf("Failed to create certificate store. Reason: %v.", err)
				os.Exit(1)
			}
			if store.IsExists(kubeadmconsts.CACertAndKeyBaseName) && !overwrite {
				fmt.Printf("CA certificate found at %s.", store.Location())
				os.Exit(1)
			}

			key, err := cert.NewPrivateKey()
			if err != nil {
				fmt.Printf("Failed to generate private key. Reason: %v.", err)
				os.Exit(1)
			}
			cert, err := cert.NewSelfSignedCACert(cfg, key)
			if err != nil {
				fmt.Printf("Failed to generate self-signed certificate. Reason: %v.", err)
				os.Exit(1)
			}
			err = store.Write(store.Filename(cfg), cert, key)
			if err != nil {
				fmt.Printf("Failed to init ca. Reason: %v.", err)
				os.Exit(1)
			}
			fmt.Println("Wrote ca certificates in ", store.Location())
			os.Exit(0)
		},
	}

	cmd.Flags().StringVar(&certDir, "cert-dir", certDir, "Path to directory where pki files are stored.")
	cmd.Flags().BoolVar(&overwrite, "overwrite", overwrite, "Overwrite existing cert/key pair")
	return cmd
}
