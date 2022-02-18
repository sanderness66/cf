// CF -- count files in directories and subdirectories (like du(1) but
// for counting files rather than file sizes)
//
// SvM 30-JAN-2021 - 11-MAR-2021
//
// created from https://raw.githubusercontent.com/missedone/dugo/master/du.go
// and https://golang.org/pkg/os/#Stat by chipping away anrything we
// don't need

package main

import (
	"flag"
	"fmt"
	"os"
)

var only_sum bool // whether -s option is active

func count_files(currPath string, depth int) int64 {
	var count int64

	stat, _ := os.Lstat(currPath)
	switch mode := stat.Mode(); {

	case mode.IsDir():
		files, err := os.ReadDir(currPath)
		if err != nil {
			fmt.Println(err)
			return count
		}

		for _, file := range files {
			if file.IsDir() {
				count += count_files(fmt.Sprintf("%s/%s", currPath, file.Name()), depth+1)
			} else if file.Type().IsRegular() {
				count++
			}
		}

	case mode.IsRegular():
		count++
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
