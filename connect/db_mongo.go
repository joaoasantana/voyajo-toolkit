package connect

import (
	"fmt"
	"github.com/joaoasantana/voyajo-toolkit/configs"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// NewMongoConnection establishes a new connection to a MongoDB database.
// It takes a Database configuration as input and returns a connected *mongo.Client instance or an error.
// The connection URI is constructed using the provided configuration details.
// Parameters:
//   - dbConfig: The database configuration containing host, port, username, and password.
//
// Returns:
//   - *mongo.Client: A MongoDB client instance if the connection is successful.
//   - error: An error if the connection fails.
func NewMongoConnection(dbConfig *configs.Database) (*mongo.Client, error) {
	uri := fmt.Sprintf(
		"mongodb://%s:%s@%s:%s/?authSource=admin",
		dbConfig.Username, dbConfig.Password, dbConfig.Host, dbConfig.Port,
	)

	clientOptions := options.Client().ApplyURI(uri)

	return mongo.Connect(clientOptions)
}
