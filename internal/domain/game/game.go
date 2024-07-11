package game

type Game struct {
	ID     uint64 `json:"id"`
	Status string `json:"status"`
	Words  map[uint]Word
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
