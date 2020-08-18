package parsables

// AWSIamUserParsable is the concept that holds the method for determining the
// shape and data that should be collected from tfstate to produce a HCL entry
type AWSIamUserParsable struct{}

// HCLShape returns the address of a new struct for the parser to use when manipulating
// the json from the tfstate. It gets returned as an interface to make it "generic"
func (i AWSIamUserParsable) HCLShape() interface{} {
	return &AWSUserData{}
}

// AWSUserData represents how fields in tfstate should be represented in transition to HCL
type AWSUserData struct {
	Path string `json:"path,omitempty"`
	Name string `json:"name,omitempty"`
}
