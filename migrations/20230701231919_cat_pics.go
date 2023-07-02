package main

import (
	"database/sql"
	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upCatPics, downCatPics)
}

func upCatPics(tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	_, err := tx.Exec(`CREATE TABLE cat_pics (
 		id varchar(6) NOT NULL UNIQUE,
		cat_pic_file_name varchar(200) NOT NULL UNIQUE,
		cat_pic_data bytea,
		cat_pic_file_type varchar(10) NOT NULL,
		PRIMARY KEY(id)
    );`)
	if err != nil {
		return err
	}
	return nil
}

func downCatPics(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	_, err := tx.Exec(`DROP TABLE IF EXISTS cat_pics;`)
	if err != nil {
		return err
	}
	return nil
}
