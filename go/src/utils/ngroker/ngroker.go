// package ngroker is intended to back an ngrok server
package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"fs/fshttp"
	"io"
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
		r.ParseForm()
		writer := &bytes.Buffer{}
		tee := io.TeeReader(r.Body, writer)
		decoder := json.NewDecoder(tee)
		req := dFlowFufillmentReq{}
		if err := decoder.Decode(&req); err != nil {
			panic(err)
		}

		rsp := dFlowResp{
			FulfillmentText: req.QueryResult.FulfillmentText,
		}

		fmt.Println(writer.String())

		fshttp.WriteJson(w, rsp)

		//if *fVerbose {
		//	if err := r.ParseForm(); err != nil {
		//		log.Printf("ERR: %s\n", err)
		//		return
		//	}
		//
		//	printer(r.Header, "HEADERS")
		//	printer(r.PostForm, "POST-FORM")
		//	printer(r.Form, "FORM")
		//	fmt.Println("BODY")
		//	b, err := ioutil.ReadAll(r.Body)
		//	if err != nil {
		//		log.Printf("ERR: %s\n", err)
		//		return
		//	}
		//	fmt.Println(string(b))
		//	fmt.Println("")
		//	fmt.Println("==================")
		//	fmt.Println("")
		//}
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

type dFlowFufillmentReq struct {
	ResponseID  string `json:"responseId"`
	QueryResult struct {
		QueryText                string `json:"queryText"`
		Parameters               map[string]interface{}
		AllRequiredParamsPresent bool   `json:"allRequiredParamsPresent"`
		FulfillmentText          string `json:"fulfillmentText"`
		FulfillmentMessages      []struct {
			Text struct {
				Text []string `json:"text"`
			} `json:"text"`
			Platform string `json:"platform,omitempty"`
		} `json:"fulfillmentMessages"`
		Intent struct {
			Name        string `json:"name"`
			DisplayName string `json:"displayName"`
		} `json:"intent"`
		IntentDetectionConfidence float64 `json:"intentDetectionConfidence"`
		DiagnosticInfo            struct {
		} `json:"diagnosticInfo"`
		LanguageCode string `json:"languageCode"`
	} `json:"queryResult"`
	OriginalDetectIntentRequest struct {
		Payload struct {
			Data struct {
				AuthedUsers []string `json:"authed_users"`
				EventID     string   `json:"event_id"`
				APIAppID    string   `json:"api_app_id"`
				TeamID      string   `json:"team_id"`
				Event       struct {
					EventTs string `json:"event_ts"`
					Channel string `json:"channel"`
					Text    string `json:"text"`
					Type    string `json:"type"`
					User    string `json:"user"`
					Ts      string `json:"ts"`
				} `json:"event"`
				Type      string  `json:"type"`
				EventTime float64 `json:"event_time"`
				Token     string  `json:"token"`
			} `json:"data"`
			Source string `json:"source"`
		} `json:"payload"`
	} `json:"originalDetectIntentRequest"`
	Session string `json:"session"`
}

type dFlowResp struct {
	FulfillmentText string `json:"fulfillmentText"`
}
