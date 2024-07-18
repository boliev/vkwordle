package game

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Calc_LettersStatuses(t *testing.T) {
	game := &Game{
		ID:     1,
		Puzzle: "пират",
		UserId: 123,
		Status: STATUS_IN_PROGRESS,
		Type:   TYPE_5_WORDS,
		Words:  make(map[int]*Word),
	}
	game.Words[0] = &Word{Word: "весна"}
	game.Words[1] = &Word{Word: "пирог"}
	game.Words[2] = &Word{Word: "осень"}
	game.Words[3] = &Word{Word: "пират"}

	game.Calc()

	assert.Equal(t, game.Words[0].Letters[0], LETTER_STATE_WRONG)
	assert.Equal(t, game.Words[0].Letters[1], LETTER_STATE_WRONG)
	assert.Equal(t, game.Words[0].Letters[2], LETTER_STATE_WRONG)
	assert.Equal(t, game.Words[0].Letters[3], LETTER_STATE_WRONG)
	assert.Equal(t, game.Words[0].Letters[4], LETTER_STATE_PRESENT)

	assert.Equal(t, game.Words[1].Letters[0], LETTER_STATE_IN_PLACE)
	assert.Equal(t, game.Words[1].Letters[1], LETTER_STATE_IN_PLACE)
	assert.Equal(t, game.Words[1].Letters[2], LETTER_STATE_IN_PLACE)
	assert.Equal(t, game.Words[1].Letters[3], LETTER_STATE_WRONG)
	assert.Equal(t, game.Words[1].Letters[4], LETTER_STATE_WRONG)

	assert.Equal(t, game.Words[2].Letters[0], LETTER_STATE_WRONG)
	assert.Equal(t, game.Words[2].Letters[1], LETTER_STATE_WRONG)
	assert.Equal(t, game.Words[2].Letters[2], LETTER_STATE_WRONG)
	assert.Equal(t, game.Words[2].Letters[3], LETTER_STATE_WRONG)
	assert.Equal(t, game.Words[2].Letters[4], LETTER_STATE_WRONG)

	assert.Equal(t, game.Words[3].Letters[0], LETTER_STATE_IN_PLACE)
	assert.Equal(t, game.Words[3].Letters[1], LETTER_STATE_IN_PLACE)
	assert.Equal(t, game.Words[3].Letters[2], LETTER_STATE_IN_PLACE)
	assert.Equal(t, game.Words[3].Letters[3], LETTER_STATE_IN_PLACE)
	assert.Equal(t, game.Words[3].Letters[4], LETTER_STATE_IN_PLACE)
}

func Test_IsLastWordCorrect(t *testing.T) {
	game := &Game{
		ID:     1,
		Puzzle: "пират",
		UserId: 123,
		Status: STATUS_IN_PROGRESS,
		Type:   TYPE_5_WORDS,
		Words:  make(map[int]*Word),
	}
	game.Words[0] = &Word{Word: "весна"}
	game.Words[1] = &Word{Word: "пирог"}
	game.Words[2] = &Word{Word: "осень"}
	game.Calc()
	assert.False(t, game.IsLastWordCorrect())

	game.Words[3] = &Word{Word: "пират"}
	game.Calc()
	assert.True(t, game.IsLastWordCorrect())

}
