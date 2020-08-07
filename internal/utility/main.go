package utility

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{Use: "doc-gen"}
)

func init() {
	commands := getCommands()
	for _, command := range commands {
		rootCmd.AddCommand(command)
	}
}

func getCommands() []*cobra.Command {
	return []*cobra.Command{
		&cobra.Command{
			Use:   "full [no options!]",
			Short: "Create default lab A4 doc with title page",
			Long:  `Create default latex document with title page, on A4 paper, with folder for images.`,
			Run:   CopyFilesWrapper("full", []string{"images"}),
		},
		&cobra.Command{
			Use:   "simple [no options!]",
			Short: "Create default LaTeX document",
			Long:  `Create default LaTeX document, minimal config: article, babel, indentfirst packages only!`,
			Run:   CopyFilesWrapper("simple", []string{}),
		},
	}
}

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}
