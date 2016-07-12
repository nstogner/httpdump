package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	port, ok := os.LookupEnv("PORT")
	if ok {
		log.Println("using 'PORT' env variable: ", port)
	} else {
		log.Println("no 'PORT' env variable found")
		log.Println("defaulting to port 4000")
		port = "4000"
	}
	log.Println("listening for requests...\n")
	log.Fatal(http.ListenAndServe(":"+port, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("===== NEW REQUEST =====")
		fmt.Println("time:   ", time.Now())
		fmt.Println("method: ", r.Method)
		fmt.Println("path:   ", r.URL.Path)
		fmt.Println("--- headers ---")
		for k, vs := range r.Header {
			for _, v := range vs {
				fmt.Println(k + ": " + v)
			}
		}
		fmt.Println("--- query params ---")
		for k, vs := range r.URL.Query() {
			for _, v := range vs {
				fmt.Println(k + ": " + v)
			}
		}
		fmt.Println("---  body  ---")
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Println("ERROR: COULD NOT READ REQUEST BODY!")
		}
		fmt.Println(string(body))
		fmt.Println("==== END OF REQUEST ====\n")
	})))
}
