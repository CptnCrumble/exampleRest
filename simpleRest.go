package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", restFace)
	serve()
}

func restFace(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		bodyReader(w, r)
	} else {
		message(w, r)
	}
}

func message(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "API Instructions")
}

func bodyReader(w http.ResponseWriter, r *http.Request) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	body := buf.String() + fmt.Sprint(" from ", getName(), "\n")
	w.Write([]byte(body))
}

func getName() string {
	if len(os.Args) == 2 {
		p := os.Args[1]
		return string(p)
	}
	return "Default name"
}

func serve() {
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe failed ", err)
	}
}
