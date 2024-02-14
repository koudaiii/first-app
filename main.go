package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/joho/godotenv"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mssqldialect"
	"github.com/uptrace/bun/extra/bundebug"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	ctx := context.Background()

	mssqldbURL := os.Getenv("MSSQL_SERVER_URL")
	if mssqldbURL == "" {
		mssqldbURL = "sqlserver://sa:yourStrong(!)Password@localhost:1433?database=testdb"
	}

	// Open a MSSQL database.
	sqldb, err := sql.Open("sqlserver", mssqldbURL)
	if err != nil {
		panic(err)
	}
	defer sqldb.Close()

	// Create a  db on top of it.
	db := bun.NewDB(sqldb, mssqldialect.New())

	// Print all queries to stdout.
	db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))

	var rnd float64

	// Select a random number.
	if err := db.NewSelect().ColumnExpr("rand()").Scan(ctx, &rnd); err != nil {
		panic(err)
	}

	fmt.Println(rnd)
}
