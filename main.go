package main

import (
	"fmt"
	"io"
	"net/http"
	"errors"
	"os"
)

func main(){
	http.HandleFunc("/", getRoot)

	err := http.ListenAndServe(":3333", nil)

	if errors.Is(err, http.ErrServerClosed){
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server %s\n", err)
		os.Exit(1)
	}
}

func getRoot(w http.ResponseWriter, r *http.Request){
	fmt.Printf("got / request\n")
	io.WriteString(w, "test")
}
