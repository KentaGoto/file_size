package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
)

// Returns the full path to the file
func dirwalk(dir string) []string {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	var paths []string
	for _, file := range files {
		if file.IsDir() {
			paths = append(paths, dirwalk(filepath.Join(dir, file.Name()))...)
			continue
		}
		paths = append(paths, filepath.Join(dir, file.Name()))
	}

	return paths
}

// Convert bytes to megabytes
func convertByte2MB(fis int64) int64 {
	fisMB := fis / 1048576
	return fisMB
}

func main() {
	// Exit unless there is a one argument
	if len(os.Args) != 3 {
		fmt.Println("The number of arguments specified is incorrect. Only one argument is allowed.")
		os.Exit(1)
	}

	var filesize int
	f := os.Args[1]
	filesize, _ = strconv.Atoi(f)
	dir := os.Args[2]

	// Target root directory
	paths := dirwalk(dir)

	flag := 0
	for _, path := range paths {
		// Get the file info
		fileinfo, staterr := os.Stat(path)
		if staterr != nil {
			fmt.Println(staterr)
			return
		}

		fis := fileinfo.Size()
		fisMB := convertByte2MB(fis)

		// Output if the file is larger than the specified MB
		if fisMB >= int64(filesize) {
			fmt.Println(path, "=>", fisMB, "MB")
			flag++
		} else {

		}
	}

	// Output when there is no one file larger than the specified MB
	if flag == 0 {
		fmt.Println("No files larger than", f+"MB were found.")
	}
}
