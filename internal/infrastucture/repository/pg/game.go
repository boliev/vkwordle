package pg

import (
	"database/sql"
	"errors"
	"github.com/boliev/vkwordle/internal/domain/game"
	"github.com/lib/pq"
)

type Game struct {
	DB *sql.DB
}

func NewGame(db *sql.DB) *Game {
	return &Game{
		DB: db,
	}
}

func (g *Game) CreateGame(userId int64, word string) (*game.Game, error) {
	const s = "INSERT INTO games (user_id, puzzle) VALUES ($1, $2) RETURNING id"

	res := g.DB.QueryRow(s, userId, word)

	var newId int64
	err := res.Scan(&newId)

	if err != nil {
		return nil, err
	}

	return &game.Game{
		ID:     newId,
		UserId: userId,
		Status: game.STATUS_IN_PROGRESS,
		Type:   game.TYPE_5_WORDS,
		Words:  make(map[int]*game.Word),
	}, nil
}

func (g *Game) GetActiveGame(userId int64, gameType int) (*game.Game, error) {
	const s = "SELECT id, puzzle, words FROM games WHERE user_id = $1 AND status = $2 AND type = $3 LIMIT 1"
	userGame := game.Game{}
	var words []string

	row := g.DB.QueryRow(s, userId, game.STATUS_IN_PROGRESS, gameType)

	err := row.Scan(&userGame.ID, &userGame.Puzzle, pq.Array(&words))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}
	userGame.Type = gameType
	userGame.Status = game.STATUS_IN_PROGRESS
	userGame.Words = make(map[int]*game.Word)

	for k, word := range words {
		gameWord := &game.Word{
			Word: word,
		}
		gameWord.CalcAgainst(userGame.Puzzle)
		userGame.Words[k] = gameWord
	}

	return &userGame, nil
}

func (g *Game) AddWord(gameId int64, word string) error {
	const s = "UPDATE games SET words = array_append(words, $1) WHERE id = $2"

	_, err := g.DB.Exec(s, word, gameId)
	if err != nil {
		return err
	}

	return nil
}

func (g *Game) Finish(gameId int64, status int) error {
	const s = "UPDATE games SET status = $1 WHERE id = $2"

	_, err := g.DB.Exec(s, status, gameId)
	if err != nil {
		return err
	}

	return nil
}
