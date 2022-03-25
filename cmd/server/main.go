package main

import (
	"context"
	"flag"
	"fmt"
	"strings"

	"github.com/lbrictson/auditmon/ent"

	"github.com/lbrictson/auditmon/pkg/configuration"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	log "github.com/sirupsen/logrus"
)

var configFlag = "config/auditmon.yaml"

func main() {
	flag.StringVar(&configFlag, "c", "config/auditmon.yaml", "Location of configuration file")
	config := configuration.MustReadConfig(configFlag)
	dbConn := &ent.Client{}
	// Connect to requested backend
	switch strings.ToLower(config.Database.Backend) {
	case "sqlite3":
		db, err := ent.Open("sqlite3", fmt.Sprintf("file:%v?_fk=1", config.Database.File))
		if err != nil {
			panic(err)
		}
		dbConn = db
	case "postgres":
		db, err := ent.Open("postgres", fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=%v",
			config.Database.Host, config.Database.Port, config.Database.Username, config.Database.Name,
			config.Database.Password, config.Database.SSLMode))
		if err != nil {
			panic(err)
		}
		dbConn = db
	default:
		panic(fmt.Sprintf("%v is not a supported backend. Supported backends are:  sqlite3, postgres", config.Database.Backend))
	}
	// Perform migrations
	err := dbConn.Schema.Create(context.Background())
	if err != nil {
		panic(fmt.Sprintf("failed to perform database migrations, application cannot start %v", err))
	}
	// Configure logger
	log.SetFormatter(&log.JSONFormatter{})
	switch strings.ToLower(config.Logging.Level) {
	case "info":
		log.SetLevel(log.InfoLevel)
	case "warn":
		log.SetLevel(log.WarnLevel)
	case "error":
		log.SetLevel(log.ErrorLevel)
	case "debug":
		log.SetLevel(log.DebugLevel)

	default:
		panic(fmt.Sprintf("%v is not a valid log level, valid log levels are: info,warn,error,debug", config.Logging.Level))
	}
}
