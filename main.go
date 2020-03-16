package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)
var dir = flag.String("dir", "/tmp/www", "www root directory")
var addr = flag.String("addr", "0.0.0.0:8000", "server listen address")

func main() {
	flag.Parse()
	fileServer := http.FileServer(http.Dir(*dir))
	err := http.ListenAndServe(*addr, fileServer)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}