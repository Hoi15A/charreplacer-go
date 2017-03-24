/*
*
* File: charreplace.go
* Author(s): Austin
* License: GPLv3
* Github: https://github.com/Hoi15A/charreplacer-go
*
*/

package main

import (
  "bufio"
  "bytes"
  "fmt"
  "io/ioutil"
  "os"
  "path/filepath"
  "time"
)

var count = 0

func main() {
  start := time.Now()

  fmt.Println("Starting replacement: ")
  err := filepath.Walk("./", visit)
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
