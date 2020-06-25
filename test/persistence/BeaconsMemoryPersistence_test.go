package test_persistence

import (
	"testing"

	bpersist "github.com/pip-services-samples/pip-services-beacons-go/persistence"
	cconf "github.com/pip-services3-go/pip-services3-commons-go/config"
)

func TestBeaconsMemoryPersistence(t *testing.T) {
	var persistence *bpersist.BeaconsMemoryPersistence
	var fixture *BeaconsPersistenceFixture

	persistence = bpersist.NewBeaconsMemoryPersistence()
	persistence.Configure(cconf.NewEmptyConfigParams())
	fixture = NewBeaconsPersistenceFixture(persistence)

	persistence.Open("")

	defer persistence.Close("")

	t.Run("BeaconsMemoryPersistence:CRUD Operations", fixture.TestCrudOperations)
	persistence.Clear("")
	t.Run("BeaconsMemoryPersistence:Get with Filters", fixture.TestGetWithFilters)
}
