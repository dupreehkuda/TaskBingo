package entity

import "time"

// GameRow is the raw projection used by the stats aggregator.
type GameRow struct {
	GameID       string
	Kind         string
	Finished     time.Time
	PackID       string
	Winner       string
	User1ID      string
	User2ID      string
	User1Bingo   int
	User2Bingo   int
	User1Numbers []int
	User2Numbers []int
	PackTitle    string
}

// Window is the time range the stats apply to.
type Window struct {
	Days int       `json:"days"`
	From time.Time `json:"from"`
	To   time.Time `json:"to"`
}

// Bucket is one bar in the chart — a day (for 7d window) or a week (for 30d).
type Bucket struct {
	Label       string `json:"label"`
	SoloBingo   int    `json:"soloBingo"`
	DuoBingo    int    `json:"duoBingo"`
	TasksClosed int    `json:"tasksClosed"`
}

// PackCount is one row in the "most played packs" ranking.
type PackCount struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Count int    `json:"count"`
}

// Summary collapses the whole window into a tile row.
type Summary struct {
	Games       int `json:"games"`
	Wins        int `json:"wins"`
	Losses      int `json:"losses"`
	TotalBingo  int `json:"totalBingo"`
	SoloBingo   int `json:"soloBingo"`
	DuoBingo    int `json:"duoBingo"`
	TasksClosed int `json:"tasksClosed"`
}

// Result is the full payload returned by the stats endpoint.
type Result struct {
	Window   Window      `json:"window"`
	Buckets  []Bucket    `json:"buckets"`
	TopPacks []PackCount `json:"topPacks"`
	Summary  Summary     `json:"summary"`
}
