package gposecurity


type GpoSecurityRestrictedGroups struct {
	// Comma separated list of group names or SIDs that this group belongs to.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ad/r/gpo_security#group_memberof GpoSecurity#group_memberof}
	GroupMemberof *string `field:"required" json:"groupMemberof" yaml:"groupMemberof"`
	// Comma separated list of group names or SIDs that are members of the group.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ad/r/gpo_security#group_members GpoSecurity#group_members}
	GroupMembers *string `field:"required" json:"groupMembers" yaml:"groupMembers"`
	// Name of the group we are managing.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ad/r/gpo_security#group_name GpoSecurity#group_name}
	GroupName *string `field:"required" json:"groupName" yaml:"groupName"`
}

