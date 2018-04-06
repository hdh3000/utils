package main

import (
	"bytes"
	"log"
	"net/http"
	"time"
	"utils/masters"
)

var lastModTime = time.Now()

func main() {
	http.DefaultServeMux.HandleFunc("/styles.css", func(w http.ResponseWriter, r *http.Request) {
		styles := masters.GetStyles()
		http.ServeContent(w, r, "styles.css", lastModTime, styles)
	})

	http.DefaultServeMux.HandleFunc("/logo.png", func(w http.ResponseWriter, r *http.Request) {
		logo := masters.GetLogo()
		http.ServeContent(w, r, "logo.png", lastModTime, logo)
	})

	http.DefaultServeMux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		lb, err := masters.GetLeaderboard()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		tmpl, err := masters.RenderTmpl(lb)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		http.ServeContent(w, r, "leaderboard", time.Now(), bytes.NewReader(tmpl))
	})

	log.Println("serving on 127.0.0.1:4040")
	http.ListenAndServe("127.0.0.1:4040", http.DefaultServeMux)
}
