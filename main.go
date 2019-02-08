package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
)

func main() {
	port := "4000"
	if p, ok := os.LookupEnv("PORT"); ok {
		port = p
	}
	log.Printf("listening on PORT %v", port)

	http.ListenAndServe(":"+port, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		dmp, err := httputil.DumpRequest(r, true)
		if err != nil {
			panic(err)
		}

		fmt.Printf("================================================================================\n%v\n", string(dmp))
	}))
}
