package cmds

import (
	"bufio"
	"crypto/rsa"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"

	"github.com/spf13/cobra"
	"k8s.io/client-go/util/cert"
)

func NewCmdGetCACert() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "get-cacert",
		Short:             "Get self-sgned CA certificate from PEM encoded RSA private key",
		DisableAutoGenTag: true,
		Run: func(cmd *cobra.Command, args []string) {
			reader := bufio.NewReader(os.Stdin)
			keyBytes, err := ioutil.ReadAll(reader)
			if err != nil {
				Fatal(fmt.Errorf("Failed to read private key. Reason: %v", err))
			}
			key, err := cert.ParsePrivateKeyPEM(keyBytes)
			if err != nil {
				Fatal(fmt.Errorf("Failed to parse private key. Reason: %v", err))
			}
			rsaKey, ok := key.(*rsa.PrivateKey)
			if !ok {
				Fatal(fmt.Errorf("Only supports rsa private key. Found %v", reflect.ValueOf(key).Kind().String()))
			}
			crt, err := cert.NewSelfSignedCACert(cert.Config{CommonName: "ca"}, rsaKey)
			if err != nil {
				Fatal(fmt.Errorf("Failed to generate self-signed certificate. Reason: %v.", err))
			}
			fmt.Println(string(cert.EncodeCertPEM(crt)))
			os.Exit(0)
		},
	}

	return cmd
}
