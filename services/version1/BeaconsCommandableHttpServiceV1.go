package services1

import (
	cref "github.com/pip-services3-go/pip-services3-commons-go/refer"
	cservices "github.com/pip-services3-go/pip-services3-rpc-go/services"
)

type BeaconsCommandableHttpServiceV1 struct {
	*cservices.CommandableHttpService
}

func NewBeaconsCommandableHttpServiceV1() *BeaconsCommandableHttpServiceV1 {
	c := &BeaconsCommandableHttpServiceV1{
		CommandableHttpService: cservices.NewCommandableHttpService("v1/beacons"),
	}
	c.CommandableHttpService.IRegisterable = c
	c.DependencyResolver.Put("controller", cref.NewDescriptor("beacons", "controller", "*", "*", "1.0"))
	return c
}

func (c *BeaconsCommandableHttpServiceV1) Register() {
	if !c.SwaggerAuto && c.SwaggerEnable {
		c.RegisterOpenApiSpecFromFile("./swagger/beacons_v1.yaml")
	}
	c.CommandableHttpService.Register()
}
