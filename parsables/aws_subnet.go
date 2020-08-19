package parsables

// AWSSubnetParsable is the concept that holds the method for determining the
// shape and data that should be collected from tfstate to produce a HCL entry
type AWSSubnetParsable struct{}

// HCLShape returns the address of a new struct for the parser to use when manipulating
// the json from the tfstate. It gets returned as an interface to make it "generic"
func (i AWSSubnetParsable) HCLShape() interface{} {
	return &AWSSubnetData{}
}

// AWSSubnetData represents how fields in tfstate should be represented in transition to HCL
type AWSSubnetData struct {
	CIDRBlock                     string `json:"cidr_block,omitempty"`
	VPCID                         string `json:"vpc_id,omitempty"`
	AssignIPV6OnAddressOnCreation bool   `json:"assign_ipv6_address_on_creation,omitempty"`
	AvailabilityZone              string `json:"availability_zone,omitempty"`
	IPV6CIDRBlock                 string `json:"ipv6_cidr_block,omitempty"`
	MapPublicIPOnLaunch           bool   `json:"map_public_ip_on_launch,omitempty"`
	OutpostARN                    string `json:"outpost_arn,omitempty"`
	Tags                          struct {
		Name string `json:"name,omitempty"`
	} `json:"tags,omitempty"`
}
