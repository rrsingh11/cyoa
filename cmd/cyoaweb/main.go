package main

import (
	"cyoa"
	"flag"
	"fmt"
	"net/http"
)

func main() {
	// Get file name from flag
	fileName := flag.String("file", "gopher.json", "Choose the file name")
	port := flag.Int("port", 2969, "Enter the port number")
	flag.Parse()

	// Open the file  and read
	file := cyoa.OpenFile(*fileName)

	// Parse JSON into Story Stuct
	story := cyoa.ParseJSON(file)
	// fmt.Println(story)

	// Server
	h := cyoa.NewHandler(story)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", *port), h); err == nil {
		cyoa.Exit(fmt.Sprintf("Server running on port :%d", *port))
	} else {
		cyoa.Exit("Unable to run the server")
	}

	// log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))
	

}
