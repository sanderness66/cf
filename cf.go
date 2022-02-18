// CF -- count files in directory and subdirectories (like du(1) but
// for counting files rather than file sizes)
//
// SvM 19-JAN-2021
//
// created from https://raw.githubusercontent.com/missedone/dugo/master/du.go
// by chipping away everything we don't need

package main

import (
	"flag"
	"fmt"
	"os"
)

func diskUsage(currPath string, depth int) int64 {
	var size int64

	dir, err := os.Open(currPath)
	if err != nil {
		fmt.Println(err)
		return size
	}
	defer dir.Close()

	files, err := dir.Readdir(-1)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, file := range files {
		if file.IsDir() {
			size += diskUsage(fmt.Sprintf("%s/%s", currPath, file.Name()), depth+1)
		} else {
			size++
		}
	}

	fmt.Printf("%d\t%s\n", size, currPath)
	return size
}

func main() {
	var dir string

	flag.Parse()
	if flag.NArg() > 0 {
		for _, dir = range flag.Args() {
			diskUsage(dir, 0)
		}
	} else {
		dir = "."
		diskUsage(dir, 0)
	}

}
