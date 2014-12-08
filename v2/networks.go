package nova

import "errors"

func (cc *ComputeClient) Networks() ([]*Network, error) {
	resp := make(map[string][]*Network)
	err := cc.AuthedReq("GET", cc.Endpoint.PublicUrl+"/os-networks", nil, &resp)
	if err != nil {
		return nil, err
	}
	return resp["networks"], nil

}

func (cc *ComputeClient) FloatingIPPools() ([]*FloatingIPPool, error) {
	resp := make(map[string][]*FloatingIPPool)
	err := cc.AuthedReq("GET", cc.Endpoint.PublicUrl+"/os-floating-ip-pools", nil, &resp)
	if err != nil {
		return nil, err
	}
	return resp["floating_ip_pools"], nil
}

func (cc *ComputeClient) FloatingIPs() ([]*FloatingIP, error) {
	resp := make(map[string][]*FloatingIP)
	err := cc.AuthedReq("GET", cc.Endpoint.PublicUrl+"/os-floating-ips", nil, &resp)
	if err != nil {
		return nil, err
	}
	return resp["floating_ips"], nil
}

func (cc *ComputeClient) CreateFloatingIP(pool string) (*FloatingIP, error) {
	resp := make(map[string]*FloatingIP)
	req := map[string]string{"pool": pool}
	err := cc.AuthedReq("POST", cc.Endpoint.PublicUrl+"/os-floating-ips", req, &resp)
	if err != nil {
		return nil, err
	}
	return resp["floating_ip"], nil
}

func (cc *ComputeClient) GetFloatingIP(ip string) (*FloatingIP, error) {
	ips, err := cc.FloatingIPs()
	if err != nil {
		return nil, err
	}

	for _, v := range ips {
		if v.IP == ip {
			return v, nil
		}
	}

	return nil, errors.New("IP not found")
}

func (cc *ComputeClient) DeleteFloatingIP(ip string) error {
	addr, err := cc.GetFloatingIP(ip)
	if err != nil {
		return err
	}

	err = cc.AuthedReq("DELETE", cc.Endpoint.PublicUrl+"/os-floating-ips/"+addr.Id, nil, nil)
	return err
}

func (cc *ComputeClient) AddFloatingIP(server, ip string) error {
	req := map[string]map[string]string{"addFloatingIp": map[string]string{"address": ip}}
	srv, err := cc.ServerByName(server)
	if err != nil {
		return err
	}

	err = cc.AuthedReq("POST", cc.Endpoint.PublicUrl+"/servers/"+srv.Id+"/action", req, nil)
	return err
}

func (cc *ComputeClient) RemoveFloatingIP(server, ip string) error {
	req := map[string]map[string]string{"removeFloatingIp": map[string]string{"address": ip}}
	srv, err := cc.ServerByName(server)
	if err != nil {
		return err
	}
	err = cc.AuthedReq("POST", cc.Endpoint.PublicUrl+"/servers/"+srv.Id+"/action", req, nil)
	return err
}
