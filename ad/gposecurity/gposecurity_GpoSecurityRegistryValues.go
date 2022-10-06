package gposecurity


type GpoSecurityRegistryValues struct {
	// Fully qualified name of the key (https://docs.microsoft.com/en-us/openspecs/windows_protocols/ms-rrp/97587de7-3524-4291-8527-39517110c0eb).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ad/r/gpo_security#key_name GpoSecurity#key_name}
	KeyName *string `field:"required" json:"keyName" yaml:"keyName"`
	// The value of the key, matching the type set in `value_type`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ad/r/gpo_security#value GpoSecurity#value}
	Value *string `field:"required" json:"value" yaml:"value"`
	// Data type of the key's value. 1: String, 2: Expand String, 3: Binary, 4: DWORD, 5: MULTI_SZ.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ad/r/gpo_security#value_type GpoSecurity#value_type}
	ValueType *string `field:"required" json:"valueType" yaml:"valueType"`
}

