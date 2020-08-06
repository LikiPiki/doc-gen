package utility

import (
	"fmt"
	"log"
	"os"

	"github.com/likipiki/doc-gen/internal/templates"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

const (
	imagesFolderName   = "images"
	templateFolderName = "template"
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

func init() {
	rootCmd.AddCommand(fullDocCmd)
}

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func createFullDoc(cmd *cobra.Command, args []string) {
	fmt.Println("Creating full doc...")

	// scan template files and copy it to repo
	files, err := templates.AssetDir("templates/full")

	if err != nil {
		log.Println(
			errors.Wrap(err, "cannot read asset directory"),
		)
		return
	}

	fmt.Println("Copy files...")
	for _, filename := range files {
		data, err := templates.Asset("templates/full/" + filename)

		if err != nil {
			log.Println(errors.Wrap(err, "cannot read file bin data"))
			return
		}

		err = writeToFile(filename, data)
		if err != nil {
			log.Println("Error, cannot write to file", err)
		}

		fmt.Printf("- %s\n", filename)
	}

	fmt.Println("Template full created!")
}

func writeToFile(filename string, content []byte) error {
	file, err := os.Create(filename)

	if err != nil {
		return errors.Wrap(
			err,
			"can not create file - "+filename,
		)
	}
	defer file.Close()

	_, err = file.Write(content)
	if err != nil {
		return errors.Wrap(
			err,
			"unable write content to file - "+filename,
		)
	}

	return nil
}

func createFoldersIfNeed(folders []string) error {
	for _, folder := range folders {

		// create folder if not exists
		_, err := os.Stat(imagesFolderName)

		if os.IsNotExist(err) {
			errDir := os.MkdirAll(imagesFolderName, 0755)

			if errDir != nil {
				return errors.Wrap(
					err,
					"cannot make directory"+folder,
				)
			}

			fmt.Println(folder + " folder created")
		} else {
			fmt.Println(folder + " folder is exists")
		}
	}

	return nil
}
