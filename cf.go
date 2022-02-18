// CF -- count files in directories and subdirectories (like du(1) but
// for counting files rather than file sizes)
//
// SvM 30-JAN-2021 - 03-FEB-2021
//
// created from https://raw.githubusercontent.com/missedone/dugo/master/du.go
// and https://golang.org/pkg/os/#Stat by chipping away anrything we
// don't need

// So, do we want to count a) only files, b) everything that isn't a
// directory, or c) everything including directories? We seem to have
// ended up doing b), but a) or c) is probably more useful.

package main

import (
	"flag"
	"fmt"
	"os"
)

var only_sum bool // whether -s option is active

func count_files(currPath string, depth int) int64 {
	var count int64

	dir, err := os.Open(currPath)
	if err != nil {
		fmt.Println(err)
		return 0 // although there's a case for return 1, since it probably does exist
	}
	defer dir.Close()

	stat, err := os.Stat(currPath)
	switch mode := stat.Mode(); {
	case mode.IsDir():
		files, err := dir.Readdir(-1)
		if err != nil {
			fmt.Println(err)
		} else {
			for _, file := range files {
				if file.IsDir() {
					count += count_files(fmt.Sprintf("%s/%s", currPath, file.Name()), depth+1)
				} else {
					count++
				}
			}
		}
	default:
		count = 1
	}

	if only_sum == false || (only_sum == true && depth == 0) {
		fmt.Printf("%d\t%s\n", count, currPath)
	}
	return count
}

func main() {
	var dir string

	flag.BoolVar(&only_sum, "s", false, "display only a total for each argument")
	flag.Parse()
	if flag.NArg() > 0 {
		for _, dir = range flag.Args() {
			count_files(dir, 0)
		}
	} else {
		dir = "."
		count_files(dir, 0)
	}
}
