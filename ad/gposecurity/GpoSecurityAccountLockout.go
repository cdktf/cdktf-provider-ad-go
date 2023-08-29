// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gposecurity


type GpoSecurityAccountLockout struct {
	// Disconnect SMB sessions when logon hours expire.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/ad/0.4.4/docs/resources/gpo_security#force_logoff_when_hour_expire GpoSecurity#force_logoff_when_hour_expire}
	ForceLogoffWhenHourExpire *string `field:"optional" json:"forceLogoffWhenHourExpire" yaml:"forceLogoffWhenHourExpire"`
	// Number of failed logon attempts until a account is locked.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/ad/0.4.4/docs/resources/gpo_security#lockout_bad_count GpoSecurity#lockout_bad_count}
	LockoutBadCount *string `field:"optional" json:"lockoutBadCount" yaml:"lockoutBadCount"`
	// Number of minutes a locked out account must remain locked out.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/ad/0.4.4/docs/resources/gpo_security#lockout_duration GpoSecurity#lockout_duration}
	LockoutDuration *string `field:"optional" json:"lockoutDuration" yaml:"lockoutDuration"`
	// Number of minutes a account will remain locked after a failed logon attempt.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/ad/0.4.4/docs/resources/gpo_security#reset_lockout_count GpoSecurity#reset_lockout_count}
	ResetLockoutCount *string `field:"optional" json:"resetLockoutCount" yaml:"resetLockoutCount"`
}

