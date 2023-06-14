package db

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/friendsofgo/errors"

	// MySQL driver
	_ "github.com/go-sql-driver/mysql"

	"github.com/htnk128/go-ddd-sample/pkg/account/domain/repository"
)

type dbConfig struct {
	host            string
	username        string
	password        string
	dbName          string
	maxIdleConns    int
	maxOpenConns    int
	connMaxLifetime time.Duration
}

func newDBConfig() (*dbConfig, error) {
	host, b := os.LookupEnv("DB_HOST")
	if !b {
		host = "127.0.0.1:3306"
	}

	userName, b := os.LookupEnv("DB_USERNAME")
	if !b {
		userName = "user"
	}

	password, b := os.LookupEnv("DB_PASSWORD")
	if !b {
		password = "password"
	}

	dbName, b := os.LookupEnv("DB_NAME")
	if !b {
		dbName = "account"
	}

	dmic, b := os.LookupEnv("DB_MAX_IDLE_CONNS")
	if !b {
		dmic = "10"
	}

	maxIdleConns, err := strconv.Atoi(dmic)
	if err != nil {
		return nil, errors.Wrap(err, "env of DB_MAX_IDLE_CONNS is not numeric.")
	}

	dmoc, b := os.LookupEnv("DB_MAX_OPEN_CONNS")
	if !b {
		dmoc = "20"
	}

	maxOpenConns, err := strconv.Atoi(dmoc)
	if err != nil {
		return nil, errors.Wrap(err, "env of DB_MAX_OPEN_CONNS is not numeric.")
	}

	dcml, b := os.LookupEnv("DB_CONN_MAX_LIFETIME")
	if !b {
		dcml = "5000"
	}

	connMaxLifetime, err := strconv.Atoi(dcml)
	if err != nil {
		return nil, errors.Wrap(err, "env of DB_CONN_MAX_LIFETIME is not numeric.")
	}

	return &dbConfig{
		host:            host,
		username:        userName,
		password:        password,
		dbName:          dbName,
		maxIdleConns:    maxIdleConns,
		maxOpenConns:    maxOpenConns,
		connMaxLifetime: time.Duration(connMaxLifetime) * time.Second,
	}, nil
}

type Repositories struct {
	AccountRepository repository.AccountRepository
}

const driverName = "mysql"

func NewRepositories() (*Repositories, error) {
	config, err := newDBConfig()
	if err != nil {
		return nil, errors.Wrap(err, "failed to create database configuration.")
	}

	dns := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", config.username, config.password, config.host, config.dbName)
	db, err := sql.Open(driverName, dns)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create database connection.")
	}

	db.SetMaxIdleConns(config.maxIdleConns)
	db.SetMaxOpenConns(config.maxOpenConns)
	db.SetConnMaxLifetime(config.connMaxLifetime)

	return &Repositories{
		AccountRepository: newUserRepository(db),
	}, nil
}
