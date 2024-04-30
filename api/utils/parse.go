package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

func ParseHtmlTemplates(root string) *template.Template {
	var directories []string
	var filenames []string

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			directories = append(directories, path)
		} else {
			filename := info.Name()
			hasRepeatedFiles := contains(filenames, filename)
			if hasRepeatedFiles {
				return fmt.Errorf("you can't have repeated template files: %s", filename)
			}

			filenames = append(filenames, filename)
		}

		return nil
	})

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	tmpl := template.Must(template.ParseGlob(root + "/*.html"))
	for _, path := range directories[1:] {
		template.Must(tmpl.ParseGlob(path + "/*.html"))
	}

	return tmpl
}

func contains(filenames []string, filename string) bool {
	for _, f := range filenames {
		if f == filename {
			return true
		}
	}
	return false
}
