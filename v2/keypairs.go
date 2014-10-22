package nova

func (cc *ComputeClient) Keypairs() ([]*Keypair, error) {
	resp := make(map[string][]map[string]*Keypair) // lol wtf the keypair response is weird
	err := cc.AuthedReq("GET",cc.Endpoint.PublicUrl+"/os-keypairs",nil,&resp)
	if err != nil {
		return nil,err
	}
	pairs := resp["keypairs"]
	ret := make([]*Keypair,len(pairs))
	for i,v := range pairs {
		ret[i] = v["keypair"]
	}
	return ret, nil
}

func (cc *ComputeClient) NewKeypair(keyPair *Keypair) error {
	resp := make(map[string]*Keypair)
	err := cc.AuthedReq("POST",cc.Endpoint.PublicUrl+"/os-keypairs",map[string]*Keypair{"keypair":keyPair},&resp)
	if err != nil {
		return err
	}
	*keyPair = *resp["keypair"]
	return nil
}

func (cc *ComputeClient) DeleteKeypair(keyPair *Keypair) error {
	return cc.AuthedReq("DELETE",cc.Endpoint.PublicUrl+"/os-keypairs/"+keyPair.Name,nil,nil)
}


func (cc *ComputeClient) GetKeypair(name string) (*Keypair,error) {
	resp := make(map[string]*Keypair)
	err := cc.AuthedReq("GET",cc.Endpoint.PublicUrl+"/os-keypairs/"+name,nil,&resp)
	if err != nil {
		return nil,err
	}
	return resp["keypair"], nil
}

