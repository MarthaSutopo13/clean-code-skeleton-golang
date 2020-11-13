package db

import (
	dbConf "basesvc_v2/internal/config/db"
	dr "basesvc_v2/pkg/adapter/db"
	"database/sql"
	"log"
)

type Author struct {
	repo dr.DbDriver
	db   *sql.DB
}

func NewAuthor(databases dbConf.DatabaseList) *Author {
	driver, err := dr.NewInstanceDb(databases.Postgres.DB)
	if err != nil {
		log.Fatalf("Failed to connect : %v", err)
	}
	return &Author{
		repo: driver,
		db:   driver.Db().(*sql.DB),
	}
}

func (r *Author) Create(item interface{}) (interface{}, error) {
	//Query Db
	return nil, nil
}

func (r *Author) FindAll(item interface{}) (interface{}, error) {
	//Query Db
	return nil, nil
}
