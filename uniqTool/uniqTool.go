package main

import (
  "fmt"
  "os"
  "log"
  // "strconv"
  "strings"
  "bufio"
)

func main() {
  var outputFilename string
  var data []byte
  var err error
  var filename string

  args := os.Args[1:]
  initialArg := args[0]
  

  if (initialArg == "-c") {
    
  } else {
    filename = args[0]
  }
  
    if (len(args) == 2) {
    outputFilename = args[1]
  }

  if (filename != "-") {
    data, err = os.ReadFile(filename)
    if (err != nil) {
      log.Fatal("Error reading file", err)
    }
  } else {
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
      data = append(data, scanner.Bytes()...)
      data = append(data, []byte("\n")...)
    }
  }

  var results string
  var previousLine string
  dataString := strings.Split(string(data), "\n")
  
  for i := 0; i < len(dataString); i++ {
    if (previousLine == "") {
      previousLine = dataString[i]
      results = dataString[i] + "\n"
    } else {
      if (dataString[i] != previousLine) {
        results = results + (dataString[i] + "\n")
      }
      previousLine = dataString[i]
    }
  }

  if (outputFilename != "") {
    err = os.WriteFile(outputFilename, []byte(results), 0666)
    if (err != nil) {
      log.Fatal("Error in writing to file", err)
    }
  } else {
    fmt.Println("hello?")
    fmt.Println(results)
  }
}
