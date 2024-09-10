package pg

import (
	"database/sql"
	"fmt"
	"github.com/boliev/vkwordle/internal/domain/game"
	"sync"
)

type Noun struct {
	DB    *sql.DB
	Mu    sync.RWMutex
	cache map[int]map[string]interface{}
}

func NewNoun(db *sql.DB) *Noun {
	return &Noun{
		DB:    db,
		Mu:    sync.RWMutex{},
		cache: make(map[int]map[string]interface{}),
	}
}

func (n *Noun) IsWordReal(word string, gameType int) (bool, error) {
	err := n.fillCacheIfNeed()
	if err != nil {
		return false, err
	}

	if gameType != game.TYPE_5_WORDS {
		return false, fmt.Errorf("wrong game type %d", gameType)
	}

	n.Mu.RLock()
	defer n.Mu.RUnlock()
	_, ok := n.cache[gameType][word]

	return ok, nil
}

func (n *Noun) fillCacheIfNeed() error {
	if len(n.cache) > 0 {
		return nil
	}

	n.cache[game.TYPE_5_WORDS] = make(map[string]interface{})
	const s = "SELECT word FROM nouns WHERE type = $1"

	rows, err := n.DB.Query(s, game.TYPE_5_WORDS)
	if err != nil {
		return err
	}
	defer rows.Close()

	n.Mu.Lock()
	defer n.Mu.Unlock()
	for rows.Next() {
		var word string
		err = rows.Scan(&word)
		if err != nil {
			return err
		}

		n.cache[game.TYPE_5_WORDS][word] = nil
	}

	return nil
}
