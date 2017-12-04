package models

import (
	"log"

	"upper.io/db.v3/lib/sqlbuilder"
	"upper.io/db.v3/postgresql"
	// PG adaptor
	_ "github.com/lib/pq"
)

// OpenDB open connection to databaase using upper.io/db.v3
func OpenDB(settings postgresql.ConnectionURL) (sqlbuilder.Database, error) {
	db, err := postgresql.Open(settings)
	if err != nil {
		log.Fatal(err)
	}
	return db, nil
}
