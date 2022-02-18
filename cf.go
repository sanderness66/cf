// CF -- count files in directories and subdirectories (like du(1) but
// for counting files rather than file sizes)
//
// SvM 30-JAN-2021 - 21-JAN-2022
//
// created from https://raw.githubusercontent.com/missedone/dugo/master/du.go
// and https://golang.org/pkg/os/#Stat by chipping away anrything we
// don't need

package main

import (
	"flag"
	"fmt"
	"os"
	"path"
)

var only_sum bool  // -s option
var print_tot bool // -c option
var print_sep bool // -S option

func count_files(currPath string, depth int) int64 {
	var count int64

	stat, err := os.Lstat(currPath)
	if err != nil {
		fmt.Println(err)
		return count
	}

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
		fmt.Printf("%d\t%s\n", count, path.Clean(currPath))
	}
	if print_sep {
		return 0
	} else {
		return count
	}
}

func main() {
	var dir string
	var tot int64

	flag.BoolVar(&only_sum, "s", false, "print only a total for each argument")
	flag.BoolVar(&print_tot, "c", false, "print grand total")
	flag.BoolVar(&print_sep, "S", false, "don't add subdirectories to directory counts")
	flag.Parse()
	if flag.NArg() > 0 {
		for _, dir = range flag.Args() {
			tot += count_files(dir, 0)
		}
	} else {
		dir = "."
		tot = count_files(dir, 0)
	}

	if print_tot {
		fmt.Printf("%d\t%s\n", tot, "total")
	}
}
