package gposecurity


type GpoSecurityAuditLog struct {
	// Control log retention.
	//
	// Values: 0: overwrite events as needed, 1: overwrite events as specified specified by `retention_days`, 2: never overwrite events.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ad/r/gpo_security#audit_log_retention_period GpoSecurity#audit_log_retention_period}
	AuditLogRetentionPeriod *string `field:"optional" json:"auditLogRetentionPeriod" yaml:"auditLogRetentionPeriod"`
	// Maximum size of log in KiloBytes. (64-4194240).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ad/r/gpo_security#maximum_log_size GpoSecurity#maximum_log_size}
	MaximumLogSize *string `field:"optional" json:"maximumLogSize" yaml:"maximumLogSize"`
	// Restrict access to logs for guest users. A non-zero value restricts access to guest users.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ad/r/gpo_security#restrict_guest_access GpoSecurity#restrict_guest_access}
	RestrictGuestAccess *string `field:"optional" json:"restrictGuestAccess" yaml:"restrictGuestAccess"`
	// Number of days before new events overwrite old events. (1-365).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ad/r/gpo_security#retention_days GpoSecurity#retention_days}
	RetentionDays *string `field:"optional" json:"retentionDays" yaml:"retentionDays"`
}

