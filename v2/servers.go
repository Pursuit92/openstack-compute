package compute

func (cc *ComputeClient) Servers() ([]*Server, error) {
	resp := make(map[string][]*Server)
	err := cc.AuthedReq("GET", cc.Endpoint.PublicUrl+"/servers", nil, &resp)
	if err != nil {
		return nil, err
	}
	return resp["servers"], nil
}

func (cc *ComputeClient) ServersDetail() ([]*Server, error) {
	resp := make(map[string][]*Server)
	err := cc.AuthedReq("GET", cc.Endpoint.PublicUrl+"/servers/detail", nil, &resp)
	if err != nil {
		return nil, err
	}
	return resp["servers"], err
}

func (cc *ComputeClient) Details(srv *Server) (*Server, error) {
	serverId := srv.Id
	resp := make(map[string]*Server)

	err := cc.AuthedReq("GET", cc.Endpoint.PublicUrl+"/servers/"+serverId, nil, &resp)
	if err != nil {
		return nil, err
	}
	return resp["server"], nil
}

func (cc *ComputeClient) Delete(srv *Server) error {
	serverId := srv.Id
	return cc.AuthedReq("DELETE", cc.Endpoint.PublicUrl+"/servers/"+serverId, nil, nil)
}

func (cc *ComputeClient) Update(srv *Server) (*Server, error) {
	resp := make(map[string]*Server)
	err := cc.AuthedReq("PUT", cc.Endpoint.PublicUrl+"/servers/"+srv.Id, srv, &resp)
	if err != nil {
		return nil, err
	}
	return resp["server"], nil
}

func NewServer() *Server {
	return &Server{ServerDetail: &ServerDetail{}, ServerCreate: &ServerCreate{}}
}

func (cc *ComputeClient) create(srv *Server) (*Server, error) {
	req := make(map[string]*Server)
	resp := make(map[string]*Server)
	if srv.TenantId == "" {
		srv.TenantId = cc.Access.Token.Tenant.Id
	}
	req["server"] = srv
	err := cc.AuthedReq("POST", cc.Endpoint.PublicUrl+"/servers", req, &resp)
	if err != nil {
		return nil, err
	}
	return resp["server"], nil
}
