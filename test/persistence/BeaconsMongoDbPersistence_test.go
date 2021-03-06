package test_persistence

import (
	"os"
	"testing"

	bpersist "github.com/pip-services-samples/pip-services-beacons-go/persistence"
	cconf "github.com/pip-services3-go/pip-services3-commons-go/config"
)

func TestBeaconsMongoDbPersistence(t *testing.T) {

	var persistence *bpersist.BeaconsMongoDbPersistence
	var fixture *BeaconsPersistenceFixture

	mongoUri := os.Getenv("MONGO_SERVICE_URI")
	mongoHost := os.Getenv("MONGO_SERVICE_HOST")

	if mongoHost == "" {
		mongoHost = "localhost"
	}
	mongoPort := os.Getenv("MONGO_SERVICE_PORT")
	if mongoPort == "" {
		mongoPort = "27017"
	}

	mongoDatabase := os.Getenv("MONGO_SERVICE_DB")
	if mongoDatabase == "" {
		mongoDatabase = "test"
	}

	// Exit if mongo connection is not set
	if mongoUri == "" && mongoHost == "" {
		return
	}

	persistence = bpersist.NewBeaconsMongoDbPersistence()
	persistence.Configure(cconf.NewConfigParamsFromTuples(
		"connection.uri", mongoUri,
		"connection.host", mongoHost,
		"connection.port", mongoPort,
		"connection.database", mongoDatabase,
	))

	fixture = NewBeaconsPersistenceFixture(persistence)

	opnErr := persistence.Open("")
	if opnErr == nil {
		persistence.Clear("")
	}

	defer persistence.Close("")

	t.Run("BeaconsMongoDbPersistence:CRUD Operations", fixture.TestCrudOperations)
	persistence.Clear("")
	t.Run("BeaconsMongoDbPersistence:Get with Filters", fixture.TestGetWithFilters)

}
