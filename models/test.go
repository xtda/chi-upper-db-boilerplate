package models

import (
	"context"
	"errors"
	"log"

	"upper.io/db.v3/lib/sqlbuilder"
)

type Test struct {
	Test string `db:"test"`
}

func AllTest(ctx context.Context) ([]Test, error) {
	db, ok := ctx.Value("db").(sqlbuilder.Database)
	if !ok {
		return nil, errors.New("models: could not get database connection pool from context")
	}
	var tests []Test
	data := db.Collection("test").Find()
	if err := data.OrderBy("test desc").All(&tests); err != nil {
		log.Fatal(err)
	}
	return tests, nil
}
