package repositories_test

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"
	"time"

	"github.com/antonve/portfolio-api/infra"
	"github.com/jmoiron/sqlx"

	txdb "github.com/DATA-DOG/go-txdb"
	"github.com/DavidHuie/gomigrate"
	_ "github.com/lib/pq"
)

func TestMain(m *testing.M) {
	dbURL := os.Getenv("TESTING_DATABASE_URL")
	// Must be called pgx so the sqlx mapper uses the correct notation
	txdb.Register("pgx", "postgres", dbURL)

	db, err := infra.NewRDB(dbURL, 1, 1)
	if err != nil {
		panic(fmt.Sprintf("could not connect to testing DB: %s", err))
	}

	migrator, _ := gomigrate.NewMigratorWithLogger(
		db.DB,
		gomigrate.Postgres{},
		"./../migrations",
		log.New(ioutil.Discard, "", log.LstdFlags),
	)

	err = migrator.Migrate()
	if err != nil {
		panic(fmt.Sprintf("could not migrate testing DB: %s", err))
	}

	code := m.Run()
	defer os.Exit(code)
}

func setupTestingSuite(t *testing.T) (*infra.RDB, func() error) {
	t.Parallel()

	db, cleanup := prepareDB(t)
	return db, cleanup
}

func prepareDB(t *testing.T) (*infra.RDB, func() error) {
	cName := fmt.Sprintf("connection_%d", time.Now().UnixNano())
	db, err := sqlx.Open("pgx", cName)

	if err != nil {
		t.Fatalf("open pgx connection: %s", err)
	}

	return db, db.Close
}
