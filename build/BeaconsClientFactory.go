package build

import (
	clients1 "github.com/pip-services-samples/pip-services-beacons-go/clients/version1"
	cref "github.com/pip-services3-go/pip-services3-commons-go/refer"
	cbuild "github.com/pip-services3-go/pip-services3-components-go/build"
)

type BeaconsClientFactory struct {
	cbuild.Factory
}

func NewBeaconsClientFactory() *BeaconsClientFactory {
	c := &BeaconsClientFactory{
		Factory: *cbuild.NewFactory(),
	}

	nullClientDescriptor := cref.NewDescriptor("beacons", "client", "null", "*", "1.0")
	directClientDescriptor := cref.NewDescriptor("beacons", "client", "direct", "*", "1.0")
	cmdHttpClientDescriptor := cref.NewDescriptor("beacons", "client", "commandable-http", "*", "1.0")
	cmdGrpcClientDescriptor := cref.NewDescriptor("beacons", "client", "commandable-grpc", "*", "1.0")

	httpClientDescriptor := cref.NewDescriptor("beacons", "client", "http", "*", "1.0")
	grpcClientDescriptor := cref.NewDescriptor("beacons", "client", "grpc", "*", "1.0")

	c.RegisterType(nullClientDescriptor, clients1.NewBeaconsNullClientV1)
	c.RegisterType(directClientDescriptor, clients1.NewBeaconsDirectClientV1)
	c.RegisterType(cmdHttpClientDescriptor, clients1.NewBeaconsCommandableHttpClientV1)
	c.RegisterType(cmdGrpcClientDescriptor, clients1.NewBeaconsCommandableGrpcClientV1)

	c.RegisterType(httpClientDescriptor, clients1.NewBeaconsRestClientV1)
	c.RegisterType(grpcClientDescriptor, clients1.NewBeaconsGrpcClientV1)

	return c
}
