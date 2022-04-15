package clean

import (
	"fmt"
	"os"
	"path/filepath"
)

func clean_project(path string) {

}

func getFilelist(path string) {
	err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}
		println(path)
		return nil
	})
	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}
}

func CleanRust(path string) error {
	getFilelist(path)
	return nil
}
