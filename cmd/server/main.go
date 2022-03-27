package main

import (
	"context"
	"flag"
	"fmt"
	"strings"

	"github.com/lbrictson/auditmon/pkg/auth"

	"github.com/lbrictson/auditmon/pkg/server"

	"github.com/lbrictson/auditmon/pkg/storage"

	"entgo.io/ent/dialect/sql/schema"

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
	err := dbConn.Schema.Create(context.Background(), schema.WithAtlas(true))
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
	userStore := storage.MustNewUserStore(storage.NewUserStoreInput{EntClient: dbConn})
	app := server.MustNewServer(server.NewServerInput{
		EventStorage: storage.MustNewEventStore(storage.NewEventStoreInput{
			EntClient:       dbConn,
			BufferDirectory: config.Durability.BufferLocation,
		}),
		UserStorage:       userStore,
		Port:              config.Server.Port,
		CallBackURL:       config.Server.RootURL,
		MaxSessionSeconds: config.Security.SessionMaxSeconds,
		SessionSecret:     config.Security.SessionSecret,
	})
	err = seedUser(userStore, config.Security.InitialUser, config.Security.InitialPassword, config.Server.RootURL)
	if err != nil {
		panic(fmt.Sprintf("unable to seed initial admin user %v", err))
	}
	app.RunServer()
}

func seedUser(store *storage.UserStore, defaultUser string, defaultPassword string, callbackURL string) error {
	all, err := store.All(context.Background())
	if err != nil {
		return err
	}
	if len(all) == 0 {
		// Seed initial admin user
		mfa := auth.GenerateMFA(fmt.Sprintf("Auditmon | %v", callbackURL), defaultUser)
		_, err := store.Create(context.Background(), storage.CreateUserInput{
			Username:       defaultUser,
			HashedPassword: auth.HashAndSalt(defaultPassword),
			Role:           "Admin",
			MFASecret:      mfa.Secret,
			MFAImage:       mfa.ImageBytes,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
