package clients1

import (
	"reflect"

	data1 "github.com/pip-services-samples/pip-services-beacons-go/data/version1"
	cdata "github.com/pip-services3-go/pip-services3-commons-go/data"
	cclients "github.com/pip-services3-go/pip-services3-grpc-go/clients"
)

type BeaconsCommandableGrpcClientV1 struct {
	*cclients.CommandableGrpcClient
	beaconV1DataPageType reflect.Type
	beaconV1Type         reflect.Type
	geoPointV1Type       reflect.Type
}

func NewBeaconsCommandableGrpcClientV1() *BeaconsCommandableGrpcClientV1 {
	c := &BeaconsCommandableGrpcClientV1{
		CommandableGrpcClient: cclients.NewCommandableGrpcClient("v1.beacons"),
		beaconV1DataPageType:  reflect.TypeOf(&data1.BeaconV1DataPage{}),
		beaconV1Type:          reflect.TypeOf(&data1.BeaconV1{}),
		geoPointV1Type:        reflect.TypeOf(&data1.GeoPointV1{}),
	}
	return c
}

func (c *BeaconsCommandableGrpcClientV1) GetBeacons(
	correlationId string, filter *cdata.FilterParams,
	paging *cdata.PagingParams) (*data1.BeaconV1DataPage, error) {

	params := cdata.NewAnyValueMapFromTuples(
		"filter", filter,
		"paging", paging,
	)

	res, err := c.CallCommand(c.beaconV1DataPageType, "get_beacons", correlationId, params)
	if err != nil {
		return nil, err
	}

	result, _ := res.(*data1.BeaconV1DataPage)
	return result, nil
}

func (c *BeaconsCommandableGrpcClientV1) GetBeaconById(
	correlationId string, beaconId string) (*data1.BeaconV1, error) {
	params := cdata.NewAnyValueMapFromTuples(
		"beacon_id", beaconId,
	)

	res, err := c.CallCommand(c.beaconV1Type, "get_beacon_by_id", correlationId, params)
	if err != nil {
		return nil, err
	}

	result, _ := res.(*data1.BeaconV1)
	return result, nil
}

func (c *BeaconsCommandableGrpcClientV1) GetBeaconByUdi(
	correlationId string, udi string) (*data1.BeaconV1, error) {
	params := cdata.NewAnyValueMapFromTuples(
		"udi", udi,
	)

	res, err := c.CallCommand(c.beaconV1Type, "get_beacon_by_udi", correlationId, params)
	if err != nil {
		return nil, err
	}

	result, _ := res.(*data1.BeaconV1)
	return result, nil
}

func (c *BeaconsCommandableGrpcClientV1) CalculatePosition(
	correlationId string, siteId string, udis []string) (*data1.GeoPointV1, error) {
	params := cdata.NewAnyValueMapFromTuples(
		"site_id", siteId,
		"udis", udis,
	)

	res, err := c.CallCommand(c.geoPointV1Type, "calculate_position", correlationId, params)
	if err != nil {
		return nil, err
	}

	result, _ := res.(*data1.GeoPointV1)
	return result, nil
}

func (c *BeaconsCommandableGrpcClientV1) CreateBeacon(
	correlationId string, beacon *data1.BeaconV1) (*data1.BeaconV1, error) {
	params := cdata.NewAnyValueMapFromTuples(
		"beacon", beacon,
	)

	res, err := c.CallCommand(c.beaconV1Type, "create_beacon", correlationId, params)
	if err != nil {
		return nil, err
	}

	result, _ := res.(*data1.BeaconV1)
	return result, nil
}

func (c *BeaconsCommandableGrpcClientV1) UpdateBeacon(
	correlationId string, beacon *data1.BeaconV1) (*data1.BeaconV1, error) {
	params := cdata.NewAnyValueMapFromTuples(
		"beacon", beacon,
	)

	res, err := c.CallCommand(c.beaconV1Type, "update_beacon", correlationId, params)
	if err != nil {
		return nil, err
	}

	result, _ := res.(*data1.BeaconV1)
	return result, nil
}

func (c *BeaconsCommandableGrpcClientV1) DeleteBeaconById(
	correlationId string, beaconId string) (*data1.BeaconV1, error) {
	params := cdata.NewAnyValueMapFromTuples(
		"beacon_id", beaconId,
	)

	res, err := c.CallCommand(c.beaconV1Type, "delete_beacon_by_id", correlationId, params)
	if err != nil {
		return nil, err
	}

	result, _ := res.(*data1.BeaconV1)
	return result, nil
}
