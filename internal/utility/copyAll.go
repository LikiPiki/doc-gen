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
	templateFolderName = "templates"
)

// CopyFilesWrapper - simple wrapper, which copies all files from the template folder and creates necessary dirs
func CopyFilesWrapper(templateName string, dirsToCreate []string) func(*cobra.Command, []string) {
	return func(cmd *cobra.Command, args []string) {
		baseCopier(cmd, args, templateName, dirsToCreate)
	}
}

func baseCopier(cmd *cobra.Command, args []string, templateName string, dirsToCreate []string) {
	fmt.Println("Creating " + templateName + " doc...")

	if len(dirsToCreate) > 0 {
		if err := createFoldersIfNeed(dirsToCreate); err != nil {
			log.Println("Errors while creating dir ", err)
		}
	}

	// scan template files and copy it to repo
	files, err := templates.AssetDir(
		fmt.Sprintf(
			"%s/%s",
			templateFolderName,
			templateName,
		),
	)

	if err != nil {
		log.Println(
			errors.Wrap(err, "cannot read asset directory"),
		)
		return
	}

	fmt.Println("Copy files...")
	for _, filename := range files {
		data, err := templates.Asset(
			fmt.Sprintf(
				"%s/%s/%s",
				templateFolderName,
				templateName,
				filename,
			),
		)

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

	fmt.Println("Template " + templateName + " created!")
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
		_, err := os.Stat(folder)

		if os.IsNotExist(err) {
			errDir := os.MkdirAll(folder, 0755)

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
