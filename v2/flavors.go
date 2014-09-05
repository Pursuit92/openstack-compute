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

func (cc *ComputeClient) FlavorDetails(flav *Flavor) (*Flavor, error) {
	resp := make(map[string]*Flavor)
	err := cc.AuthedReq("GET", cc.Endpoint.PublicUrl+"/flavors/"+flav.Id, nil, &resp)
	if err != nil {
		return nil, err
	}
	return resp["flavor"], nil
}

