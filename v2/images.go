package compute

func (cc *ComputeClient) Images() ([]*Image, error) {
	resp := make(map[string][]*Image)
	err := cc.AuthedReq("GET", cc.Endpoint.PublicUrl+"/images", nil, &resp)
	if err != nil {
		return nil, err
	}
	return resp["images"], nil
}

func (cc *ComputeClient) ImagesDetail() ([]*Image, error) {
	resp := make(map[string][]*Image)
	err := cc.AuthedReq("GET", cc.Endpoint.PublicUrl+"/images/detail", nil, &resp)
	if err != nil {
		return nil, err
	}
	return resp["images"], nil
}

func (cc *ComputeClient) ImageDetails(img *Image) (*Image, error) {
	resp := make(map[string]*Image)
	err := cc.AuthedReq("GET", cc.Endpoint.PublicUrl+"/images/"+img.Id, nil, &resp)
	if err != nil {
		return nil, err
	}
	return resp["image"], nil
}

func (cc *ComputeClient) DeleteImage(img *Image) error {
	err := cc.AuthedReq("DELETE", cc.Endpoint.PublicUrl+"/images/"+img.Id, nil, nil)
	if err != nil {
		return err
	}
	return nil
}

func (cc *ComputeClient) GetImageMeta(img *Image) (map[string]string, error) {
	resp := make(map[string]map[string]string)
	err := cc.AuthedReq("GET", cc.Endpoint.PublicUrl+"/images/"+img.Id+"/metadata", nil, &resp)
	if err != nil {
		return nil, err
	}
	return resp["metadata"], nil
}

func (cc *ComputeClient) SetImageMeta(img *Image, meta map[string]string) (map[string]string, error) {
	resp := make(map[string]map[string]string)
	req := make(map[string]map[string]string)
	req["metadata"] = meta
	err := cc.AuthedReq("GET", cc.Endpoint.PublicUrl+"/images/"+img.Id+"/metadata", req, &resp)
	if err != nil {
		return nil, err
	}
	return resp["metadata"], nil
}
