package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
)

// ルートディレクトリから再帰で潜ってファイルのフルパスを返す
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
	// 引数が1つ以外は終了
	if len(os.Args) != 3 {
		fmt.Println("The number of arguments specified is incorrect. Only one argument is allowed.")
		os.Exit(1)
	}

	var filesize int
	f := os.Args[1]
	filesize, _ = strconv.Atoi(f)
	//fmt.Println(filesize)
	dir := os.Args[2]

	// 対象のルートディレクトリ
	paths := dirwalk(dir)

	flag := 0

	for _, path := range paths {
		// ファイル情報を得る
		fileinfo, staterr := os.Stat(path)
		//fmt.Println(fileinfo.Size())
		if staterr != nil {
			fmt.Println(staterr)
			return
		}

		fis := fileinfo.Size()
		fisMB := convertByte2MB(fis)

		// 指定のMB以上のファイルがあれば出力
		if fisMB >= int64(filesize) {
			fmt.Println(path, "=>", fisMB, "MB")
			flag++
		} else {

		}
	}

	// 指定のMB以上のファイルが一つも無ければ出力
	if flag == 0 {
		fmt.Println("No files larger than", f+"MB were found.")
	}
}
