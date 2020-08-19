package parsables

// AWSInstanceParsable is the concept that holds the method for determining the
// shape and data that should be collected from tfstate to produce a HCL entry
type AWSInstanceParsable struct{}

// HCLShape returns the address of a new struct for the parser to use when manipulating
// the json from the tfstate. It gets returned as an interface to make it "generic"
func (i AWSInstanceParsable) HCLShape() interface{} {
	return &AWSInstanceData{}
}

// AWSInstanceData represents how fields in tfstate should be represented in transition to HCL
type AWSInstanceData struct {
	InstanceType      string `json:"instance_type,omitempty"`
	AMI               string `json:"ami,omitempty"`
	AvailabilityZone  string `json:"availability_zone,omitempty"`
	CPUCoreCount      int    `json:"cpu_core_count,omitempty"`
	CPUThreadsPerCore int    `json:"cpu_threads_per_core,omitempty"`
	HostID            string `json:"host_id,omitempty"`
	PlacementGroup    string `json:"placement_group,omitempty"`
	Tenancy           string `json:"tenancy,omitempty"`
}
