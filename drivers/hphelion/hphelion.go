package openstack

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"github.com/codegangsta/cli"
	"github.com/docker/machine/drivers"
	"github.com/docker/machine/log"
	"github.com/docker/machine/ssh"
	"github.com/docker/machine/state"
	"github.com/docker/machine/utils"
)

type Driver struct {
	*drivers.BaseDriver


	AuthUrl          string
	ActiveTimeout    int
	Insecure         bool
	DomainID         string
	DomainName       string
	Username         string
	Password         string
	TenantName       string
	TenantId         string
	Region           string
	AvailabilityZone string
	EndpointType     string
	MachineId        string
	FlavorName       string
	FlavorId         string
	ImageName        string
	ImageId          string
	KeyPairName      string
	NetworkName      string
	NetworkId        string
	SecurityGroups   []string
	FloatingIpPool   string
	FloatingIpPoolId string

	//hp helion specific feature
	DnsName			string

	client           Client
}

func init() {
	drivers.Register("hphelion", &drivers.RegisteredDriver{
		New:			NewDriver, 
		GetCreateFlags: GetCreateFlags,
	})
}

func GetCreateFlags() []cli.Flag {
	return [].cli.Flag{

	}
}