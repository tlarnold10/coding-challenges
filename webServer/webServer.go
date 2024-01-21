package main

import (
  "fmt"
  "net/http"
  "time"
)

func getRoot(res http.ResponseWriter, req *http.Request) {
  time.Sleep(5 * time.Second)
  if (req.URL.Path != "/") {
    errorHandler(res, req, http.StatusNotFound)
    return
  }
  fmt.Println("welcome to the root")
  http.ServeFile(res, req, "./www/index.html")
}

func getAbout(res http.ResponseWriter, req *http.Request) {
  time.Sleep(5 * time.Second)
  if (req.URL.Path != "/about") {
    errorHandler(res, req, http.StatusNotFound)
    return
  }
  fmt.Println("welcome to the about section")
  http.ServeFile(res, req, "./www/about.html")
}

func errorHandler(res http.ResponseWriter, req *http.Request, status int) {
  time.Sleep(5 * time.Second)
  res.WriteHeader(status)
  if (status == http.StatusNotFound) {
    http.ServeFile(res, req, "./www/notfound.html")
  }
}

func main() {
  http.HandleFunc("/", getRoot)
  http.HandleFunc("/about", getAbout)
  err := http.ListenAndServe(":6969", nil)

  if (err != nil) {
    fmt.Println(err)
  }
}
