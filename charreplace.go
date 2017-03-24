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

// Valid command line flags (POSIX and GNU long-flags)
var flags = []string{"--path", "-p"}

var count = 0				// Count of files replaced
var runPath = "./"	// Default path
var argsDefined = 0

func main() {
  start := time.Now()

	if len(os.Args) > 1 {
		// Arguments were passed
		for i := 0; i < len(os.Args); i++ {
			cArg := os.Args[i]

			// Check if flag is at all valid
			if strings.HasPrefix(cArg, "-") {
				if !checkFlags(flags, cArg) {
					fmt.Println("Aborting!")
					os.Exit(1)
				}
			}

			// Define a custom path to run in
			if cArg == "--path" || cArg == "-p" {
				argsDefined++

				if len(os.Args) > i + 1 {
					runPath = os.Args[i + 1]
				} else {
					fmt.Printf("Error: Please supply a run path.")
					os.Exit(1)
				}
			}

			// Further flags can be placed here
		}
	}

	// Warn before running in current working directory
	if argsDefined == 0 {
		var confirm string
		dir, _ := os.Getwd()
		fmt.Printf("Warning: This action will possibly corrupt files in the directory %s\n", dir)
		fmt.Printf("Continue? [y,n] ")
		fmt.Scanln(&confirm)

		if(strings.ToLower(confirm) != "y") {
			fmt.Println("Aborting!")
			os.Exit(0)
		}
	}

  fmt.Println("Starting replacement: ")

	err := filepath.Walk(runPath, visit)
  fmt.Printf("filepath.Walk() returned %v\n", err)

  elapsed := time.Since(start)
  fmt.Println("Time taken: ", elapsed)
  fmt.Println("Files checked: ", count)
  fmt.Println("Push [ENTER] to exit...")
  bufio.NewReader(os.Stdin).ReadBytes('\n')
}

// Check if passed argument is within the allowed array of valid flags
func checkFlags(flags []string, flag string) bool {
	for _, a := range flags {
		if a == flag {
			// flag is valid
			return true
		}
	}

	fmt.Printf("'%s' is not a valid flag. ", flag)
	return false
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
