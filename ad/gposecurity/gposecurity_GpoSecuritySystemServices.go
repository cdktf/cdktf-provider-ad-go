package gposecurity


type GpoSecuritySystemServices struct {
	// Security descriptor to apply. (https://docs.microsoft.com/en-us/openspecs/windows_protocols/ms-dtyp/f4296d69-1c0f-491f-9587-a960b292d070).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ad/r/gpo_security#acl GpoSecurity#acl}
	Acl *string `field:"required" json:"acl" yaml:"acl"`
	// Name of the service.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ad/r/gpo_security#service_name GpoSecurity#service_name}
	ServiceName *string `field:"required" json:"serviceName" yaml:"serviceName"`
	// Startup mode of the service. Possible values are 2: Automatic, 3: Manual, 4: Disabled.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ad/r/gpo_security#startup_mode GpoSecurity#startup_mode}
	StartupMode *string `field:"required" json:"startupMode" yaml:"startupMode"`
}

