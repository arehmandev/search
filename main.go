package search

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
)
// Search - this is a function for searching a directory for a file/folder/both
func Search(pattern, pathforsearching, filetype string) []string {

	var files []string

	if filetype == "all" {
		filepath.Walk(pathforsearching, func(path string, f os.FileInfo, _ error) error {
			// if !f.IsDir() {
			r, err := regexp.MatchString(pattern, f.Name())
			if err == nil && r {
				absolutefilepath, err := filepath.Abs(path)
				if err != nil {
					log.Fatal(err)
				}
				files = append(files, absolutefilepath)
				// }
			}
			return nil
		})
	}

	if filetype == "file" {
		filepath.Walk(pathforsearching, func(path string, f os.FileInfo, _ error) error {
			if !f.IsDir() {
				r, err := regexp.MatchString(pattern, f.Name())
				if err == nil && r {
					absolutefilepath, err := filepath.Abs(path)
					if err != nil {
						log.Fatal(err)
					}
					files = append(files, absolutefilepath)
				}
			}
			return nil
		})
	}

	if filetype == "folder" {
		filepath.Walk(pathforsearching, func(path string, f os.FileInfo, _ error) error {
			if f.IsDir() {
				r, err := regexp.MatchString(pattern, f.Name())
				if err == nil && r {
					absolutefilepath, err := filepath.Abs(path)
					if err != nil {
						log.Fatal(err)
					}
					files = append(files, absolutefilepath)
				}
			}
			return nil
		})
	}

	return files
}
