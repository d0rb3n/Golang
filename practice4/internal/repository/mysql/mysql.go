package mysql

import (
	"context"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"

	"practice4/pkg/modules"
)

type Dialect struct {
	DB *sqlx.DB
}

func NewMySQLDialect(ctx context.Context, cfg *modules.MySQLConfig) *Dialect {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DBName)

	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}

	AutoMigrate(cfg)

	return &Dialect{DB: db}
}

func AutoMigrate(cfg *modules.MySQLConfig) {
	sourceURL := "file://database/migrations"
	databaseURL := fmt.Sprintf("mysql://%s:%s@tcp(%s:%s)/%s",
		cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DBName)

	m, err := migrate.New(sourceURL, databaseURL)
	if err != nil {
		panic(err)
	}

	if err = m.Up(); err != nil && err != migrate.ErrNoChange {
		panic(err)
	}
}
