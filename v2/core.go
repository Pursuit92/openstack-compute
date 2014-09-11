package compute

import (
	"errors"
	keystone "github.com/Pursuit92/openstack-identity/v2_0"
)

var (
	ErrNoComputeSvc = errors.New("No compute service in catalog.")
)

type ComputeClient struct {
	*keystone.IdentityClient
	Endpoint keystone.Endpoint
}

func NewClient(authUrl string) (*ComputeClient, error) {
	cc := &ComputeClient{}
	var err error
	cc.IdentityClient, err = keystone.NewClient(authUrl)
	return cc, err
}

func (cc *ComputeClient) Authenticate() error {
	err := cc.IdentityClient.Authenticate()
	if err != nil {
		return err
	}
	tenants, err := cc.Tenants()
	if err != nil {
		return err
	}
	if len(tenants) > 0 {
		cc.TenantId(tenants[0].Id)
		for _,v := range tenants {
			if v.Name == cc.Username() {
				cc.TenantId(v.Id)
			}
		}
	}
	cc.IdentityClient.Authenticate()

	found := false
	for _, v := range cc.Access.ServiceCatalog {
		if v.Type == "compute" {
			if len(v.Endpoints) >= 1 {
				cc.Endpoint = v.Endpoints[0]
				found = true
				break
			}
		}
	}
	if !found {
		return ErrNoComputeSvc
	}
	return nil
}
