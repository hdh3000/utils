package masters

type LeaderBoardArgs struct {
	TopScore    int
	Contestants []*Contestant
}

type Contestant struct {
	Name    string
	Score   int
	Choices []*Choice
}

type Choice struct {
	OptionGroup string
	Picks       []*RankedPlayer
}

type RankedPlayer struct {
	Name  string
	Score int
	Place int
}
