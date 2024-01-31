package main

import (
  "fmt"
  "os"
  "log"
  "strconv"
  "strings"
  "bufio"
)

var outputFilename string
var data []byte
var err error
var filename string
var results string
var previousLine string

func postResults(results string) {
   if (outputFilename != "") {
    fmt.Println("Writing to file " + outputFilename)
    err = os.WriteFile(outputFilename, []byte(results), 0666)
    if (err != nil) {
      log.Fatal("Error in writing to file", err)
    }
  } else {
    fmt.Println(results)
  } 
}

func parseArgs(args []string) []string {
  filename = args[0] 
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

  return strings.Split(string(data), "\n")
}

func createCounts(dataString []string) map[string]int {
  countResults := make(map[string]int)
  for i := 0; i < len(dataString); i++ {
    if (previousLine == "") {
      previousLine = dataString[i]
      countResults[dataString[i]] = 1
    } else {
      if (dataString[i] != previousLine) {
        countResults[dataString[i]] = 1
      } else {
        countResults[dataString[i]] = countResults[dataString[i]] + 1
      }
      previousLine = dataString[i]
    }
  }

  return countResults
}

func standardMode(args []string) {
  dataString := parseArgs(args)
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
  postResults(results)
}

func countMode(args []string) {
  dataString := parseArgs(args[1:])
  countResults := createCounts(dataString)

  for key, value := range countResults {
    if (key != "") {
      results = results + strconv.Itoa(value) + "\t" + key + "\n"
    }
  }
  postResults(results)
}

func repeatMode(args []string) {
  dataString := parseArgs(args[1:])
  countResults := createCounts(dataString)

  for key, value := range countResults {
    if (key != "" && value > 1) {
      results = results + key + "\n"
    }
  }
  postResults(results)
}

func uniqueMode(args []string) {
  dataString := parseArgs(args[1:])
  countResults := createCounts(dataString)

  for key, value := range countResults {
    if (key != "" && value == 1) {
      results = results + key + "\n"
    }
  }
  postResults(results)
}

func main() {
  args := os.Args[1:]
  initialArg := args[0]
  
  if (initialArg == "-c") {
    fmt.Println("Running in count mode")
    countMode(args)
  } else if (initialArg == "-d") {
    fmt.Println("Running in repeated mode")
    repeatMode(args)
  } else if (initialArg == "-u") {
    fmt.Println("Running in unique mode")
    uniqueMode(args)
  } else {
    fmt.Println("Running in standard mode")
    standardMode(args)
  }
}
