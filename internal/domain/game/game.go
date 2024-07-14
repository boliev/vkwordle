package game

const STATUS_IN_PROGRESS = 0
const TYPE_5_WORDS = 5

type Game struct {
	ID     int64 `json:"id"`
	UserId int64 `json:"user_id"`
	Status int8  `json:"status"`
	Type   int8  `json:"type"`
	Words  map[int8]*Word
}

type Word struct {
	Word    string
	Letters map[rune]string
}

type Puzzle struct {
	Word     string
	Category string
	Hint     string
}
