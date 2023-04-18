package dataadcomputer

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataAdComputerConfig struct {
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
	// The OU's identifier. It can be the OU's GUID, SID, Distinguished Name, or SAM Account Name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/ad/0.4.4/docs/data-sources/computer#computer_id DataAdComputer#computer_id}
	ComputerId *string `field:"optional" json:"computerId" yaml:"computerId"`
	// The Distinguished Name of the computer object.
	//
	// This field is deprecated in favour of `computer_id`. In the future this field will be read-only.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/ad/0.4.4/docs/data-sources/computer#dn DataAdComputer#dn}
	Dn *string `field:"optional" json:"dn" yaml:"dn"`
	// The GUID of the computer object.
	//
	// This field is deprecated in favour of `computer_id`. In the future this field will be read-only.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/ad/0.4.4/docs/data-sources/computer#guid DataAdComputer#guid}
	Guid *string `field:"optional" json:"guid" yaml:"guid"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/ad/0.4.4/docs/data-sources/computer#id DataAdComputer#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

