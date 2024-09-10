package main

import (
	"database/sql"
	"fmt"
	"github.com/boliev/vkwordle/internal/app/api"
	"github.com/boliev/vkwordle/internal/domain/game"
	"github.com/boliev/vkwordle/internal/infrastucture/repository/pg"
	migrations_fs "github.com/boliev/vkwordle/migrations"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	log "github.com/sirupsen/logrus"
	"os"
)

func main() {
	createLogger()
	db, err := prepareDB()
	if err != nil {
		panic(err)
	}

	puzzleRepo := pg.NewPuzzle(db)
	gameRepo := pg.NewGame(db)
	nounRepo := pg.NewNoun(db)

	gameService := game.NewService(puzzleRepo, gameRepo, nounRepo)
	app := api.NewApi(gameService)
	app.Run()
}

func prepareDB() (*sql.DB, error) {
	dbString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		"localhost", 5532, "admin", "123456", "vkwordle")

	db, err := sql.Open("postgres", dbString)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	log.Infoln("DB created")

	err = migrate(db)

	if err != nil {
		return nil, err
	}

	return db, nil
}

func createLogger() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetReportCaller(true)
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
	log.Infoln("logger is ready")
}

func migrate(db *sql.DB) error {
	if err := goose.SetDialect("postgres"); err != nil {
		return err
	}

	goose.SetBaseFS(migrations_fs.EmbedMigrations)

	if err := goose.Up(db, "."); err != nil {
		return err
	}

	log.Infoln("DB migrated")
	return nil
}
