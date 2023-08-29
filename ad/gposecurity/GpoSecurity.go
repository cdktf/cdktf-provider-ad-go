// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gposecurity

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-ad-go/ad/v6/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-ad-go/ad/v6/gposecurity/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/ad/0.4.4/docs/resources/gpo_security ad_gpo_security}.
type GpoSecurity interface {
	cdktf.TerraformResource
	AccountLockout() GpoSecurityAccountLockoutOutputReference
	AccountLockoutInput() *GpoSecurityAccountLockout
	ApplicationLog() GpoSecurityApplicationLogOutputReference
	ApplicationLogInput() *GpoSecurityApplicationLog
	AuditLog() GpoSecurityAuditLogOutputReference
	AuditLogInput() *GpoSecurityAuditLog
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	EventAudit() GpoSecurityEventAuditOutputReference
	EventAuditInput() *GpoSecurityEventAudit
	Filesystem() GpoSecurityFilesystemList
	FilesystemInput() interface{}
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GpoContainer() *string
	SetGpoContainer(val *string)
	GpoContainerInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	KerberosPolicy() GpoSecurityKerberosPolicyOutputReference
	KerberosPolicyInput() *GpoSecurityKerberosPolicy
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
	PasswordPolicies() GpoSecurityPasswordPoliciesOutputReference
	PasswordPoliciesInput() *GpoSecurityPasswordPolicies
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	RegistryKeys() GpoSecurityRegistryKeysList
	RegistryKeysInput() interface{}
	RegistryValues() GpoSecurityRegistryValuesList
	RegistryValuesInput() interface{}
	RestrictedGroups() GpoSecurityRestrictedGroupsList
	RestrictedGroupsInput() interface{}
	SystemLog() GpoSecuritySystemLogOutputReference
	SystemLogInput() *GpoSecuritySystemLog
	SystemServices() GpoSecuritySystemServicesList
	SystemServicesInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutAccountLockout(value *GpoSecurityAccountLockout)
	PutApplicationLog(value *GpoSecurityApplicationLog)
	PutAuditLog(value *GpoSecurityAuditLog)
	PutEventAudit(value *GpoSecurityEventAudit)
	PutFilesystem(value interface{})
	PutKerberosPolicy(value *GpoSecurityKerberosPolicy)
	PutPasswordPolicies(value *GpoSecurityPasswordPolicies)
	PutRegistryKeys(value interface{})
	PutRegistryValues(value interface{})
	PutRestrictedGroups(value interface{})
	PutSystemLog(value *GpoSecuritySystemLog)
	PutSystemServices(value interface{})
	ResetAccountLockout()
	ResetApplicationLog()
	ResetAuditLog()
	ResetEventAudit()
	ResetFilesystem()
	ResetId()
	ResetKerberosPolicy()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPasswordPolicies()
	ResetRegistryKeys()
	ResetRegistryValues()
	ResetRestrictedGroups()
	ResetSystemLog()
	ResetSystemServices()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for GpoSecurity
type jsiiProxy_GpoSecurity struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GpoSecurity) AccountLockout() GpoSecurityAccountLockoutOutputReference {
	var returns GpoSecurityAccountLockoutOutputReference
	_jsii_.Get(
		j,
		"accountLockout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GpoSecurity) AccountLockoutInput() *GpoSecurityAccountLockout {
	var returns *GpoSecurityAccountLockout
	_jsii_.Get(
		j,
		"accountLockoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GpoSecurity) ApplicationLog() GpoSecurityApplicationLogOutputReference {
	var returns GpoSecurityApplicationLogOutputReference
	_jsii_.Get(
		j,
		"applicationLog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GpoSecurity) ApplicationLogInput() *GpoSecurityApplicationLog {
	var returns *GpoSecurityApplicationLog
	_jsii_.Get(
		j,
		"applicationLogInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GpoSecurity) AuditLog() GpoSecurityAuditLogOutputReference {
	var returns GpoSecurityAuditLogOutputReference
	_jsii_.Get(
		j,
		"auditLog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GpoSecurity) AuditLogInput() *GpoSecurityAuditLog {
	var returns *GpoSecurityAuditLog
	_jsii_.Get(
		j,
		"auditLogInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GpoSecurity) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GpoSecurity) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GpoSecurity) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GpoSecurity) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GpoSecurity) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GpoSecurity) EventAudit() GpoSecurityEventAuditOutputReference {
	var returns GpoSecurityEventAuditOutputReference
	_jsii_.Get(
		j,
		"eventAudit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GpoSecurity) EventAuditInput() *GpoSecurityEventAudit {
	var returns *GpoSecurityEventAudit
	_jsii_.Get(
		j,
		"eventAuditInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GpoSecurity) Filesystem() GpoSecurityFilesystemList {
	var returns GpoSecurityFilesystemList
	_jsii_.Get(
		j,
		"filesystem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GpoSecurity) FilesystemInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"filesystemInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GpoSecurity) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GpoSecurity) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GpoSecurity) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GpoSecurity) GpoContainer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gpoContainer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GpoSecurity) GpoContainerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gpoContainerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GpoSecurity) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GpoSecurity) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GpoSecurity) KerberosPolicy() GpoSecurityKerberosPolicyOutputReference {
	var returns GpoSecurityKerberosPolicyOutputReference
	_jsii_.Get(
		j,
		"kerberosPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GpoSecurity) KerberosPolicyInput() *GpoSecurityKerberosPolicy {
	var returns *GpoSecurityKerberosPolicy
	_jsii_.Get(
		j,
		"kerberosPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GpoSecurity) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GpoSecurity) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GpoSecurity) PasswordPolicies() GpoSecurityPasswordPoliciesOutputReference {
	var returns GpoSecurityPasswordPoliciesOutputReference
	_jsii_.Get(
		j,
		"passwordPolicies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GpoSecurity) PasswordPoliciesInput() *GpoSecurityPasswordPolicies {
	var returns *GpoSecurityPasswordPolicies
	_jsii_.Get(
		j,
		"passwordPoliciesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GpoSecurity) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GpoSecurity) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GpoSecurity) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GpoSecurity) RegistryKeys() GpoSecurityRegistryKeysList {
	var returns GpoSecurityRegistryKeysList
	_jsii_.Get(
		j,
		"registryKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GpoSecurity) RegistryKeysInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"registryKeysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GpoSecurity) RegistryValues() GpoSecurityRegistryValuesList {
	var returns GpoSecurityRegistryValuesList
	_jsii_.Get(
		j,
		"registryValues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GpoSecurity) RegistryValuesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"registryValuesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GpoSecurity) RestrictedGroups() GpoSecurityRestrictedGroupsList {
	var returns GpoSecurityRestrictedGroupsList
	_jsii_.Get(
		j,
		"restrictedGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GpoSecurity) RestrictedGroupsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"restrictedGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GpoSecurity) SystemLog() GpoSecuritySystemLogOutputReference {
	var returns GpoSecuritySystemLogOutputReference
	_jsii_.Get(
		j,
		"systemLog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GpoSecurity) SystemLogInput() *GpoSecuritySystemLog {
	var returns *GpoSecuritySystemLog
	_jsii_.Get(
		j,
		"systemLogInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GpoSecurity) SystemServices() GpoSecuritySystemServicesList {
	var returns GpoSecuritySystemServicesList
	_jsii_.Get(
		j,
		"systemServices",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GpoSecurity) SystemServicesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"systemServicesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GpoSecurity) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GpoSecurity) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GpoSecurity) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/ad/0.4.4/docs/resources/gpo_security ad_gpo_security} Resource.
func NewGpoSecurity(scope constructs.Construct, id *string, config *GpoSecurityConfig) GpoSecurity {
	_init_.Initialize()

	if err := validateNewGpoSecurityParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GpoSecurity{}

	_jsii_.Create(
		"@cdktf/provider-ad.gpoSecurity.GpoSecurity",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/ad/0.4.4/docs/resources/gpo_security ad_gpo_security} Resource.
func NewGpoSecurity_Override(g GpoSecurity, scope constructs.Construct, id *string, config *GpoSecurityConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-ad.gpoSecurity.GpoSecurity",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GpoSecurity)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GpoSecurity)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GpoSecurity)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GpoSecurity)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GpoSecurity)SetGpoContainer(val *string) {
	if err := j.validateSetGpoContainerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gpoContainer",
		val,
	)
}

func (j *jsiiProxy_GpoSecurity)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GpoSecurity)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GpoSecurity)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GpoSecurity)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func GpoSecurity_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGpoSecurity_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-ad.gpoSecurity.GpoSecurity",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GpoSecurity_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGpoSecurity_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-ad.gpoSecurity.GpoSecurity",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GpoSecurity_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGpoSecurity_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-ad.gpoSecurity.GpoSecurity",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GpoSecurity_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-ad.gpoSecurity.GpoSecurity",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GpoSecurity) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GpoSecurity) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := g.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GpoSecurity) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GpoSecurity) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := g.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		g,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GpoSecurity) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := g.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GpoSecurity) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := g.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		g,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GpoSecurity) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := g.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		g,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GpoSecurity) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := g.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		g,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GpoSecurity) GetStringAttribute(terraformAttribute *string) *string {
	if err := g.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GpoSecurity) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := g.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		g,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GpoSecurity) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GpoSecurity) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GpoSecurity) PutAccountLockout(value *GpoSecurityAccountLockout) {
	if err := g.validatePutAccountLockoutParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAccountLockout",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GpoSecurity) PutApplicationLog(value *GpoSecurityApplicationLog) {
	if err := g.validatePutApplicationLogParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putApplicationLog",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GpoSecurity) PutAuditLog(value *GpoSecurityAuditLog) {
	if err := g.validatePutAuditLogParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAuditLog",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GpoSecurity) PutEventAudit(value *GpoSecurityEventAudit) {
	if err := g.validatePutEventAuditParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putEventAudit",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GpoSecurity) PutFilesystem(value interface{}) {
	if err := g.validatePutFilesystemParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putFilesystem",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GpoSecurity) PutKerberosPolicy(value *GpoSecurityKerberosPolicy) {
	if err := g.validatePutKerberosPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putKerberosPolicy",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GpoSecurity) PutPasswordPolicies(value *GpoSecurityPasswordPolicies) {
	if err := g.validatePutPasswordPoliciesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPasswordPolicies",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GpoSecurity) PutRegistryKeys(value interface{}) {
	if err := g.validatePutRegistryKeysParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putRegistryKeys",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GpoSecurity) PutRegistryValues(value interface{}) {
	if err := g.validatePutRegistryValuesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putRegistryValues",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GpoSecurity) PutRestrictedGroups(value interface{}) {
	if err := g.validatePutRestrictedGroupsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putRestrictedGroups",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GpoSecurity) PutSystemLog(value *GpoSecuritySystemLog) {
	if err := g.validatePutSystemLogParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSystemLog",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GpoSecurity) PutSystemServices(value interface{}) {
	if err := g.validatePutSystemServicesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSystemServices",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GpoSecurity) ResetAccountLockout() {
	_jsii_.InvokeVoid(
		g,
		"resetAccountLockout",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GpoSecurity) ResetApplicationLog() {
	_jsii_.InvokeVoid(
		g,
		"resetApplicationLog",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GpoSecurity) ResetAuditLog() {
	_jsii_.InvokeVoid(
		g,
		"resetAuditLog",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GpoSecurity) ResetEventAudit() {
	_jsii_.InvokeVoid(
		g,
		"resetEventAudit",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GpoSecurity) ResetFilesystem() {
	_jsii_.InvokeVoid(
		g,
		"resetFilesystem",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GpoSecurity) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GpoSecurity) ResetKerberosPolicy() {
	_jsii_.InvokeVoid(
		g,
		"resetKerberosPolicy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GpoSecurity) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GpoSecurity) ResetPasswordPolicies() {
	_jsii_.InvokeVoid(
		g,
		"resetPasswordPolicies",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GpoSecurity) ResetRegistryKeys() {
	_jsii_.InvokeVoid(
		g,
		"resetRegistryKeys",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GpoSecurity) ResetRegistryValues() {
	_jsii_.InvokeVoid(
		g,
		"resetRegistryValues",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GpoSecurity) ResetRestrictedGroups() {
	_jsii_.InvokeVoid(
		g,
		"resetRestrictedGroups",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GpoSecurity) ResetSystemLog() {
	_jsii_.InvokeVoid(
		g,
		"resetSystemLog",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GpoSecurity) ResetSystemServices() {
	_jsii_.InvokeVoid(
		g,
		"resetSystemServices",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GpoSecurity) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GpoSecurity) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GpoSecurity) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GpoSecurity) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

