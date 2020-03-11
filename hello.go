package main

import (
    "expvar"
    "fmt"
    "net/http"
    "io"
    "os"
)

var numCalls = expvar.NewInt("num_calls")

func HelloServer(w http.ResponseWriter, req *http.Request) {

    numCalls.Add(1)
    name, err := os.Hostname()
	if err != nil {
		panic(err)
	}

    fmt.Printf(name+" hit %d times \n", numCalls)
    io.WriteString(w, "Hello from server "+ name +"\n")
}

func main() {
    http.HandleFunc("/", HelloServer)
    http.ListenAndServe(":8080", nil)
}
