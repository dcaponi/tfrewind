package parsables

type OneloginUsersParsable struct{}

func (i OneloginUsersParsable) HCLShape() interface{} {
	return &UserData{}
}

// UserData represents how fields in tfstate should be represented in transition to HCL
type UserData struct {
	Firstname            *string `json:"firstname,omitempty"`
	Lastname             *string `json:"lastname,omitempty"`
	Username             *string `json:"username,omitempty"`
	Email                *string `json:"email,omitempty"`
	DistinguishedName    *string `json:"distinguished_name,omitempty"`
	Samaccountname       *string `json:"samaccountname,omitempty"`
	UserPrincipalName    *string `json:"userprincipalname,omitempty"`
	MemberOf             *string `json:"member_of,omitempty"`
	Phone                *string `json:"phone,omitemepty"`
	Title                *string `json:"title,omitempty"`
	Company              *string `json:"company,omitempty"`
	Department           *string `json:"department,omitempty"`
	Comment              *string `json:"comment,omitempty"`
	State                *int32  `json:"state,omitempty"`
	Status               *int32  `json:"status,omitempty"`
	InvalidLoginAttempts *int32  `json:"invalid_login_attempts,omitempty"`
	GroupID              *int32  `json:"group_id,omitempty"`
	DirectoryID          *int32  `json:"directory_id,omitempty"`
	TrustedIDPID         *int32  `json:"trusted_idp_id,omitempty"`
	ManagerADID          *int32  `json:"manager_ad_id,omitempty"`
	ManagerUserID        *int32  `json:"manager_user_id,omitempty"`
	ExternalID           *int32  `json:"external_id,omitempty"`
}
