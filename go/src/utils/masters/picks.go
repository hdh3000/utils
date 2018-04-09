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
	AllPicks        []*RankedPlayer
}

type Choice struct {
	OptionGroup string
	Picks       []*RankedPlayer
}

type RankedPlayer struct {
	Name            string // LastName
	Ref             string
	Score           int
	Position        int    // 4
	DisplayPosition string // T4
	Thru            int
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
					{Ref: "Woods"},
				},
			},
			{
				OptionGroup: "Group B",
				Picks: []*RankedPlayer{
					{Ref: "McIlroy"},
					{Ref: "Rose"},
				},
			},
			{
				OptionGroup: "Group C",
				Picks: []*RankedPlayer{
					{Ref: "Watson"},
					{Ref: "Casey"},
					{Ref: "Hoffman"},
				},
			},
			{
				OptionGroup: "Group D",
				Picks: []*RankedPlayer{
					{Ref: "Scott"},
					{Ref: "DeChambeau"},
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
					{Ref: "Woods"},
				},
			},
			{
				OptionGroup: "Group B",
				Picks: []*RankedPlayer{
					{Ref: "Spieth"},
					{Ref: "Fowler"},
				},
			},
			{
				OptionGroup: "Group C",
				Picks: []*RankedPlayer{
					{Ref: "Watson"},
					{Ref: "Stenson"},
					{Ref: "Kuchar"},
				},
			},
			{
				OptionGroup: "Group D",
				Picks: []*RankedPlayer{
					{Ref: "Scott"},
					{Ref: "Cabrera"},
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
					{Ref: "Woods"},
				},
			},
			{
				OptionGroup: "Group B",
				Picks: []*RankedPlayer{
					{Ref: "DJohnson"},
					{Ref: "Thomas"},
				},
			},
			{
				OptionGroup: "Group C",
				Picks: []*RankedPlayer{
					{Ref: "Oosthuizen"},
					{Ref: "Grace"},
					{Ref: "Hoffman"},
				},
			},
			{
				OptionGroup: "Group D",
				Picks: []*RankedPlayer{
					{Ref: "Scott"},
					{Ref: "Pieters"},
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
					{Ref: "Woods"},
				},
			},
			{
				OptionGroup: "Group B",
				Picks: []*RankedPlayer{
					{Ref: "Fowler"},
					{Ref: "Rose"},
				},
			},
			{
				OptionGroup: "Group C",
				Picks: []*RankedPlayer{
					{Ref: "Kuchar"},
					{Ref: "Cantlay"},
					{Ref: "Casey"},
				},
			},
			{
				OptionGroup: "Group D",
				Picks: []*RankedPlayer{
					{Ref: "Singh"},
					{Ref: "Kim"},
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
					{Ref: "Woods"},
				},
			},
			{
				OptionGroup: "Group B",
				Picks: []*RankedPlayer{
					{Ref: "Spieth"},
					{Ref: "DJohnson"},
				},
			},
			{
				OptionGroup: "Group C",
				Picks: []*RankedPlayer{
					{Ref: "Watson"},
					{Ref: "Oosthuizen"},
					{Ref: "Kuchar"},
				},
			},
			{
				OptionGroup: "Group D",
				Picks: []*RankedPlayer{
					{Ref: "Scott"},
					{Ref: "Willett"},
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
					{Ref: "Mickelson"},
				},
			},
			{
				OptionGroup: "Group B",
				Picks: []*RankedPlayer{
					{Ref: "Spieth"},
					{Ref: "DJohnson"},
				},
			},
			{
				OptionGroup: "Group C",
				Picks: []*RankedPlayer{
					{Ref: "Kuchar"},
					{Ref: "Grace"},
					{Ref: "Watson"},
				},
			},
			{
				OptionGroup: "Group D",
				Picks: []*RankedPlayer{
					{Ref: "Scott"},
					{Ref: "ZJohnson"},
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
					{Ref: "Woods"},
				},
			},
			{
				OptionGroup: "Group B",
				Picks: []*RankedPlayer{
					{Ref: "Spieth"},
					{Ref: "DJohnson"},
				},
			},
			{
				OptionGroup: "Group C",
				Picks: []*RankedPlayer{
					{Ref: "Kuchar"},
					{Ref: "Chappell"},
					{Ref: "Hatton"},
				},
			},
			{
				OptionGroup: "Group D",
				Picks: []*RankedPlayer{
					{Ref: "Scott"},
					{Ref: "Schwartzel"},
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
					{Ref: "Mickelson"},
				},
			},
			{
				OptionGroup: "Group B",
				Picks: []*RankedPlayer{
					{Ref: "Rose"},
					{Ref: "McIlroy"},
				},
			},
			{
				OptionGroup: "Group C",
				Picks: []*RankedPlayer{
					{Ref: "Noren"},
					{Ref: "Stenson"},
					{Ref: "Schauffele"},
				},
			},
			{
				OptionGroup: "Group D",
				Picks: []*RankedPlayer{
					{Ref: "Scott"},
					{Ref: "ZJohnson"},
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
					{Ref: "Woods"},
				},
			},
			{
				OptionGroup: "Group B",
				Picks: []*RankedPlayer{
					{Ref: "Spieth"},
					{Ref: "Garcia"},
				},
			},
			{
				OptionGroup: "Group C",
				Picks: []*RankedPlayer{
					{Ref: "Watson"},
					{Ref: "Fisher"},
					{Ref: "Fleetwood"},
				},
			},
			{
				OptionGroup: "Group D",
				Picks: []*RankedPlayer{
					{Ref: "Scott"},
					{Ref: "Singh"},
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
					{Ref: "Mickelson"},
				},
			},
			{
				OptionGroup: "Group B",
				Picks: []*RankedPlayer{
					{Ref: "Spieth"},
					{Ref: "Day"},
				},
			},
			{
				OptionGroup: "Group C",
				Picks: []*RankedPlayer{
					{Ref: "Kuchar"},
					{Ref: "Oosthuizen"},
					{Ref: "Watson"},
				},
			},
			{
				OptionGroup: "Group D",
				Picks: []*RankedPlayer{
					{Ref: "Vegas"},
					{Ref: "ZJohnson"},
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
					{Ref: "Woods"},
				},
			},
			{
				OptionGroup: "Group B",
				Picks: []*RankedPlayer{
					{Ref: "Thomas"},
					{Ref: "Rose"},
				},
			},
			{
				OptionGroup: "Group C",
				Picks: []*RankedPlayer{
					{Ref: "Kuchar"},
					{Ref: "Noren"},
					{Ref: "Hoffman"},
				},
			},
			{
				OptionGroup: "Group D",
				Picks: []*RankedPlayer{
					{Ref: "Scott"},
					{Ref: "Kaymer"},
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
					{Ref: "Woods"},
				},
			},
			{
				OptionGroup: "Group B",
				Picks: []*RankedPlayer{
					{Ref: "Thomas"},
					{Ref: "McIlroy"},
				},
			},
			{
				OptionGroup: "Group C",
				Picks: []*RankedPlayer{
					{Ref: "Watson"},
					{Ref: "Casey"},
					{Ref: "Hoffman"},
				},
			},
			{
				OptionGroup: "Group D",
				Picks: []*RankedPlayer{
					{Ref: "Scott"},
					{Ref: "ZJohnson"},
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
					{Ref: "Woods"},
				},
			},
			{
				OptionGroup: "Group B",
				Picks: []*RankedPlayer{
					{Ref: "Spieth"},
					{Ref: "Fowler"},
				},
			},
			{
				OptionGroup: "Group C",
				Picks: []*RankedPlayer{
					{Ref: "Hoffman"},
					{Ref: "Stenson"},
					{Ref: "Oosthuizen"},
				},
			},
			{
				OptionGroup: "Group D",
				Picks: []*RankedPlayer{
					{Ref: "Schwartzel"},
					{Ref: "Li"},
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
					{Ref: "Mickelson"},
				},
			},
			{
				OptionGroup: "Group B",
				Picks: []*RankedPlayer{
					{Ref: "Spieth"},
					{Ref: "McIlroy"},
				},
			},
			{
				OptionGroup: "Group C",
				Picks: []*RankedPlayer{
					{Ref: "Watson"},
					{Ref: "Fitzpatrick"},
					{Ref: "Berger"},
				},
			},
			{
				OptionGroup: "Group D",
				Picks: []*RankedPlayer{
					{Ref: "ZJohnson"},
					{Ref: "Cabrera"},
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
					{Ref: "Mickelson"},
				},
			},
			{
				OptionGroup: "Group B",
				Picks: []*RankedPlayer{
					{Ref: "Spieth"},
					{Ref: "Fowler"},
				},
			},
			{
				OptionGroup: "Group C",
				Picks: []*RankedPlayer{
					{Ref: "Aphibarnrat"},
					{Ref: "Finau"},
					{Ref: "Kisner"},
				},
			},
			{
				OptionGroup: "Group D",
				Picks: []*RankedPlayer{
					{Ref: "Vegas"},
					{Ref: "Ikeda"},
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
					{Ref: "Woods"},
				},
			},
			{
				OptionGroup: "Group B",
				Picks: []*RankedPlayer{
					{Ref: "Spieth"},
					{Ref: "Thomas"},
				},
			},
			{
				OptionGroup: "Group C",
				Picks: []*RankedPlayer{
					{Ref: "Oosthuizen"},
					{Ref: "Kuchar"},
					{Ref: "Casey"},
				},
			},
			{
				OptionGroup: "Group D",
				Picks: []*RankedPlayer{
					{Ref: "Scott"},
					{Ref: "Pieters"},
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
					{Ref: "Woods"},
				},
			},
			{
				OptionGroup: "Group B",
				Picks: []*RankedPlayer{
					{Ref: "DJohnson"},
					{Ref: "Thomas"},
				},
			},
			{
				OptionGroup: "Group C",
				Picks: []*RankedPlayer{
					{Ref: "Fleetwood"},
					{Ref: "Casey"},
					{Ref: "Fitzpatrick"},
				},
			},
			{
				OptionGroup: "Group D",
				Picks: []*RankedPlayer{
					{Ref: "Kizzire"},
					{Ref: "Willett"},
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
					{Ref: "Woods"},
				},
			},
			{
				OptionGroup: "Group B",
				Picks: []*RankedPlayer{
					{Ref: "Spieth"},
					{Ref: "McIlroy"},
				},
			},
			{
				OptionGroup: "Group C",
				Picks: []*RankedPlayer{
					{Ref: "Aphibarnrat"},
					{Ref: "Watson"},
					{Ref: "Casey"},
				},
			},
			{
				OptionGroup: "Group D",
				Picks: []*RankedPlayer{
					{Ref: "Scott"},
					{Ref: "Willett"},
				},
			},
		},
	},
}
