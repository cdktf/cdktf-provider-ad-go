package gposecurity


type GpoSecurityFilesystem struct {
	// Security descriptor to apply. (https://docs.microsoft.com/en-us/openspecs/windows_protocols/ms-dtyp/f4296d69-1c0f-491f-9587-a960b292d070).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ad/r/gpo_security#acl GpoSecurity#acl}
	Acl *string `field:"required" json:"acl" yaml:"acl"`
	// Path of the file or directory.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ad/r/gpo_security#path GpoSecurity#path}
	Path *string `field:"required" json:"path" yaml:"path"`
	// Control permission propagation.
	//
	// 0: Propagate permissions to all subfolders and files, 1: Replace existing permissions on all subfolders and files, 2: Do not allow permissions to be replaced.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ad/r/gpo_security#propagation_mode GpoSecurity#propagation_mode}
	PropagationMode *string `field:"required" json:"propagationMode" yaml:"propagationMode"`
}

