package main

import (
	"encoding/json"
	"log"
	"net/http"
	"utils/masters"
)

func main() {
	//req, _ := http.NewRequest("GET", "https://statdata.pgatour.com/r/014/leaderboard-v2mini.json", &bytes.Buffer{})
	//	req.Header.Set("Accept":"*/*"
	//"Accept-Encoding", "gzip, deflate, br"
	//"Accept-Language", "en-US,en;q=0.9"
	//"Connection", "keep-alive"
	//"Host: statdata.pgatour.com
	//"Origin: https://www.pgatour.com
	//"Referer: https://www.pgatour.com/leaderboard.html
	//"User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/65.0.3325.181 Safari/537.36)

	resp, err := http.Get("https://statdata.pgatour.com/r/014/leaderboard-v2mini.json")
	if err != nil {
		log.Fatal(err)
	}

	leaderBoard := masters.LeaderBoardResp{}
	if err := json.NewDecoder(resp.Body).Decode(&leaderBoard); err != nil {
		log.Fatal(err)
	}

	buf, err := json.MarshalIndent(&leaderBoard, "", " ")
	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(buf))

}
