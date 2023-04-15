package main

import (
	"flag"
	"fmt"
	"os"
	"runtime/debug"
	"sync"
	"todotech.henrry.online/internal/database"
	leveledlog "todotech.henrry.online/internal/leveledlog"
	"todotech.henrry.online/internal/version"
)

type config struct {
	baseURL string
	port    int
	apiKey  string
	dsn     string
}

type store struct {
	config config
	db     any // *database.DB
	logger *leveledlog.Logger
	wg     sync.WaitGroup
}

func main() {
	logger := leveledlog.NewLogger(os.Stdout, leveledlog.LevelAll, true)
	err := run(logger)
	if err != nil {
		trace := debug.Stack()
		logger.Fatal(err, trace)
	}
}

func run(logger *leveledlog.Logger) error {
	var cfg config

	// Flags for config

	flag.StringVar(&cfg.baseURL, "base-url", "http://localhost:4444", "base URL for the application")
	flag.IntVar(&cfg.port, "http-port", 4444, "port to listen on for HTTP requests")
	flag.StringVar(&cfg.apiKey, "api-key", "08d98afd-951a-4d21-959f-55ee18f413c1", "API Key for request headers")
	flag.StringVar(&cfg.dsn, "db-dsn", "postgres:Reyshell@localhost:5432/store", "dsn for connect to the database")

	showVersion := flag.Bool("version", false, version.Get())

	flag.Parse()

	if *showVersion {
		fmt.Printf("version: %s\n", version.Get())
		return nil
	}

	db, err := database.New(cfg.dsn)
	if err != nil {
		return err
	}
	defer db.Close()

	st := &store{
		config: cfg,
		db:     db,
		logger: logger,
	}

	return st.serveHTTP()
}
