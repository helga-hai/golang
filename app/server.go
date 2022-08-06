package main

import (
  "fmt"
  "log"
  "net/http"
  "time"
	"io"
)

func appHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "This is my website!\n")
  fmt.Println(time.Now(), "Hello from my new fresh server")
}

func getAbout(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /about request\n")
	io.WriteString(w, "About page, HTTP!\n")
  fmt.Println(time.Now(), "Hello from my new fresh server | About")
}

func main() {
  http.HandleFunc("/", appHandler)
  http.HandleFunc("/about", getAbout)

  log.Println("Started, serving on port 8080")
  err := http.ListenAndServe(":8080", nil)

  if err != nil {
		fmt.Printf("error starting server: %s\n", err)
    log.Fatal(err.Error())
  }
}