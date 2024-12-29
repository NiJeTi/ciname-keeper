package main

import (
	"embed"
	"log/slog"
	"os"

	"github.com/pressly/goose/v3"

	"github.com/nijeti/cinema-keeper/internal/db"
	cfgpkg "github.com/nijeti/cinema-keeper/internal/pkg/config"
)

type config struct {
	DB db.Config `conf:"db"`
}

const (
	codeOk  = 0
	codeErr = 1
)

const migrationsTable = "migrations"

//go:embed *.sql
var migrations embed.FS

func main() {
	code := run()
	os.Exit(code)
}

func run() int {
	// logging
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	logger.Info("starting")

	// config
	cfgpkg.SetLogger(logger)
	cfg, err := cfgpkg.ReadConfig[config]()
	if err != nil {
		logger.Error("failed to read config", "error", err)
		return codeErr
	}

	// db
	dbConn, err := db.Connect(cfg.DB.ConnectionString)
	if err != nil {
		logger.Error("failed to connect to db", "error", err)
		return codeErr
	}
	defer dbConn.Close()

	// migrator
	goose.SetLogger(gooseLogger{logger})
	goose.SetBaseFS(migrations)
	goose.SetTableName(migrationsTable)
	goose.SetSequential(true)

	err = goose.SetDialect(string(goose.DialectPostgres))
	if err != nil {
		logger.Error("failed to set dialect", "error", err)
		return codeErr
	}

	// run
	logger.Info("begin database migration")

	err = goose.Up(dbConn.DB, ".")
	if err != nil {
		logger.Error("failed to migrate database", "error", err)
		return codeErr
	}

	logger.Info("database migration complete")
	return codeOk
}
