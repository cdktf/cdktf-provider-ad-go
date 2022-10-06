package groupmembership

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GroupMembershipConfig struct {
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
	// The ID of the group.
	//
	// This can be a GUID, a SID, a Distinguished Name, or the SAM Account Name of the group.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ad/r/group_membership#group_id GroupMembership#group_id}
	GroupId *string `field:"required" json:"groupId" yaml:"groupId"`
	// A list of member AD Principals.
	//
	// Each principal can be identified by its GUID, SID, Distinguished Name, or SAM Account Name. Only one is required
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ad/r/group_membership#group_members GroupMembership#group_members}
	GroupMembers *[]*string `field:"required" json:"groupMembers" yaml:"groupMembers"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ad/r/group_membership#id GroupMembership#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

