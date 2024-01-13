package main

import (
  "fmt"
  "os"
  "log"
  "strconv"
  "strings"
  "bufio"
)

func getLineCount(data []byte) int {
  dataString := strings.Split(string(data), "\n")
  lineCount := len(dataString)
  if (len(dataString[len(dataString) - 1]) == 0) {
    lineCount = lineCount - 1
  }
  return lineCount
}

func getWordCount(data []byte) int {
  dataString := strings.Split(string(data), " ")
  return len(dataString)
}

func getCharCount(data []byte) int {
  dataString := string(data)
  return len(dataString)
}

func getByteCount(data []byte) int {
  return len(data)
}

func main() {
  args := os.Args[1:]
  var lineCount int
  var wordCount int
  var charCount int
  var byteCount int
  var fileName string
  var mode string
  var data []byte
  var err error
  noArgs := false
  stdIn := false
  if (len(args) == 0) {
    stdIn = true
    noArgs = true
    mode = "all"
  } else {
    noArgs = false
  }
  if (!noArgs) {
    if (strings.Contains(args[0], "-") && len(args) == 1) {
      stdIn = true
    }
  }
  
  if (len(args) == 1 && !stdIn) {
    fileName = args[0]
    mode = "all"
  } else if (stdIn && !noArgs) {
    mode = args[0]
    fileName = ""
  } else if (!noArgs) {
    mode = args[0]
    fileName = args[1]
  }
  if (fileName != "") {
    data, err = os.ReadFile(fileName)
    if (err != nil) {
      log.Fatal("Error in reading file", err)
    }
  } else {
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
      data = append(data, scanner.Bytes()...)
      data = append(data, []byte("\n")...)
    }
  }
  
  switch mode {
  case "-l":
    lineCount = getLineCount(data)
  case "-c":
    byteCount = getByteCount(data)
  case "-w":
    wordCount = getWordCount(data)
  case "-m":
    charCount = getCharCount(data)
  default:
    lineCount = getLineCount(data)
    byteCount = getByteCount(data)
    wordCount = getWordCount(data)
  }
  if (byteCount != 0) {
    fmt.Print(strconv.Itoa(byteCount) + "\t")
  }
  if (lineCount != 0) {
    fmt.Print(strconv.Itoa(lineCount) + "\t")
  }
  if (wordCount != 0) {
    fmt.Print(strconv.Itoa(wordCount) + "\t")
  }
  if (charCount != 0) {
    fmt.Print(strconv.Itoa(charCount) + "\t")
  }
  fmt.Println(fileName)
}
