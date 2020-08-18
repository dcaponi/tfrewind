package parsables

// Parsable is an interface for returning the address of a pre-defined struct
// where the fields help the json package determine which fields on the .tfstate json
// object map to fields in the HCL plan file
type Parsable interface {
	HCLShape() interface{} // dictates what fields on tfstate should be represented in HCL files
}
