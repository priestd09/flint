package commands

import (
	"fmt"
	"os"

	"github.com/astrocorp42/flint/config"
	"github.com/astrocorp42/flint/lint"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "flint",
	Short: "Flint is a fast and configurable filesystem (file and directory names) linter",
	Long: `A Fast and configurable filesystem (file and directory names) linter.
More information here: https://github.com/astrocorp42/flint`,
	Run: func(cmd *cobra.Command, args []string) {
		conf, err := config.Get()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		linter := lint.NewLinter()
		issuesc, _ := linter.Lint(conf)
		for issue := range issuesc {
			fmt.Printf("%s: [%s] %s\n", issue.File.Path, issue.RuleName, issue.Message)
		}

		os.Exit(int(linter.ExitCode))
	},
}
