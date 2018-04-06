package masters

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func OrderContestants(players []*Player, c []*Contestant) {
	playerMap := make(map[string]*RankedPlayer)
	maxPosition := 0
	for _, p := range players {
		playerPosition, err := strconv.Atoi(strings.Trim(p.CurrentPosition, "T"))
		if err != nil {
			panic(err)
		}

		if playerPosition > maxPosition {
			maxPosition = playerPosition
		}

		playerMap[strings.ToLower(p.PlayerBio.LastName)] = &RankedPlayer{
			Name:            strings.ToLower(p.PlayerBio.LastName),
			Score:           p.Total,
			Position:        playerPosition,
			DisplayPosition: p.CurrentPosition,
		}
		fmt.Printf("%s: %d\n", p.PlayerBio.LastName, playerPosition)
	}

	// Insert kopeka at end as he dropped out
	playerMap["koepka"] = &RankedPlayer{
		Name:     "koepka",
		Position: maxPosition,
	}

	for _, cont := range c {
		for _, choice := range cont.Choices {
			for i, rp := range choice.Picks {
				pick, ok := playerMap[strings.ToLower(rp.Name)]
				if !ok {
					panic(fmt.Sprintf("%s is not a player", strings.ToLower(rp.Name)))
				}
				choice.Picks[i] = pick
			}
		}

		cont.Score = ComputeScore(cont)
	}

	sort.Slice(c, func(i, j int) bool {
		return c[i].Score < c[j].Score
	})
}

func ComputeScore(c *Contestant) int {
	score := 0
	for _, choice := range c.Choices {
		for _, rp := range choice.Picks {
			score += rp.Position
		}
	}

	return score
}
