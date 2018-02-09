package cmds

import (
	"fmt"
	"os"

	"github.com/hashicorp/go-version"
	"github.com/spf13/cobra"
)

func NewCmdSemver() *cobra.Command {
	var (
		base  bool
		check string
	)
	cmd := &cobra.Command{
		Use:               "semver",
		Short:             "Print sanitized semver version",
		DisableAutoGenTag: true,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) > 1 {
				Fatal(fmt.Errorf("multiple version found: %+v", args))
			}
			if len(args) == 0 {
				Fatal(fmt.Errorf("missing version"))
			}
			gitVersion := args[0]

			gv, err := version.NewVersion(gitVersion)
			if err != nil {
				Fatal(fmt.Errorf("invalid version %s. reason: %s", gitVersion, err))
			}
			m := gv.ToMutator().ResetMetadata().ResetPrerelease()
			if base {
				m = m.ResetPatch()
			}
			if check == "" {
				fmt.Print(m.String())
				return
			}

			c, err := version.NewConstraint(check)
			if err != nil {
				Fatal(fmt.Errorf("invalid constraint %s. reason: %s", gitVersion, err))
			}
			if !c.Check(m.Done()) {
				os.Exit(1)
			}
		},
	}

	cmd.Flags().BoolVar(&base, "base", base, "print major.minor.0 version")
	cmd.Flags().StringVar(&check, "check", check, "check constraint")
	return cmd
}
