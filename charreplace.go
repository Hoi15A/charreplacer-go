/*
*
* File: charreplace.go
* Author(s): Austin Albrecht, Severin Fürbringer
* License: GPLv3
* Github: https://github.com/Hoi15A/charreplacer-go
*
*/

package main

import (
	"bytes"
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var count = 0				// Count of files replaced
var runpath = "./"	// Default path

func main() {
  start := time.Now()

	if len(os.Args) <= 1 {
		// Warn before running in current working directory
		var confirm string
		dir, _ := os.Getwd()
		fmt.Printf("Warning: This action will possibly corrupt files in the directory %s", dir)
		fmt.Printf("Continue [y,n] ")
		fmt.Scanln(&confirm)
		if(strings.ToLower(confirm) != "y") {
			fmt.Printf("Aborting!")
			os.Exit(0)
		}
	} else {
		// Arguments were passed
		for i := 0; i < len(os.Args); i++ {
			cArg := os.Args[i]
			if cArg == "--path" || cArg == "-p" {
				if len(os.Args) > i + 1 {
					runpath = os.Args[i + 1]
				} else {
					fmt.Printf("Error: Please supply a run path.")
					os.Exit(1)
				}
			}
		}
	}
  fmt.Println("Starting replacement: ")

	err := filepath.Walk(runpath, visit)
  fmt.Printf("filepath.Walk() returned %v\n", err)

  elapsed := time.Since(start)
  fmt.Println("Time taken: ", elapsed)
  fmt.Println("Files checked: ", count)
  fmt.Println("Push [ENTER] to exit...")
  bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func visit(path string, f os.FileInfo, err error) error {
  stat, err := os.Stat(path)
    if err != nil {
        fmt.Println(err)
        return nil
    }

  switch x := stat.Mode(); {
    case x.IsDir():
        fmt.Printf("Directory: %s\n", path)
    case x.IsRegular():

        input, err := ioutil.ReadFile(path)
        if err != nil {
          fmt.Println(err)
        }

        output := bytes.Replace(input, []byte(";"), []byte(";"), -1)
        fmt.Println(path)
        if path != "kek.exe" && path != "kek.go" {
          if err = ioutil.WriteFile(path, output, 0666); err != nil {
            fmt.Println(err)
          } else {
            fmt.Println("Replaced all semicolons in:", path)
            count++
          }
        } else {
          fmt.Println("Found ", path, " but left it alone.")
        }
    }

  return nil
}
