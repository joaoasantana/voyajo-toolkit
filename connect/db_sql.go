package connect

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/joaoasantana/voyajo-toolkit/configs"
)

// NewSQLConnection establishes a new connection to a SQL database.
// It takes a Database configuration as input and returns a connected *sqlx.DB instance or an error.
// The connection string is constructed using the provided configuration details.
func NewSQLConnection(dbConfig *configs.Database) (*sqlx.DB, error) {
	uri := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbConfig.Host, dbConfig.Port, dbConfig.Username, dbConfig.Password, dbConfig.Name,
	)

	return sqlx.Connect(dbConfig.Driver, uri)
}
