package db

import (
	dbConf "basesvc_v2/internal/config/db"
	"basesvc_v2/pkg/shared/enum"
	"errors"
	"log"
	"sync"

	_ "github.com/lib/pq"
)

var (
	errorInvalidDbInstance = errors.New("Invalid db instance")
)

var atomicinz uint64
var instanceDb map[string]DbDriver = make(map[string]DbDriver)

var once sync.Once

// DbDriver is object DB
type DbDriver interface {
	Db() interface{}
	// Dml(command string, params ...string) error
	// Query(command string, params ...string) (*sql.Rows, error)
}

// NewInstanceDb is used to create a new instance DB
func NewInstanceDb(config dbConf.Database) (DbDriver, error) {
	var err error
	var dbName = config.Dbname

	once.Do(func() {
		switch config.Adapter {
		case enum.Postgres:
			dbConn, sqlErr := NewPostgresDriver(config)
			if sqlErr != nil {
				err = sqlErr
				// TODO: need handle if database disconnect
				log.Fatalf("Disconnect Database %v", err)
			}
			instanceDb[dbName] = dbConn
		default:
			err = errorInvalidDbInstance
		}
	})

	return instanceDb[dbName], err
}
