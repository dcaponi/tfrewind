// Package parsables parsable_factory.go
// This module creates a list of parsable instances so each can be instantiated once and shared among other callers.
//
// Adding parsables
// Add a case to the switch statement where the key is the name of the resource as it should be represented in terraform
// then call the requisite client method if needed. Be sure the client exists in the clients package to avoid comiler errors.
// Finally, add the parsable to the map of parsables so it can be fetched by referencing the terraform name via the terraform naming convention
package parsables

import "log"

// ParsableList is the list of created parsables referenced by a map where the key is the name used to identify it in terraform
type ParsableList struct {
	parsables map[string]Parsable
}

// New allocates memory for the list of parsables and returns its address
func New() *ParsableList {
	imf := ParsableList{}
	imf.parsables = map[string]Parsable{}
	return &imf
}

// GetParsable will instantiate if not created then return the requested parsable from the list
func (imf *ParsableList) GetParsable(parsableType string) Parsable {
	if imf.parsables[parsableType] == nil {
		switch parsableType {
		case "aws_iam_user":
			imf.parsables[parsableType] = &AWSIamUserParsable{}
		case "onelogin_users":
			imf.parsables[parsableType] = &OneloginUsersParsable{}
		case "onelogin_apps", "onelogin_saml_apps", "onelogin_oidc_apps":
			imf.parsables[parsableType] = &OneloginAppsParsable{}
		case "aws_instance":
			imf.parsables[parsableType] = &AWSInstanceParsable{}
		case "aws_security_group":
			imf.parsables[parsableType] = &AWSSecurityGroupParsable{}
		case "aws_subnet":
			imf.parsables[parsableType] = &AWSSubnetParsable{}
		case "aws_vpc":
			imf.parsables[parsableType] = &AWSVPCParsable{}
		// case "provider_resourceName":
		//   imf.parsables[parsableType] = &ProviderResourceName{}
		default:
			log.Printf("The parsable %s is not configured\n", parsableType)
		}
	}
	return imf.parsables[parsableType]
}
