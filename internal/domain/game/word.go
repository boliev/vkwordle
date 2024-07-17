package game

const LETTER_STATE_IN_PLACE = "in place"
const LETTER_STATE_PRESENT = "present"
const LETTER_STATE_WRONG = "wrong"

type Word struct {
	Word    string
	Letters map[int]string
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
