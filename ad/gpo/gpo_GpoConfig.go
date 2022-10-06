package gpo

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GpoConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
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
	// Name of the GPO.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ad/r/gpo#name Gpo#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Description of the GPO.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ad/r/gpo#description Gpo#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Domain of the GPO.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ad/r/gpo#domain Gpo#domain}
	Domain *string `field:"optional" json:"domain" yaml:"domain"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ad/r/gpo#id Gpo#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Status of the GPO. Can be one of `AllSettingsEnabled`, `UserSettingsDisabled`, `ComputerSettingsDisabled`, or `AllSettingsDisabled` (case sensitive).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ad/r/gpo#status Gpo#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
}

