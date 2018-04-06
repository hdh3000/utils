package masters

type LeaderBoardResp struct {
	Debug       *Debug       `json:"debug"`
	Leaderboard *Leaderboard `json:"leaderboard"`
}

type Leaderboard struct {
	TourCode         string    `json:"tour_code"`
	TourName         string    `json:"tour_name"`
	TournamentID     string    `json:"tournament_id"`
	TournamentName   string    `json:"tournament_name"`
	StartDate        string    `json:"start_date"`
	EndDate          string    `json:"end_date"`
	TournamentFormat string    `json:"tournament_format"`
	ScoringType      string    `json:"scoring_type"`
	InCup            bool      `json:"in_cup"`
	TotalRounds      int       `json:"total_rounds"`
	IsStarted        bool      `json:"is_started"`
	IsFinished       bool      `json:"is_finished"`
	CurrentRound     int       `json:"current_round"`
	RoundState       string    `json:"round_state"`
	InPlayoff        bool      `json:"in_playoff"`
	Courses          []Course  `json:"courses"`
	CutLine          CutLine   `json:"cut_line"`
	Players          []*Player `json:"players"`
}

type Player struct {
	PlayerID        string     `json:"player_id"`
	PlayerBio       *PlayerBio `json:"player_bio"`
	CurrentPosition string     `json:"current_position"`
	StartPosition   string     `json:"start_position"`
	Status          string     `json:"status"`
	Thru            int        `json:"thru"`
	StartHole       int        `json:"start_hole"`
	CourseID        string     `json:"course_id"`
	CurrentRound    int        `json:"current_round"`
	CourseHole      int        `json:"course_hole"`
	Today           int        `json:"today"`
	Total           int        `json:"total"`
	TotalStrokes    int        `json:"total_strokes"`
	Rounds          []*Round   `json:"rounds"`
	Rankings        *Rankings  `json:"rankings"`
	GroupID         string     `json:"group_id"`
}

type Round struct {
	RoundNumber int         `json:"round_number"`
	Strokes     interface{} `json:"strokes"`
	TeeTime     string      `json:"tee_time"`
}

type Rankings struct {
	CupPoints               int         `json:"cup_points"`
	CupRank                 string      `json:"cup_rank"`
	CupTrailing             int         `json:"cup_trailing"`
	ProjectedCupPointsTotal int         `json:"projected_cup_points_total"`
	ProjectedCupPointsEvent int         `json:"projected_cup_points_event"`
	ProjectedCupRank        string      `json:"projected_cup_rank"`
	ProjectedMoneyTotal     int         `json:"projected_money_total"`
	ProjectedMoneyEvent     int         `json:"projected_money_event"`
	ProjectedMoneyRank      string      `json:"projected_money_rank"`
	PriorityProjRank        string      `json:"priority_proj_rank"`
	PriorityProjSort        interface{} `json:"priority_proj_sort"`
	PriorityStartRank       string      `json:"priority_start_rank"`
	PriorityStartSort       interface{} `json:"priority_start_sort"`
	PrioritySeed            interface{} `json:"priority_seed"`
	SchwabStartRank         string      `json:"schwab_start_rank"`
	SchwabProjRank          string      `json:"schwab_proj_rank"`
	MoneyStartRank          string      `json:"money_start_rank"`
	MoneyProjRank           string      `json:"money_proj_rank"`
	Top25Seed               interface{} `json:"top25_seed"`
	StartRank               string      `json:"start_rank"`
	ProjRank                string      `json:"proj_rank"`
	ProjSort                interface{} `json:"proj_sort"`
}

type PlayerBio struct {
	IsAmateur bool   `json:"is_amateur"`
	FirstName string `json:"first_name"`
	ShortName string `json:"short_name"`
	LastName  string `json:"last_name"`
	Country   string `json:"country"`
	IsMember  bool   `json:"is_member"`
}

type Course struct {
	CourseID      string `json:"course_id"`
	CourseCode    string `json:"course_code"`
	CourseName    string `json:"course_name"`
	IsHost        bool   `json:"is_host"`
	ParIn         string `json:"par_in"`
	ParOut        string `json:"par_out"`
	ParTotal      string `json:"par_total"`
	DistanceIn    int    `json:"distance_in"`
	DistanceOut   int    `json:"distance_out"`
	DistanceTotal int    `json:"distance_total"`
}

type CutLine struct {
	ShowCutLine          bool `json:"show_cut_line"`
	CutCount             int  `json:"cut_count"`
	CutLineScore         int  `json:"cut_line_score"`
	ShowProjected        bool `json:"show_projected"`
	ProjectedCount       int  `json:"projected_count"`
	PaidPlayersMakingCut int  `json:"paid_players_making_cut"`
}
type Debug struct {
	MsgID                         string `json:"msg_id"`
	SetupFileFound                bool   `json:"setup_file_found"`
	SetupGenerated                string `json:"setup_generated"`
	SetupYear                     string `json:"setup_year"`
	CurrentRoundInSetup           int    `json:"current_round_in_setup"`
	LastRoundInSetup              int    `json:"last_round_in_setup"`
	ScheduleFileFound             bool   `json:"schedule_file_found"`
	ScheduleGenerated             string `json:"schedule_generated"`
	TournamentInScheduleFileFound bool   `json:"tournament_in_schedule_file_found"`
	TournamentInScheduleFileName  string `json:"tournament_in_schedule_file_name"`
	FormatInScheduleFileName      string `json:"format_in_schedule_file_name"`
}

type AutoGenerated struct {
	LastUpdated string `json:"last_updated"`
	TimeStamp   string `json:"time_stamp"`
}
