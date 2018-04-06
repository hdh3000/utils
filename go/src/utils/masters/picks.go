package masters

type LeaderBoardArgs struct {
	Contestants []*Contestant
	LastUpdated string
}

type Contestant struct {
	Name            string
	Position        int
	DisplayPosition string
	TieBreaker      int
	Score           int
	Choices         []*Choice
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

var Contestants = []*Contestant{
	{

		/*
			Tiger Woods
			Rory McIlroy
			Justin Rose
			Bubba Watson
			Paul Casey
			Charley Hoffman
			Adam Scott
			Bryson DeChambeau
		*/

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
					{Name: "McIlroy"},
					{Name: "Rose"},
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

		/*
			Tiger Woods
			Jordan Spieth
			Rickie Fowler
			Bubba Watson
			Henrik Stenson
			Matt Kuchar
			Adam Scott
			Angel Cabrera
		*/
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
		/*
			Tiger Woods
			Dustin Johnson
			Justin Thomas
			Brooks Koepka -- > Oosthuizen
			Branden Grace
			Charley Hoffman
			Adam Scott
			Thomas Pieters
		*/
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
					{Name: "Oosthuizen"},
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
		/*
			Tiger Woods
			Rickie Fowler
			Justin Rose
			Matt Kuchar
			Patrick Cantlay
			Paul Casey
			Vijay Singh
			Siwoo Kim
		*/

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
		/*
			Tiger Woods
			Jordan Spieth
			Dustin Johnson
			Bubba Watson
			Louis Oosthuizen
			Matt Kuchar
			Adam Scott
			Danny Willett
		*/
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
		/*
			Phil Mickelson
			Jordan Spieth
			Dustin Johnson
			Matt Kuchar
			Brooks Koepka --> Grace
			Bubba Watson
			Adam Scott
			Zach Johnson
		*/
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
					{Name: "Grace"},
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

		/*
			Tiger Woods
			Jordan Spieth
			Dustin Johnson
			Matt Kuchar
			Kevin Chappell
			Brooks Koepka --> hatton
			Adam Scott
			Charl Schwartzel
		*/

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
					{Name: "Spieth"},
					{Name: "Johnson"},
				},
			},
			{
				OptionGroup: "Group C",
				Picks: []*RankedPlayer{
					{Name: "Kuchar"},
					{Name: "Chappell"},
					{Name: "Hatton"},
				},
			},
			{
				OptionGroup: "Group D",
				Picks: []*RankedPlayer{
					{Name: "Scott"},
					{Name: "Schwartzel"},
				},
			},
		},
	},
	{

		/*
			Phil Mickelson
			Justin Rose
			Rory McIlroy
			Alex Noren
			Henrik Stenson
			Xander Schauffele
			Adam Scott
			Zach Johnson
		*/
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

		/*
			Tiger Woods
			Jordan Spieth
			Sergio Garcia
			Bubba Watson
			Ross Fisher
			Tommy Fleetwood
			Adam Scott
			Vijay Singh
		*/
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

		/*
			Phil Mickelson
			Jordan Spieth
			Jason Day
			Matt Kuchar
			Louis Oosthuizen
			Bubba Watson
			Jhonottan Vegas
			Zach Johnson
		*/
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
		/*
			Tiger Woods
			Justin Thomas
			Justin Rose
			Matt Kuchar
			Alex Noren
			Charley Hoffman
			Adam Scott
			Martin Kaymer
		*/
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

		/*
			Tiger Woods
			Justin Thomas
			Rory McIlroy
			Bubba Watson
			Paul Casey
			Charley Hoffman
			Adam Scott
			Zach Johnson
		*/
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

		/*
			Tiger Woods
			Jordan Spieth
			Rickie Fowler
			Charley Hoffman
			Henrik Stenson
			Louis Oosthuizen
			Charl Schwartzel
			Li Haotang
		*/
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
		/*
			Phil Mickelson
			Jordan Spieth
			Rory McIlroy
			Bubba Watson
			Matthew Fitzpatrick
			Daniel Berger
			Zach Johnson
			Angel Cabrera
		*/
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
		/*
			Phil Mickelson
			Jordan Spieth
			Rickie Fowler
			Kiradech Aphibarnrat
			Tony Finau
			Kevin Kisner
			Jhonottan Vegas
			Yuta Ikeda
		*/
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
		/*
			Tiger Woods
			Jordan Spieth
			Justin Thomas
			Louis Oosthuizen
			Matt Kuchar
			Paul Casey
			Adam Scott
			Thomas Pieters
		*/
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
					{Name: "Scott"},
					{Name: "Pieters"},
				},
			},
		},
	},
	{
		/*
			Tiger Woods
			Dustin Johnson
			Justin Thomas
			Tommy Fleetwood
			Paul Casey
			Matthew Fitzpatrick
			Patton Kizzire
			Danny Willett
		*/
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
		/*
			Tiger Woods
			Jordan Spieth
			Rory McIlroy
			Kiradech Aphibarnrat
			Bubba Watson
			Paul Casey
			Adam Scott
			Danny Willett
		*/
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
