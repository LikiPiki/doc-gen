package utility

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

const (
	imagesFolderName = "images"
)

var (
	rootCmd    = &cobra.Command{Use: "doc-gen"}
	fullDocCmd = &cobra.Command{
		Use:   "full [no options!]",
		Short: "Create default lab A4 doc with title page",
		Long:  `Create default latex document with title page, on A4 paper, with folder for images.`,
		Run:   createFullDoc,
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(fullDocCmd)
}

func createFullDoc(cmd *cobra.Command, args []string) {
	fmt.Println("Creating full doc...")
	name, _ := os.Executable()

	fmt.Println(name)

	_, err := os.Stat(imagesFolderName)

	if os.IsNotExist(err) {
		errDir := os.MkdirAll(imagesFolderName, 0755)
		if errDir != nil {
			log.Fatal(err)
			return
		}
		fmt.Println("image folder created")
	} else {
		fmt.Println("image folder is exists")
	}
}
