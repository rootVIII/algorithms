package main

// rootVIII
// recursively list all files in a directory (and all subdirectories)

import (
	"fmt"
	"io/ioutil"
)

func processed(fileName string, processedDirectories []string) bool {
	for i := 0; i < len(processedDirectories); i++ {
		if processedDirectories[i] != fileName {
			continue
		}
		return true
	}
	return false
}

func listDirContents(path string, dirs []string) {
	files, _ := ioutil.ReadDir(path)

	for _, f := range files {
		var newPath string
		if path != "/" {
			newPath = fmt.Sprintf("%s/%s", path, f.Name())
		} else {
			newPath = fmt.Sprintf("%s%s", path, f.Name())
		}

		if f.IsDir() {
			if !processed(newPath, dirs) {
				dirs = append(dirs, newPath)
				listDirContents(newPath, dirs)
			}
		} else {
			fmt.Println(newPath)
		}
	}
}

func main() {
	// recursively list all files in a directory (and all subdirectories)
	fmt.Printf("Recursive folder list without 'walk'")
	var dirPaths []string
	listDirContents("/home", dirPaths)
}
