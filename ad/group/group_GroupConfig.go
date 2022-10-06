package group

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GroupConfig struct {
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
	// A DN of a container object holding the group.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ad/r/group#container Group#container}
	Container *string `field:"required" json:"container" yaml:"container"`
	// The name of the group.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ad/r/group#name Group#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The pre-win2k name of the group.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ad/r/group#sam_account_name Group#sam_account_name}
	SamAccountName *string `field:"required" json:"samAccountName" yaml:"samAccountName"`
	// The group's category. Can be one of `distribution` or `security` (case sensitive).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ad/r/group#category Group#category}
	Category *string `field:"optional" json:"category" yaml:"category"`
	// Description of the Group.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ad/r/group#description Group#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ad/r/group#id Group#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The group's scope. Can be one of `global`, `domainlocal`, or `universal` (case sensitive).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ad/r/group#scope Group#scope}
	Scope *string `field:"optional" json:"scope" yaml:"scope"`
}

