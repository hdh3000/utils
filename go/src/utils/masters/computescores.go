package masters

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func OrderContestants(players []*Player, c []*Contestant) error {
	playerMap := make(map[string]*RankedPlayer)
	for _, p := range players {

		var playerPosition int
		if p.CurrentPosition != "" {
			var err error
			playerPosition, err = strconv.Atoi(strings.Trim(p.CurrentPosition, "T"))
			if err != nil {
				return err
			}
		} else {
			p.CurrentPosition = "C51"
			playerPosition = 51
		}

		key := p.PlayerBio.LastName
		if p.PlayerID == "24024" {
			key = "zjohnson"
		} else if p.PlayerID == "30925" {
			key = "djohnson"
		}

		playerMap[strings.ToLower(key)] = &RankedPlayer{
			Name:            fmt.Sprintf("%s. %s", strings.ToLower(p.PlayerBio.ShortName), strings.ToLower(p.PlayerBio.LastName)),
			Score:           p.Total,
			Position:        playerPosition,
			DisplayPosition: p.CurrentPosition,
			Thru:            p.Thru,
		}

	}

	for _, cont := range c {
		var allPicks []*RankedPlayer
		for _, choice := range cont.Choices {
			for i, rp := range choice.Picks {
				pick, ok := playerMap[strings.ToLower(rp.Name)]
				if !ok {
					return fmt.Errorf("%s is not a player", strings.ToLower(rp.Name))
				}
				choice.Picks[i] = pick
			}

			allPicks = append(allPicks, choice.Picks...)
		}

		sort.Slice(allPicks, func(i, j int) bool {
			// Sort by position
			if allPicks[i].Position == allPicks[j].Position {
				return allPicks[i].Name < allPicks[j].Name
			}

			return allPicks[i].Position < allPicks[j].Position

		})

		cont.AllPicks = allPicks
		cont.Score = ComputeScore(cont)
	}

	sort.Slice(c, func(i, j int) bool {
		return c[i].Score < c[j].Score
	})

	var place, prevScore int
	offset := 1
	for i, cont := range c {
		if i == 0 {
			prevScore = cont.Score
			cont.DisplayPosition = "1"
			place = 1
			continue
		}

		if cont.Score == prevScore {
			offset++
			c[i-1].DisplayPosition = fmt.Sprintf("T%d", place)
			c[i].DisplayPosition = fmt.Sprintf("T%d", place)
		} else {
			// they are not equal, so the place goes up a the number of tied players
			place += offset
			offset = 1
			prevScore = cont.Score
			c[i].DisplayPosition = fmt.Sprintf("%d", place)
		}

	}

	return nil
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
