package parsables

// AWSSecurityGroupParsable is the concept that holds the method for determining the
// shape and data that should be collected from tfstate to produce a HCL entry
type AWSSecurityGroupParsable struct{}

// HCLShape returns the address of a new struct for the parser to use when manipulating
// the json from the tfstate. It gets returned as an interface to make it "generic"
func (i AWSSecurityGroupParsable) HCLShape() interface{} {
	return &AWSSecurityGroupData{}
}

// AWSSecurityGroupData represents how fields in tfstate should be represented in transition to HCL
type AWSSecurityGroupData struct {
	Description         string `json:"description,omitempty"`
	Name                string `json:"name,omitempt"`
	RevokeRulesOnDelete bool   `json:"revoke_rules_on_delete,omitempt"`
	Tags                struct {
		Name string `json:"name,omitempty"`
	} `json:"tags,omitempt"`
	VPCID   string                      `json:"vpc_id,omitempt"`
	Ingress AWSSecurityGroupIngressData `json:"ingress,omitempty"`
	Egress  AWSSecurityGroupEgressData  `json:"egress,omitempty"`
}

type AWSSecurityGroupIngressData struct {
	CIDRBlocks     []string `json:"cidr_blocks,omitempty"`
	Description    string   `json:"description,omitempty"`
	FromPort       int      `json:"from_port,omitempty"`
	IPV6CIDRBlocks []string `json:"ipv6_cidr_blocks,omitempty"`
	PrefixListIDs  []string `json:"prefix_list_ids,omitempty"`
	Protocol       string   `json:"protocol,omitempty"`
	SecurityGroups []string `json:"security_groups,omitempty"`
	Self           bool     `json:"self,omitempty"`
	ToPort         int      `json:"to_port,omitempty"`
}

type AWSSecurityGroupEgressData struct {
	CIDRBlocks     []string `json:"cidr_blocks,omitempty"`
	Description    string   `json:"description,omitempty"`
	FromPort       int      `json:"from_port,omitempty"`
	IPV6CIDRBlocks []string `json:"ipv6_cidr_blocks,omitempty"`
	PrefixListIDs  []string `json:"prefix_list_ids,omitempty"`
	Protocol       string   `json:"protocol,omitempty"`
	SecurityGroups []string `json:"security_groups,omitempty"`
	Self           bool     `json:"self,omitempty"`
	ToPort         int      `json:"to_port,omitempty"`
}
