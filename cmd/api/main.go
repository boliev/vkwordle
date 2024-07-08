package main

import (
	"database/sql"
	"fmt"
	"github.com/boliev/vkwordle/internal/app/api"
	pg_repository "github.com/boliev/vkwordle/internal/infrastucture/repository/pg/puzzle"
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

	puzzleRepo := pg_repository.NewPuzzle(db)
	app := api.NewApi(puzzleRepo)
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

	if err := goose.Up(db, "migrations"); err != nil {
		return err
	}

	log.Infoln("DB migrated")
	return nil
}
