package test_clients1

import (
	"testing"

	clients1 "github.com/pip-services-samples/pip-services-beacons-go/clients/version1"
	logic "github.com/pip-services-samples/pip-services-beacons-go/logic"
	persist "github.com/pip-services-samples/pip-services-beacons-go/persistence"
	services1 "github.com/pip-services-samples/pip-services-beacons-go/services/version1"
	cconf "github.com/pip-services3-go/pip-services3-commons-go/config"
	cref "github.com/pip-services3-go/pip-services3-commons-go/refer"
)

type beaconsCommandableHttpClientV1Test struct {
	persistence *persist.BeaconsMemoryPersistence
	controller  *logic.BeaconsController
	service     *services1.BeaconsCommandableHttpServiceV1
	client      *clients1.BeaconsCommandableHttpClientV1
	fixture     *BeaconsClientV1Fixture
}

func newBeaconsCommandableHttpClientV1Test() *beaconsCommandableHttpClientV1Test {
	persistence := persist.NewBeaconsMemoryPersistence()
	persistence.Configure(cconf.NewEmptyConfigParams())

	controller := logic.NewBeaconsController()
	controller.Configure(cconf.NewEmptyConfigParams())

	httpConfig := cconf.NewConfigParamsFromTuples(
		"connection.protocol", "http",
		"connection.port", "3000",
		"connection.host", "localhost",
	)

	service := services1.NewBeaconsCommandableHttpServiceV1()
	service.Configure(httpConfig)

	client := clients1.NewBeaconsCommandableHttpClientV1()
	client.Configure(httpConfig)

	references := cref.NewReferencesFromTuples(
		cref.NewDescriptor("pip-services-beacons", "persistence", "memory", "default", "1.0"), persistence,
		cref.NewDescriptor("pip-services-beacons", "controller", "default", "default", "1.0"), controller,
		cref.NewDescriptor("pip-services-beacons", "service", "http", "default", "1.0"), service,
		cref.NewDescriptor("pip-services-beacons", "client", "http", "default", "1.0"), client,
	)
	controller.SetReferences(references)
	service.SetReferences(references)
	client.SetReferences(references)

	fixture := NewBeaconsClientV1Fixture(client)

	return &beaconsCommandableHttpClientV1Test{
		persistence: persistence,
		controller:  controller,
		service:     service,
		client:      client,
		fixture:     fixture,
	}
}

func (c *beaconsCommandableHttpClientV1Test) setup(t *testing.T) {
	err := c.persistence.Open("")
	if err != nil {
		t.Error("Failed to open persistence", err)
	}

	err = c.service.Open("")
	if err != nil {
		t.Error("Failed to open service", err)
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

func (c *beaconsCommandableHttpClientV1Test) teardown(t *testing.T) {
	err := c.client.Close("")
	if err != nil {
		t.Error("Failed to close client", err)
	}

	err = c.service.Close("")
	if err != nil {
		t.Error("Failed to close service", err)
	}

	err = c.persistence.Close("")
	if err != nil {
		t.Error("Failed to close persistence", err)
	}
}

func TestBeaconsCommandableHttpClientV1(t *testing.T) {
	c := newBeaconsCommandableHttpClientV1Test()

	c.setup(t)

	t.Run("CRUD Operations", c.fixture.TestCrudOperations)
	c.teardown(t)

	c.setup(t)

	t.Run("Calculate Positions", c.fixture.TestCalculatePosition)
	c.teardown(t)
}
