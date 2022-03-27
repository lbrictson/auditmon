package main

import (
	"context"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql/schema"

	"github.com/lbrictson/auditmon/ent"
	"github.com/lbrictson/auditmon/pkg/configuration"
	"github.com/lbrictson/auditmon/pkg/storage"

	"github.com/google/uuid"

	"github.com/lbrictson/auditmon/pkg/models"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	c := configuration.MustReadConfig("config/auditmon.yaml")
	dbConn := &ent.Client{}
	// Connect to requested backend
	switch strings.ToLower(c.Database.Backend) {
	case "sqlite3":
		db, err := ent.Open("sqlite3", fmt.Sprintf("file:%v?_fk=1", c.Database.File))
		if err != nil {
			panic(err)
		}
		dbConn = db
	case "postgres":
		db, err := ent.Open("postgres", fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=%v",
			c.Database.Host, c.Database.Port, c.Database.Username, c.Database.Name,
			c.Database.Password, c.Database.SSLMode))
		if err != nil {
			panic(err)
		}
		dbConn = db
	default:
		panic(fmt.Sprintf("%v is not a supported backend. Supported backends are:  sqlite3, postgres", c.Database.Backend))
	}
	store := storage.MustNewEventStore(storage.NewEventStoreInput{
		EntClient:       dbConn,
		BufferDirectory: "test_data/",
	})
	err := dbConn.Schema.Create(context.Background(), schema.WithAtlas(true))
	if err != nil {
		panic(fmt.Sprintf("failed to perform database migrations, application cannot start %v", err))
	}
	seedCount := 500
	seedRounds := 50
	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)
	fmt.Println("Executing seeds")
	for i := 1; i < seedRounds; i++ {
		events := []models.Event{}
		for y := 1; y < seedCount; y++ {
			events = append(events, models.Event{
				EventName:       seedEventNames[r.Intn(len(seedEventNames))],
				EventTime:       time.Now().Add(-1 * time.Duration(r.Intn(1000)) * time.Minute),
				Username:        seedUsers[r.Intn(len(seedUsers))],
				Resource:        seedResources[r.Intn(len(seedResources))],
				EventSource:     "auditmon_seed_events",
				SourceIPAddress: seedIPs[r.Intn(len(seedIPs))],
				RequestID:       uuid.New().String(),
				ReadOnly:        seedBool[r.Intn(len(seedBool))],
			})
		}
		store.Create(context.TODO(), events)
	}
}

var seedBool = []bool{
	true,
	false,
}

var seedUsers = []string{
	"fake_admin",
	"dan",
	"paula",
	"jamie",
	"max",
	"tim",
	"janet",
	"mark",
	"debra",
	"peter",
	"paul",
	"scott",
}

var seedResources = []string{
	"company_laptop",
	"website_a",
	"website_b",
	"self",
	"-",
	"payment",
}

var seedEventNames = []string{
	"create_login",
	"destroy_session",
	"create_payment",
	"read_user_details",
	"unmask_credit_card_number",
	"read_ssn",
	"read_secret_a",
	"publish_page",
	"review_audit",
}

var seedIPs = []string{
	"196.168.1.1",
	"196.168.1.2",
	"196.168.1.3",
	"196.168.1.4",
	"196.168.1.5",
	"196.168.1.6",
}
