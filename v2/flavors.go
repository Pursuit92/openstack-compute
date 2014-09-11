package compute

func (cc *ComputeClient) Flavors() ([]*Flavor, error) {
	resp := make(map[string][]*Flavor)
	err := cc.AuthedReq("GET", cc.Endpoint.PublicUrl+"/flavors", nil, &resp)
	if err != nil {
		return nil, err
	}
	return resp["flavors"], nil
}

func (cc *ComputeClient) FlavorsDetail() ([]*Flavor, error) {
	resp := make(map[string][]*Flavor)
	err := cc.AuthedReq("GET", cc.Endpoint.PublicUrl+"/flavors/detail", nil, &resp)
	if err != nil {
		return nil, err
	}
	return resp["flavors"], nil
}

func (cc *ComputeClient) FlavorDetails(flavId string) (*Flavor, error) {
	resp := make(map[string]*Flavor)
	err := cc.AuthedReq("GET", cc.Endpoint.PublicUrl+"/flavors/"+flavId, nil, &resp)
	if err != nil {
		return nil, err
	}
	return resp["flavor"], nil
}

