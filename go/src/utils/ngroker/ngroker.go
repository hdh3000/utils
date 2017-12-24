// package ngroker is intended to back an ngrok server
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"net/http"
)

var fPort = flag.Int("p", 8080, "the port this server will listen on")
var fStatus = flag.Int("sc", 200, "the status code to return to every request")
var fVerbose = flag.Bool("v", false, "verbose prints the entire request out.")

func main() {
	flag.Parse()
	fmt.Printf("Starting up on port: %d, with status-code: %d (verbose:%v)\n", *fPort, *fStatus, *fVerbose)
	fmt.Println("==================")
	fmt.Println("")

	handler := func(w http.ResponseWriter, r *http.Request) {
		defer w.WriteHeader(*fStatus)
		if *fVerbose {
			if err := r.ParseForm(); err != nil {
				log.Printf("ERR: %s\n", err)
				return
			}

			printer(r.Header, "HEADERS")
			printer(r.PostForm, "POST-FORM")
			printer(r.Form, "FORM")
			fmt.Println("BODY")
			b, err := ioutil.ReadAll(r.Body)
			if err != nil {
				log.Printf("ERR: %s\n", err)
				return
			}
			fmt.Println(string(b))
			fmt.Println("")
			fmt.Println("==================")
			fmt.Println("")
		}
	}

	if err := http.ListenAndServe(fmt.Sprintf("127.0.0.1:%d", *fPort), &h{handler}); err != nil {
		log.Printf("ERR: %s\n", err)
		return
	}
}

func printer(m map[string][]string, title string) {
	fmt.Println(title)
	for k, v := range m {
		fmt.Printf("%s: %s\n", k, v[0])
	}
	fmt.Println("")
}

type h struct {
	handler func(w http.ResponseWriter, r *http.Request)
}

func (h *h) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.handler(w, r)
}
