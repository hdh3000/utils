package masters

type LeaderBoardArgs struct {
	TopScore    int
	Contestants []*Contestant
}

type Contestant struct {
	Name       string
	TieBreaker int
	Score      int
	Choices    []*Choice
}

type Choice struct {
	OptionGroup string
	Picks       []*RankedPlayer
}

type RankedPlayer struct {
	Name            string // LastName
	Score           int
	Position        int    // 4
	DisplayPosition string // T4
}

var Contestents = []*Contestant{
	{
		Name:       "The Bachelor",
		TieBreaker: 274,
		Choices: []*Choice{
			{
				OptionGroup: "Group A",
				Picks: []*RankedPlayer{
					{Name: "Woods"},
				},
			},
			{
				OptionGroup: "Group B",
				Picks: []*RankedPlayer{
					{Name: "Rose"},
					{Name: "McIlroy"},
				},
			},
			{
				OptionGroup: "Group C",
				Picks: []*RankedPlayer{
					{Name: "Watson"},
					{Name: "Casey"},
					{Name: "Hoffman"},
				},
			},
			{
				OptionGroup: "Group D",
				Picks: []*RankedPlayer{
					{Name: "Scott"},
					{Name: "DeChambeau"},
				},
			},
		},
	},
	{
		Name:       "Zimmy",
		TieBreaker: 272,
		Choices: []*Choice{
			{
				OptionGroup: "Group A",
				Picks: []*RankedPlayer{
					{Name: "Woods"},
				},
			},
			{
				OptionGroup: "Group B",
				Picks: []*RankedPlayer{
					{Name: "Spieth"},
					{Name: "Fowler"},
				},
			},
			{
				OptionGroup: "Group C",
				Picks: []*RankedPlayer{
					{Name: "Watson"},
					{Name: "Stenson"},
					{Name: "Kuchar"},
				},
			},
			{
				OptionGroup: "Group D",
				Picks: []*RankedPlayer{
					{Name: "Scott"},
					{Name: "Cabrera"},
				},
			},
		},
	},
	{
		Name:       "Maksi",
		TieBreaker: 269,
		Choices: []*Choice{
			{
				OptionGroup: "Group A",
				Picks: []*RankedPlayer{
					{Name: "Woods"},
				},
			},
			{
				OptionGroup: "Group B",
				Picks: []*RankedPlayer{
					{Name: "Johnson"},
					{Name: "Thomas"},
				},
			},
			{
				OptionGroup: "Group C",
				Picks: []*RankedPlayer{
					{Name: "Koepka"},
					{Name: "Grace"},
					{Name: "Hoffman"},
				},
			},
			{
				OptionGroup: "Group D",
				Picks: []*RankedPlayer{
					{Name: "Scott"},
					{Name: "Pieters"},
				},
			},
		},
	},
	{
		Name:       "Wolffman",
		TieBreaker: 277,
		Choices: []*Choice{
			{
				OptionGroup: "Group A",
				Picks: []*RankedPlayer{
					{Name: "Woods"},
				},
			},
			{
				OptionGroup: "Group B",
				Picks: []*RankedPlayer{
					{Name: "Fowler"},
					{Name: "Rose"},
				},
			},
			{
				OptionGroup: "Group C",
				Picks: []*RankedPlayer{
					{Name: "Kuchar"},
					{Name: "Cantlay"},
					{Name: "Casey"},
				},
			},
			{
				OptionGroup: "Group D",
				Picks: []*RankedPlayer{
					{Name: "Singh"},
					{Name: "Kim"},
				},
			},
		},
	},
	{
		Name:       "Guy",
		TieBreaker: 276,
		Choices: []*Choice{
			{
				OptionGroup: "Group A",
				Picks: []*RankedPlayer{
					{Name: "Woods"},
				},
			},
			{
				OptionGroup: "Group B",
				Picks: []*RankedPlayer{
					{Name: "Spieth"},
					{Name: "Johnson"},
				},
			},
			{
				OptionGroup: "Group C",
				Picks: []*RankedPlayer{
					{Name: "Watson"},
					{Name: "Oosthuizen"},
					{Name: "Kuchar"},
				},
			},
			{
				OptionGroup: "Group D",
				Picks: []*RankedPlayer{
					{Name: "Scott"},
					{Name: "Willett"},
				},
			},
		},
	},
	{
		Name:       "Engh",
		TieBreaker: 278,
		Choices: []*Choice{
			{
				OptionGroup: "Group A",
				Picks: []*RankedPlayer{
					{Name: "Mickelson"},
				},
			},
			{
				OptionGroup: "Group B",
				Picks: []*RankedPlayer{
					{Name: "Spieth"},
					{Name: "Johnson"},
				},
			},
			{
				OptionGroup: "Group C",
				Picks: []*RankedPlayer{
					{Name: "Kuchar"},
					{Name: "Koepka"},
					{Name: "Watson"},
				},
			},
			{
				OptionGroup: "Group D",
				Picks: []*RankedPlayer{
					{Name: "Scott"},
					{Name: "Johnson"},
				},
			},
		},
	},
	{
		Name:       "Barfusen",
		TieBreaker: 283,
		Choices: []*Choice{
			{
				OptionGroup: "Group A",
				Picks: []*RankedPlayer{
					{Name: "Woods"},
				},
			},
			{
				OptionGroup: "Group B",
				Picks: []*RankedPlayer{
					{Name: "Rose"},
					{Name: "McIlroy"},
				},
			},
			{
				OptionGroup: "Group C",
				Picks: []*RankedPlayer{
					{Name: "Watson"},
					{Name: "Casey"},
					{Name: "Hoffman"},
				},
			},
			{
				OptionGroup: "Group D",
				Picks: []*RankedPlayer{
					{Name: "Scott"},
					{Name: "DeChambeau"},
				},
			},
		},
	},
	{
		Name:       "Bruce",
		TieBreaker: 278,
		Choices: []*Choice{
			{
				OptionGroup: "Group A",
				Picks: []*RankedPlayer{
					{Name: "Mickelson"},
				},
			},
			{
				OptionGroup: "Group B",
				Picks: []*RankedPlayer{
					{Name: "Rose"},
					{Name: "McIlroy"},
				},
			},
			{
				OptionGroup: "Group C",
				Picks: []*RankedPlayer{
					{Name: "Noren"},
					{Name: "Stenson"},
					{Name: "Schauffele"},
				},
			},
			{
				OptionGroup: "Group D",
				Picks: []*RankedPlayer{
					{Name: "Scott"},
					{Name: "Johnson"},
				},
			},
		},
	},
	{
		Name:       "T-Law",
		TieBreaker: 281,
		Choices: []*Choice{
			{
				OptionGroup: "Group A",
				Picks: []*RankedPlayer{
					{Name: "Woods"},
				},
			},
			{
				OptionGroup: "Group B",
				Picks: []*RankedPlayer{
					{Name: "Spieth"},
					{Name: "Garcia"},
				},
			},
			{
				OptionGroup: "Group C",
				Picks: []*RankedPlayer{
					{Name: "Watson"},
					{Name: "Fisher"},
					{Name: "Fleetwood"},
				},
			},
			{
				OptionGroup: "Group D",
				Picks: []*RankedPlayer{
					{Name: "Scott"},
					{Name: "Singh"},
				},
			},
		},
	},
	{
		Name:       "Daddy",
		TieBreaker: 276,
		Choices: []*Choice{
			{
				OptionGroup: "Group A",
				Picks: []*RankedPlayer{
					{Name: "Mickelson"},
				},
			},
			{
				OptionGroup: "Group B",
				Picks: []*RankedPlayer{
					{Name: "Spieth"},
					{Name: "Day"},
				},
			},
			{
				OptionGroup: "Group C",
				Picks: []*RankedPlayer{
					{Name: "Kuchar"},
					{Name: "Oosthuizen"},
					{Name: "Watson"},
				},
			},
			{
				OptionGroup: "Group D",
				Picks: []*RankedPlayer{
					{Name: "Vegas"},
					{Name: "Johnson"},
				},
			},
		},
	},
	{
		Name:       "Hudson",
		TieBreaker: 281,
		Choices: []*Choice{
			{
				OptionGroup: "Group A",
				Picks: []*RankedPlayer{
					{Name: "Woods"},
				},
			},
			{
				OptionGroup: "Group B",
				Picks: []*RankedPlayer{
					{Name: "Thomas"},
					{Name: "Rose"},
				},
			},
			{
				OptionGroup: "Group C",
				Picks: []*RankedPlayer{
					{Name: "Kuchar"},
					{Name: "Noren"},
					{Name: "Hoffman"},
				},
			},
			{
				OptionGroup: "Group D",
				Picks: []*RankedPlayer{
					{Name: "Scott"},
					{Name: "Kaymer"},
				},
			},
		},
	},
	{
		Name:       "Dave",
		TieBreaker: 280,
		Choices: []*Choice{
			{
				OptionGroup: "Group A",
				Picks: []*RankedPlayer{
					{Name: "Woods"},
				},
			},
			{
				OptionGroup: "Group B",
				Picks: []*RankedPlayer{
					{Name: "Thomas"},
					{Name: "McIlroy"},
				},
			},
			{
				OptionGroup: "Group C",
				Picks: []*RankedPlayer{
					{Name: "Watson"},
					{Name: "Casey"},
					{Name: "Hoffman"},
				},
			},
			{
				OptionGroup: "Group D",
				Picks: []*RankedPlayer{
					{Name: "Scott"},
					{Name: "Johnson"},
				},
			},
		},
	},
	{
		Name:       "Jeffinitely",
		TieBreaker: 281,
		Choices: []*Choice{
			{
				OptionGroup: "Group A",
				Picks: []*RankedPlayer{
					{Name: "Woods"},
				},
			},
			{
				OptionGroup: "Group B",
				Picks: []*RankedPlayer{
					{Name: "Spieth"},
					{Name: "Fowler"},
				},
			},
			{
				OptionGroup: "Group C",
				Picks: []*RankedPlayer{
					{Name: "Hoffman"},
					{Name: "Stenson"},
					{Name: "Oosthuizen"},
				},
			},
			{
				OptionGroup: "Group D",
				Picks: []*RankedPlayer{
					{Name: "Schwartzel"},
					{Name: "Li"},
				},
			},
		},
	},
	{
		Name:       "Phil for Phil",
		TieBreaker: 277,
		Choices: []*Choice{
			{
				OptionGroup: "Group A",
				Picks: []*RankedPlayer{
					{Name: "Mickelson"},
				},
			},
			{
				OptionGroup: "Group B",
				Picks: []*RankedPlayer{
					{Name: "Spieth"},
					{Name: "McIlroy"},
				},
			},
			{
				OptionGroup: "Group C",
				Picks: []*RankedPlayer{
					{Name: "Watson"},
					{Name: "Fitzpatrick"},
					{Name: "Berger"},
				},
			},
			{
				OptionGroup: "Group D",
				Picks: []*RankedPlayer{
					{Name: "Johnson"},
					{Name: "Cabrera"},
				},
			},
		},
	},
	{
		Name:       "Donald Jr",
		TieBreaker: 280,
		Choices: []*Choice{
			{
				OptionGroup: "Group A",
				Picks: []*RankedPlayer{
					{Name: "Mickelson"},
				},
			},
			{
				OptionGroup: "Group B",
				Picks: []*RankedPlayer{
					{Name: "Spieth"},
					{Name: "Fowler"},
				},
			},
			{
				OptionGroup: "Group C",
				Picks: []*RankedPlayer{
					{Name: "Aphibarnrat"},
					{Name: "Finau"},
					{Name: "Kisner"},
				},
			},
			{
				OptionGroup: "Group D",
				Picks: []*RankedPlayer{
					{Name: "Vegas"},
					{Name: "Ikeda"},
				},
			},
		},
	},

	{
		Name:       "Hank",
		TieBreaker: 278,
		Choices: []*Choice{
			{
				OptionGroup: "Group A",
				Picks: []*RankedPlayer{
					{Name: "Woods"},
				},
			},
			{
				OptionGroup: "Group B",
				Picks: []*RankedPlayer{
					{Name: "Spieth"},
					{Name: "Thomas"},
				},
			},
			{
				OptionGroup: "Group C",
				Picks: []*RankedPlayer{
					{Name: "Oosthuizen"},
					{Name: "Kuchar"},
					{Name: "Casey"},
				},
			},
			{
				OptionGroup: "Group D",
				Picks: []*RankedPlayer{
					{Name: "Kizzire"},
					{Name: "Willett"},
					{Name: "Scott"},
					{Name: "Pieters"},
				},
			},
		},
	},
	{
		Name:       "Bobby The D",
		TieBreaker: 273,
		Choices: []*Choice{
			{
				OptionGroup: "Group A",
				Picks: []*RankedPlayer{
					{Name: "Woods"},
				},
			},
			{
				OptionGroup: "Group B",
				Picks: []*RankedPlayer{
					{Name: "Johnson"},
					{Name: "Thomas"},
				},
			},
			{
				OptionGroup: "Group C",
				Picks: []*RankedPlayer{
					{Name: "Fleetwood"},
					{Name: "Casey"},
					{Name: "Fitzpatrick"},
				},
			},
			{
				OptionGroup: "Group D",
				Picks: []*RankedPlayer{
					{Name: "Kizzire"},
					{Name: "Willett"},
				},
			},
		},
	},
	{
		Name:       "Bobby No D",
		TieBreaker: 277,
		Choices: []*Choice{
			{
				OptionGroup: "Group A",
				Picks: []*RankedPlayer{
					{Name: "Woods"},
				},
			},
			{
				OptionGroup: "Group B",
				Picks: []*RankedPlayer{
					{Name: "Spieth"},
					{Name: "McIlroy"},
				},
			},
			{
				OptionGroup: "Group C",
				Picks: []*RankedPlayer{
					{Name: "Aphibarnrat"},
					{Name: "Watson"},
					{Name: "Casey"},
				},
			},
			{
				OptionGroup: "Group D",
				Picks: []*RankedPlayer{
					{Name: "Scott"},
					{Name: "Willett"},
				},
			},
		},
	},
}
