// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gposecurity

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GpoSecurityConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The GUID of the container the security settings belong to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/ad/0.4.4/docs/resources/gpo_security#gpo_container GpoSecurity#gpo_container}
	GpoContainer *string `field:"required" json:"gpoContainer" yaml:"gpoContainer"`
	// account_lockout block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/ad/0.4.4/docs/resources/gpo_security#account_lockout GpoSecurity#account_lockout}
	AccountLockout *GpoSecurityAccountLockout `field:"optional" json:"accountLockout" yaml:"accountLockout"`
	// application_log block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/ad/0.4.4/docs/resources/gpo_security#application_log GpoSecurity#application_log}
	ApplicationLog *GpoSecurityApplicationLog `field:"optional" json:"applicationLog" yaml:"applicationLog"`
	// audit_log block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/ad/0.4.4/docs/resources/gpo_security#audit_log GpoSecurity#audit_log}
	AuditLog *GpoSecurityAuditLog `field:"optional" json:"auditLog" yaml:"auditLog"`
	// event_audit block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/ad/0.4.4/docs/resources/gpo_security#event_audit GpoSecurity#event_audit}
	EventAudit *GpoSecurityEventAudit `field:"optional" json:"eventAudit" yaml:"eventAudit"`
	// filesystem block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/ad/0.4.4/docs/resources/gpo_security#filesystem GpoSecurity#filesystem}
	Filesystem interface{} `field:"optional" json:"filesystem" yaml:"filesystem"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/ad/0.4.4/docs/resources/gpo_security#id GpoSecurity#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// kerberos_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/ad/0.4.4/docs/resources/gpo_security#kerberos_policy GpoSecurity#kerberos_policy}
	KerberosPolicy *GpoSecurityKerberosPolicy `field:"optional" json:"kerberosPolicy" yaml:"kerberosPolicy"`
	// password_policies block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/ad/0.4.4/docs/resources/gpo_security#password_policies GpoSecurity#password_policies}
	PasswordPolicies *GpoSecurityPasswordPolicies `field:"optional" json:"passwordPolicies" yaml:"passwordPolicies"`
	// registry_keys block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/ad/0.4.4/docs/resources/gpo_security#registry_keys GpoSecurity#registry_keys}
	RegistryKeys interface{} `field:"optional" json:"registryKeys" yaml:"registryKeys"`
	// registry_values block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/ad/0.4.4/docs/resources/gpo_security#registry_values GpoSecurity#registry_values}
	RegistryValues interface{} `field:"optional" json:"registryValues" yaml:"registryValues"`
	// restricted_groups block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/ad/0.4.4/docs/resources/gpo_security#restricted_groups GpoSecurity#restricted_groups}
	RestrictedGroups interface{} `field:"optional" json:"restrictedGroups" yaml:"restrictedGroups"`
	// system_log block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/ad/0.4.4/docs/resources/gpo_security#system_log GpoSecurity#system_log}
	SystemLog *GpoSecuritySystemLog `field:"optional" json:"systemLog" yaml:"systemLog"`
	// system_services block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/ad/0.4.4/docs/resources/gpo_security#system_services GpoSecurity#system_services}
	SystemServices interface{} `field:"optional" json:"systemServices" yaml:"systemServices"`
}

