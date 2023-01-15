package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
)

func handler(w http.ResponseWriter, r *http.Request) {
    // Log the request
    log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)

    // Write the response
    fmt.Fprintf(w, "Bienvenidos a Moonlight Falls!")
}

func main() {
    // Create a log file
    f, err := os.OpenFile("moonlight.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
    if err != nil {
        log.Fatalf("Error opening log file: %v", err)
    }
    defer f.Close()

    // Use the log file for logging
    log.SetOutput(f)

    // Handle requests
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
