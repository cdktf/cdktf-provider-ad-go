// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gposecurity


type GpoSecurityPasswordPolicies struct {
	// Store password with reversible encryption (0-2^16).
	//
	// The password will not be stored with reversible encryption if the value is set to 0. Reversible encryption will be used in any other case.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/ad/0.4.4/docs/resources/gpo_security#clear_text_password GpoSecurity#clear_text_password}
	ClearTextPassword *string `field:"optional" json:"clearTextPassword" yaml:"clearTextPassword"`
	// Number of days before password expires (-1-999). If set to -1, it means the password never expires.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/ad/0.4.4/docs/resources/gpo_security#maximum_password_age GpoSecurity#maximum_password_age}
	MaximumPasswordAge *string `field:"optional" json:"maximumPasswordAge" yaml:"maximumPasswordAge"`
	// Number of days a password must be used before changing it (0-999).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/ad/0.4.4/docs/resources/gpo_security#minimum_password_age GpoSecurity#minimum_password_age}
	MinimumPasswordAge *string `field:"optional" json:"minimumPasswordAge" yaml:"minimumPasswordAge"`
	// Minimum number of characters used in a password (0-2^16). If set to 0, it means no password is required.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/ad/0.4.4/docs/resources/gpo_security#minimum_password_length GpoSecurity#minimum_password_length}
	MinimumPasswordLength *string `field:"optional" json:"minimumPasswordLength" yaml:"minimumPasswordLength"`
	// Password must meet complexity requirements (0-2^16).
	//
	// If set to 0, then requirements do not apply, any other value means requirements are applied
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/ad/0.4.4/docs/resources/gpo_security#password_complexity GpoSecurity#password_complexity}
	PasswordComplexity *string `field:"optional" json:"passwordComplexity" yaml:"passwordComplexity"`
	// The number of unique new passwords that are required before an old password can be reused in association with a user account (0-2^16).
	//
	// A value of 0 indicates that the password history is disabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/ad/0.4.4/docs/resources/gpo_security#password_history_size GpoSecurity#password_history_size}
	PasswordHistorySize *string `field:"optional" json:"passwordHistorySize" yaml:"passwordHistorySize"`
}

