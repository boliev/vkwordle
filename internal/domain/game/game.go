package game

const STATUS_IN_PROGRESS = 0
const TYPE_5_WORDS = 5

const LETTER_STATE_IN_PLACE = "in place"
const LETTER_STATE_PRESENT = "present"
const LETTER_STATE_WRONG = "wrong"

type Game struct {
	ID     int64  `json:"id"`
	Puzzle string `json:"-"`
	UserId int64  `json:"user_id"`
	Status int8   `json:"status"`
	Type   int8   `json:"type"`
	Words  map[int8]*Word
}

type Word struct {
	Word    string
	Letters map[int]string
}

type Puzzle struct {
	Word     string
	Category string
	Hint     string
}

func (w *Word) CalcAgainst(puzzle string) {
	wordRune := []rune(w.Word)
	puzzleRune := []rune(puzzle)
	w.Letters = make(map[int]string)
	for i := 0; i < len(wordRune); i++ {
		w.Letters[i] = LETTER_STATE_WRONG
		if wordRune[i] == puzzleRune[i] {
			w.Letters[i] = LETTER_STATE_IN_PLACE
		} else {
			for j := 0; j < len(puzzleRune); j++ {
				if wordRune[i] == puzzleRune[j] {
					w.Letters[i] = LETTER_STATE_PRESENT
				}
			}
		}
	}
}
