package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"utils/masters"
)

func main() {
	resp, err := http.Get("https://statdata.pgatour.com/r/014/leaderboard-v2mini.json")
	if err != nil {
		log.Fatal(err)
	}

	leaderBoard := masters.LeaderBoardResp{}
	if err := json.NewDecoder(resp.Body).Decode(&leaderBoard); err != nil {
		log.Fatal(err)
	}

	masters.OrderContestants(leaderBoard.Leaderboard.Players, masters.Contestents)

	for i, cont := range masters.Contestents {
		fmt.Printf("%d. %s: %d\n", i+1, cont.Name, cont.Score)
	}

}
