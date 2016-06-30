package main

import (
	"flag"
	"fmt"
	"net/http"
	//"os"
)

func main() {

	// Get current working directory to get the file from it
	/* cwd, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error while getting current directory.")
		return
	}

	*/
	// Command line parsing
	bind := flag.String("bind", ":80", "Bind address")
	// root_folder := flag.String("root", cwd, "Root folder")

	flag.Parse()

	fileServer := http.FileServer(http.Dir("."))
	fmt.Printf("Sharing current directory on %s ...\n", *bind)
	http.ListenAndServe((*bind), fileServer)

}
