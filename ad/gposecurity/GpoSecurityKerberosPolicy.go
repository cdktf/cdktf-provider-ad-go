package gposecurity


type GpoSecurityKerberosPolicy struct {
	// Maximum time difference, in minutes, between the client clock and the server clock. (0-99999).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/ad/0.4.4/docs/resources/gpo_security#max_clock_skew GpoSecurity#max_clock_skew}
	MaxClockSkew *string `field:"optional" json:"maxClockSkew" yaml:"maxClockSkew"`
	// Number of days during which a ticket-granting ticket can be renewed (0-99999).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/ad/0.4.4/docs/resources/gpo_security#max_renew_age GpoSecurity#max_renew_age}
	MaxRenewAge *string `field:"optional" json:"maxRenewAge" yaml:"maxRenewAge"`
	// Maximum amount of minutes a ticket must be valid to access a service or resource.
	//
	// Minimum should be 10 and maximum should be equal to `max_ticket_age`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/ad/0.4.4/docs/resources/gpo_security#max_service_age GpoSecurity#max_service_age}
	MaxServiceAge *string `field:"optional" json:"maxServiceAge" yaml:"maxServiceAge"`
	// Maximum amount of hours a ticket-granting ticket is valid (0-99999).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/ad/0.4.4/docs/resources/gpo_security#max_ticket_age GpoSecurity#max_ticket_age}
	MaxTicketAge *string `field:"optional" json:"maxTicketAge" yaml:"maxTicketAge"`
	// Control if the session ticket is validated for every request. A non-zero value disables the policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/ad/0.4.4/docs/resources/gpo_security#ticket_validate_client GpoSecurity#ticket_validate_client}
	TicketValidateClient *string `field:"optional" json:"ticketValidateClient" yaml:"ticketValidateClient"`
}

