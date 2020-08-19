package parsables

// AWSVPCParsable is the concept that holds the method for determining the
// shape and data that should be collected from tfstate to produce a HCL entry
type AWSVPCParsable struct{}

// HCLShape returns the address of a new struct for the parser to use when manipulating
// the json from the tfstate. It gets returned as an interface to make it "generic"
func (i AWSVPCParsable) HCLShape() interface{} {
	return &AWSVPCData{}
}

// AWSVPCData represents how fields in tfstate should be represented in transition to HCL
type AWSVPCData struct {
	CIDRBlock                 string `json:"cidr_block,omitempty"`
	EnableClasslink           bool   `json:"enable_classiclink,omitempty"`
	EnableClasslinkDNSSupport bool   `json:"enable_classiclink_dns_support,omitempty"`
	EnableDNSHostnames        bool   `json:"enable_dns_hostnames,omitempty"`
	EnableDNSSupport          bool   `json:"enable_dns_support,omitempty"`
	InstanceTenancy           string `json:"instance_tenancy,omitempty"`
	Tags                      struct {
		Name string `json:"name,omitempty"`
	} `json:"tags,omitempty"`
}
