package gposecurity


type GpoSecurityEventAudit struct {
	// Audit credential validation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/ad/0.4.4/docs/resources/gpo_security#audit_account_logon GpoSecurity#audit_account_logon}
	AuditAccountLogon *string `field:"optional" json:"auditAccountLogon" yaml:"auditAccountLogon"`
	// Audit account management events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/ad/0.4.4/docs/resources/gpo_security#audit_account_manage GpoSecurity#audit_account_manage}
	AuditAccountManage *string `field:"optional" json:"auditAccountManage" yaml:"auditAccountManage"`
	// Audit access attempts to AD objects.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/ad/0.4.4/docs/resources/gpo_security#audit_ds_access GpoSecurity#audit_ds_access}
	AuditDsAccess *string `field:"optional" json:"auditDsAccess" yaml:"auditDsAccess"`
	// Audit logon events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/ad/0.4.4/docs/resources/gpo_security#audit_logon_events GpoSecurity#audit_logon_events}
	AuditLogonEvents *string `field:"optional" json:"auditLogonEvents" yaml:"auditLogonEvents"`
	// Audit access attempts to non-AD objects.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/ad/0.4.4/docs/resources/gpo_security#audit_object_access GpoSecurity#audit_object_access}
	AuditObjectAccess *string `field:"optional" json:"auditObjectAccess" yaml:"auditObjectAccess"`
	// Audit attempts to change a policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/ad/0.4.4/docs/resources/gpo_security#audit_policy_change GpoSecurity#audit_policy_change}
	AuditPolicyChange *string `field:"optional" json:"auditPolicyChange" yaml:"auditPolicyChange"`
	// Audit user attempts of exercising user rights.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/ad/0.4.4/docs/resources/gpo_security#audit_privilege_use GpoSecurity#audit_privilege_use}
	AuditPrivilegeUse *string `field:"optional" json:"auditPrivilegeUse" yaml:"auditPrivilegeUse"`
	// Audit process related events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/ad/0.4.4/docs/resources/gpo_security#audit_process_tracking GpoSecurity#audit_process_tracking}
	AuditProcessTracking *string `field:"optional" json:"auditProcessTracking" yaml:"auditProcessTracking"`
	// Audit system events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/ad/0.4.4/docs/resources/gpo_security#audit_system_events GpoSecurity#audit_system_events}
	AuditSystemEvents *string `field:"optional" json:"auditSystemEvents" yaml:"auditSystemEvents"`
}

