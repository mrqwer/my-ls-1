package ls

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
)

func Program(splitParams map[string][]string) {

}

func DefaultCase() {
	dirPath := "."
	fileInfos, err := getDirContent(dirPath)
	if err != nil {
		log.Fatal(err)
	}
	printDirContent(fileInfos)
}

func listFilesRecursive(dirPath string) error {
	return filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Print the path (file or directory)
		fmt.Println(path)

		// If it's a directory (excluding the starting directory), list its content recursively
		if info.IsDir() && path != dirPath {
			return listFilesRecursive(path)
		}

		return nil
	})
}

func getDirContent(dirPath string) ([]os.DirEntry, error) {
	fileInfos, err := os.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}

	sort.Slice(fileInfos, func(i, j int) bool {
		return fileInfos[i].Name() < fileInfos[j].Name()
	})

	return fileInfos, nil
}

func printDirContent(fileInfos []os.DirEntry) {
	for _, fileInfo := range fileInfos {
		fileName := fileInfo.Name()
		fileSize := fileInfo.Size()
		fileMode := fileInfo.Mode()
		modTime := fileInfo.ModTime().Format("Jan _2 15:04")

		// Formatting the output
		fileType := " "
		if fileMode.IsDir() {
			fileType = "d"
		}

		// File permissions
		permString := ""
		perm := fileMode.Perm()
		for i := 2; i >= 0; i-- {
			permString += perm.String()[i : i+1]
		}

		fmt.Printf("%s%s %8d %s %s\n", fileType, permString, fileSize, modTime, fileName)
	}
}
