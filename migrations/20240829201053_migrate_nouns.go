package migrations

import (
	"bufio"
	"context"
	"database/sql"
	"fmt"
	"github.com/pressly/goose/v3"
	"os"
	"strings"
)

func init() {
	goose.AddMigrationContext(upMigrateNouns, downMigrateNouns)
}

func upMigrateNouns(ctx context.Context, tx *sql.Tx) error {
	const (
		batchSize   = 1000
		paramsCount = 2
	)
	sql := "INSERT INTO nouns (word, type) VALUES "
	dataFile, err := os.Open("migrations/data/words_all_5.txt")
	if err != nil {
		return err
	}

	defer dataFile.Close()

	scanner := bufio.NewScanner(dataFile)

	i := 0
	sqlString := strings.Builder{}
	sqlString.WriteString(sql)
	params := make([]interface{}, 0, paramsCount*batchSize)
	for scanner.Scan() {
		if i == batchSize {
			fmt.Println(sqlString.String())
			_, err := tx.ExecContext(ctx, sqlString.String(), params...)
			if err != nil {
				return err
			}
			i = 0
			sqlString = strings.Builder{}
			sqlString.WriteString(sql)
			params = make([]interface{}, 0, paramsCount*batchSize)
		}

		if i > 0 {
			sqlString.WriteString(", ")
		}

		sqlString.WriteString(fmt.Sprintf("($%d, $%d)", i*paramsCount+1, i*paramsCount+2))
		params = append(params, scanner.Text(), 5)

		i++
	}

	if len(params) > 0 {
		_, err := tx.ExecContext(ctx, sqlString.String(), params...)
		if err != nil {
			return err
		}
	}

	return nil
}

func downMigrateNouns(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	return nil
}
