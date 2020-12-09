package test_clients1

import (
	"testing"

	clients1 "github.com/pip-services-samples/pip-services-beacons-go/clients/version1"
	logic "github.com/pip-services-samples/pip-services-beacons-go/logic"
	persist "github.com/pip-services-samples/pip-services-beacons-go/persistence"
	cconf "github.com/pip-services3-go/pip-services3-commons-go/config"
	cref "github.com/pip-services3-go/pip-services3-commons-go/refer"
)

type beaconsDirectClientV1Test struct {
	persistence *persist.BeaconsMemoryPersistence
	controller  *logic.BeaconsController
	client      *clients1.BeaconsDirectClientV1
	fixture     *BeaconsClientV1Fixture
}

func newBeaconsDirectClientV1Test() *beaconsDirectClientV1Test {
	persistence := persist.NewBeaconsMemoryPersistence()
	persistence.Configure(cconf.NewEmptyConfigParams())

	controller := logic.NewBeaconsController()
	controller.Configure(cconf.NewEmptyConfigParams())

	client := clients1.NewBeaconsDirectClientV1()
	client.Configure(cconf.NewEmptyConfigParams())

	references := cref.NewReferencesFromTuples(
		cref.NewDescriptor("pip-services-beacons", "persistence", "memory", "default", "1.0"), persistence,
		cref.NewDescriptor("pip-services-beacons", "controller", "default", "default", "1.0"), controller,
		cref.NewDescriptor("pip-services-beacons", "client", "direct", "default", "1.0"), client,
	)
	controller.SetReferences(references)
	client.SetReferences(references)

	fixture := NewBeaconsClientV1Fixture(client)

	return &beaconsDirectClientV1Test{
		persistence: persistence,
		controller:  controller,
		client:      client,
		fixture:     fixture,
	}
}

func (c *beaconsDirectClientV1Test) setup(t *testing.T) {
	err := c.persistence.Open("")
	if err != nil {
		t.Error("Failed to open persistence", err)
	}

	err = c.client.Open("")
	if err != nil {
		t.Error("Failed to open client", err)
	}

	err = c.persistence.Clear("")
	if err != nil {
		t.Error("Failed to clear persistence", err)
	}
}

func (c *beaconsDirectClientV1Test) teardown(t *testing.T) {
	err := c.client.Close("")
	if err != nil {
		t.Error("Failed to close client", err)
	}

	err = c.persistence.Close("")
	if err != nil {
		t.Error("Failed to close persistence", err)
	}
}

func TestBeaconsDirectClientV1(t *testing.T) {
	c := newBeaconsDirectClientV1Test()

	c.setup(t)
	t.Run("CRUD Operations", c.fixture.TestCrudOperations)
	c.teardown(t)

	c.setup(t)
	t.Run("Calculate Positions", c.fixture.TestCalculatePosition)
	c.teardown(t)
}
