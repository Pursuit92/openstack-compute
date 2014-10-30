package nova

import (
	"github.com/Pursuit92/openstack-core"
	glance "github.com/Pursuit92/openstack-image/v2"
	neutron "github.com/Pursuit92/openstack-network/v2_0"
)

type Image glance.Image
type Network neutron.Network

// Fully describes a server
type Server struct {
	Name  string      `json:"name,omitempty"`
	Id    string      `json:"id,omitempty"`
	Links []core.Link `json:"links,omitempty"`
	*ServerDetail
	*ServerCreate
	*CreateResp
}

// fields unique to detailed server information
type ServerDetail struct {
	Updated        string               `json:"updated,omitempty"`
	TenantId       string               `json:"tenant_id,omitempty"`
	Status         string               `json:"status,omitempty"`
	Progress       int                  `json:"progress,omitempty"`
	Image          *Image               `json:"image,omitempty"`
	HostId         string               `json:"host_id,omitempty"`
	Flavor         *Flavor              `json:"flavor,omitempty"`
	KeyName        string               `json:"key_name,omitempty"`
	Created        string               `json:"created,omitempty"`
	Addresses      map[string][]Address `json:"addresses,omitempty"`
	Metadata       map[string]string    `json:"metadata,omitempty"`
	AccessIPv4     string               `json:"accessIPv4,omitempty"`
	AccessIPv6     string               `json:"accessIPv6,omitempty"`
	SecurityGroups []SecurityGroup      `json:"security_groups,omitempty"`
	ConfigDrive    string               `json:"config_drive,omitempty"`
}

type Address struct {
	Addr    string `json:"addr,omitempty"`
	Version int    `json:"version,omitempty"`
}

// fields unique to server creation
type ServerCreate struct {
	UserData         string        `json:"user_data,omitempty"`
	AvailabilityZone string        `json:"availability_zone,omitempty"`
	ImageRef         string        `json:"imageRef,omitempty"`
	FlavorRef        string        `json:"flavorRef,omitempty"`
	Networks         []NetConf     `json:"networks,omitempty"`
	Personality      []Personality `json:"personality,omitempty"`
	NetNames         []string      `json:"-"`
}

type Personality struct {
	Path     string `json:"path,omitempty"`
	Contents string `json:"contents,omitempty"`
}

// fields unique to the server creation response
type CreateResp struct {
	AdminPass string `json:"adminPass,omitempty"`
}

type NetConf struct {
	// necessary if the port isn't specified
	Uuid string `json:"uuid,omitempty"`
	// necessary if the uuid isn't specified
	Port    string `json:"port,omitempty"`
	FixedIp string `json:"fixed_ip,omitempty"`
}

type SecurityGroup struct {
	Name string `json:"name,omitempty"`
}

type Flavor struct {
	Name  string      `json:"name,omitempty"`
	Id    string      `json:"id,omitempty"`
	Links []core.Link `json:"links,omitempty"`
	Disk  int         `json:"disk,omitempty"`
	Ram   int         `json:"ram,omitempty"`
	Vcpus int         `json:"vcpus,omitempty"`
}

type Keypair struct {
	Fingerprint string `json:"fingerprint,omitempty"`
	Name        string `json:"name,omitempty"`
	PublicKey   string `json:"public_key,omitempty"`
}
