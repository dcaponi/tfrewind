package parsables

// OneloginAppsImportable is the concept that holds the method for determining the
// shape and data that should be collected from tfstate to produce a HCL entry
type OneloginAppsParsable struct{}

// HCLShape returns the address of a new struct for the parser to use when manipulating
// the json from the tfstate. It gets returned as an interface to make it "generic"
func (i OneloginAppsParsable) HCLShape() interface{} {
	return &AppData{}
}

// AppData represents how fields in tfstate should be represented in transition to HCL
type AppData struct {
	AllowAssumedSignin *bool                `json:"allow_assumed_signin,omitempty"`
	ConnectorID        *int32               `json:"connector_id,omitempty"`
	Description        *string              `json:"description,omitempty"`
	Name               *string              `json:"name,omitempty"`
	Notes              *string              `json:"notes,omitempty"`
	Visible            *bool                `json:"visible,omitempty"`
	Configuration      AppConfigurationData `json:"configuration,omitempty"`
	Provisioning       AppProvisioningData  `json:"provisioning,omitempty"`
	Parameters         []AppParametersData  `json:"parameters,omitempty"`
	Rules              []AppRuleData        `json:"rules,omitempty"`
}

// AppProvisioningData represents how fields in tfstate should be represented in transition to HCL
type AppProvisioningData struct {
	Enabled *bool `json:"enabled,omitempty"`
}

// AppConfigurationData represents how fields in tfstate should be represented in transition to HCL
type AppConfigurationData struct {
	RedirectURI                   *string `json:"redirect_uri,omitempty"`
	RefreshTokenExpirationMinutes *string `json:"refresh_token_expiration_minutes,omitempty"`
	LoginURL                      *string `json:"login_url,omitempty"`
	OidcApplicationType           *string `json:"oidc_application_type,omitempty"`
	TokenEndpointAuthMethod       *string `json:"token_endpoint_auth_method,omitempty"`
	AccessTokenExpirationMinutes  *string `json:"access_token_expiration_minutes,omitempty"`
	ProviderArn                   *string `json:"provider_arn,omitempty"`
	SignatureAlgorithm            *string `json:"signature_algorithm,omitempty"`
}

// AppParametersData represents how fields in tfstate should be represented in transition to HCL
type AppParametersData struct {
	ID                        *int32  `json:"id,omitempty"`
	Label                     *string `json:"label,omitempty"`
	UserAttributeMappings     *string `json:"user_attribute_mappings,omitempty"`
	UserAttributeMacros       *string `json:"user_attribute_macros,omitempty"`
	AttributesTransformations *string `json:"attributes_transformations,omitempty"`
	SkipIfBlank               *bool   `json:"skip_if_blank,omitempty"`
	Values                    *string `json:"values,omitempty,omitempty"`
	DefaultValues             *string `json:"default_values,omitempty"`
	ParamKeyName              *string `json:"param_key_name,omitempty"`
	ProvisionedEntitlements   *bool   `json:"provisioned_entitlements,omitempty"`
	SafeEntitlementsEnabled   *bool   `json:"safe_entitlements_enabled,omitempty"`
	IncludeInSamlAssertion    *bool   `json:"include_in_saml_assertion,omitempty"`
}

// AppRuleData represents how fields in tfstate should be represented in transition to HCL
type AppRuleData struct {
	Name       *string                 `json:"name,omitempty"`
	Match      *string                 `json:"match,omitempty"`
	Enabled    *bool                   `json:"enabled,omitempty"`
	Conditions []AppRuleConditionsData `json:"conditions,omitempty"`
	Actions    []AppRuleActionsData    `json:"actions,omitempty"`
}

// AppRuleActionsData represents how fields in tfstate should be represented in transition to HCL
type AppRuleActionsData struct {
	Action     *string  `json:"action,omitempty"`
	Value      []string `json:"value,omitempty"`
	Expression *string  `json:"expression,omitempty"`
}

// AppRuleConditionsData represents how fields in tfstate should be represented in transition to HCL
type AppRuleConditionsData struct {
	Source   *string `json:"source,omitempty"`
	Operator *string `json:"operator,omitempty"`
	Value    *string `json:"value,omitempty"`
}
