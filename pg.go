package pg

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	"github.com/loadimpact/k6/js/modules"
)

// Register the extension on module initialization, available to
// import from JS as "k6/x/PG".
func init() {
	modules.Register("k6/x/pg", new(PG))
}

// PG is the k6 extension for a PG connection.
type PG struct{}

// NewClient creates a new PG client with the provided options.
func (*PG) NewClient(conninfo string) (*sql.DB, error) {
	print("Opening connection")
	//defer db.Close()

	return sql.Open("postgres", conninfo)
}

// Insert the given key with the given value and expiration time.
func (*PG) Insert(client *sql.DB, query string) {
	_, err := client.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}

// Get returns the value for the given key.
func (*PG) Get(client *sql.DB, query string) (*sql.Rows, error) {
	return client.Query(query)
}

// Close closes the db connection
func (*PG) Close(client *sql.DB) {
	client.Close()
}

