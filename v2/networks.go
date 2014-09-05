package compute

func (cc *ComputeClient) Networks() ([]*Network, error) {
	resp := make(map[string][]*Network)
	err := cc.AuthedReq("GET", cc.Endpoint.PublicUrl+"/os-networks", nil, &resp)
	if err != nil {
		return nil, err
	}
	return resp["networks"], nil

}
