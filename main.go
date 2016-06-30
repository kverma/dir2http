package main

import (
	"flag"
	"fmt"
	"net/http"
)

func main() {
	// Command line parsing
	bind := flag.String("bind", ":8080", "Bind address")

	flag.Parse()

	fileServer := http.FileServer(http.Dir("."))
	fmt.Printf("Sharing current directory on %s ...\n", *bind)
	http.ListenAndServe((*bind), fileServer)
}
