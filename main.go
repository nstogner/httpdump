package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
)

func main() {
	port := "8080"
	if p, ok := os.LookupEnv("PORT"); ok {
		port = p
	}
	log.Printf("using PORT %v", port)
	http.ListenAndServe(":"+port, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		dmp, err := httputil.DumpRequest(r, true)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(dmp))
	}))
}
