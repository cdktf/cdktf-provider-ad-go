// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-ad-go/ad/v7/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-ad-go/ad/v7/provider/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/ad/0.4.4/docs ad}.
type AdProvider interface {
	cdktf.TerraformProvider
	Alias() *string
	SetAlias(val *string)
	AliasInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	DomainController() *string
	SetDomainController(val *string)
	DomainControllerInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	KrbConf() *string
	SetKrbConf(val *string)
	KrbConfInput() *string
	KrbKeytab() *string
	SetKrbKeytab(val *string)
	KrbKeytabInput() *string
	KrbRealm() *string
	SetKrbRealm(val *string)
	KrbRealmInput() *string
	KrbSpn() *string
	SetKrbSpn(val *string)
	KrbSpnInput() *string
	// Experimental.
	MetaAttributes() *map[string]interface{}
	// The tree node.
	Node() constructs.Node
	// Experimental.
	RawOverrides() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformProviderSource() *string
	// Experimental.
	TerraformResourceType() *string
	WinrmHostname() *string
	SetWinrmHostname(val *string)
	WinrmHostnameInput() *string
	WinrmInsecure() interface{}
	SetWinrmInsecure(val interface{})
	WinrmInsecureInput() interface{}
	WinrmPassCredentials() interface{}
	SetWinrmPassCredentials(val interface{})
	WinrmPassCredentialsInput() interface{}
	WinrmPassword() *string
	SetWinrmPassword(val *string)
	WinrmPasswordInput() *string
	WinrmPort() *float64
	SetWinrmPort(val *float64)
	WinrmPortInput() *float64
	WinrmProto() *string
	SetWinrmProto(val *string)
	WinrmProtoInput() *string
	WinrmUseNtlm() interface{}
	SetWinrmUseNtlm(val interface{})
	WinrmUseNtlmInput() interface{}
	WinrmUsername() *string
	SetWinrmUsername(val *string)
	WinrmUsernameInput() *string
	// Experimental.
	AddOverride(path *string, value interface{})
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	ResetAlias()
	ResetDomainController()
	ResetKrbConf()
	ResetKrbKeytab()
	ResetKrbRealm()
	ResetKrbSpn()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetWinrmInsecure()
	ResetWinrmPassCredentials()
	ResetWinrmPort()
	ResetWinrmProto()
	ResetWinrmUseNtlm()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for AdProvider
type jsiiProxy_AdProvider struct {
	internal.Type__cdktfTerraformProvider
}

func (j *jsiiProxy_AdProvider) Alias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdProvider) AliasInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aliasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdProvider) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdProvider) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdProvider) DomainController() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainController",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdProvider) DomainControllerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainControllerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdProvider) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdProvider) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdProvider) KrbConf() *string {
	var returns *string
	_jsii_.Get(
		j,
		"krbConf",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdProvider) KrbConfInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"krbConfInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdProvider) KrbKeytab() *string {
	var returns *string
	_jsii_.Get(
		j,
		"krbKeytab",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdProvider) KrbKeytabInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"krbKeytabInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdProvider) KrbRealm() *string {
	var returns *string
	_jsii_.Get(
		j,
		"krbRealm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdProvider) KrbRealmInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"krbRealmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdProvider) KrbSpn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"krbSpn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdProvider) KrbSpnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"krbSpnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdProvider) MetaAttributes() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"metaAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdProvider) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdProvider) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdProvider) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdProvider) TerraformProviderSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformProviderSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdProvider) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdProvider) WinrmHostname() *string {
	var returns *string
	_jsii_.Get(
		j,
		"winrmHostname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdProvider) WinrmHostnameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"winrmHostnameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdProvider) WinrmInsecure() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"winrmInsecure",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdProvider) WinrmInsecureInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"winrmInsecureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdProvider) WinrmPassCredentials() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"winrmPassCredentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdProvider) WinrmPassCredentialsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"winrmPassCredentialsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdProvider) WinrmPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"winrmPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdProvider) WinrmPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"winrmPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdProvider) WinrmPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"winrmPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdProvider) WinrmPortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"winrmPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdProvider) WinrmProto() *string {
	var returns *string
	_jsii_.Get(
		j,
		"winrmProto",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdProvider) WinrmProtoInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"winrmProtoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdProvider) WinrmUseNtlm() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"winrmUseNtlm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdProvider) WinrmUseNtlmInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"winrmUseNtlmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdProvider) WinrmUsername() *string {
	var returns *string
	_jsii_.Get(
		j,
		"winrmUsername",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdProvider) WinrmUsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"winrmUsernameInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/ad/0.4.4/docs ad} Resource.
func NewAdProvider(scope constructs.Construct, id *string, config *AdProviderConfig) AdProvider {
	_init_.Initialize()

	if err := validateNewAdProviderParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AdProvider{}

	_jsii_.Create(
		"@cdktf/provider-ad.provider.AdProvider",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/ad/0.4.4/docs ad} Resource.
func NewAdProvider_Override(a AdProvider, scope constructs.Construct, id *string, config *AdProviderConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-ad.provider.AdProvider",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AdProvider)SetAlias(val *string) {
	_jsii_.Set(
		j,
		"alias",
		val,
	)
}

func (j *jsiiProxy_AdProvider)SetDomainController(val *string) {
	_jsii_.Set(
		j,
		"domainController",
		val,
	)
}

func (j *jsiiProxy_AdProvider)SetKrbConf(val *string) {
	_jsii_.Set(
		j,
		"krbConf",
		val,
	)
}

func (j *jsiiProxy_AdProvider)SetKrbKeytab(val *string) {
	_jsii_.Set(
		j,
		"krbKeytab",
		val,
	)
}

func (j *jsiiProxy_AdProvider)SetKrbRealm(val *string) {
	_jsii_.Set(
		j,
		"krbRealm",
		val,
	)
}

func (j *jsiiProxy_AdProvider)SetKrbSpn(val *string) {
	_jsii_.Set(
		j,
		"krbSpn",
		val,
	)
}

func (j *jsiiProxy_AdProvider)SetWinrmHostname(val *string) {
	_jsii_.Set(
		j,
		"winrmHostname",
		val,
	)
}

func (j *jsiiProxy_AdProvider)SetWinrmInsecure(val interface{}) {
	if err := j.validateSetWinrmInsecureParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"winrmInsecure",
		val,
	)
}

func (j *jsiiProxy_AdProvider)SetWinrmPassCredentials(val interface{}) {
	if err := j.validateSetWinrmPassCredentialsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"winrmPassCredentials",
		val,
	)
}

func (j *jsiiProxy_AdProvider)SetWinrmPassword(val *string) {
	_jsii_.Set(
		j,
		"winrmPassword",
		val,
	)
}

func (j *jsiiProxy_AdProvider)SetWinrmPort(val *float64) {
	_jsii_.Set(
		j,
		"winrmPort",
		val,
	)
}

func (j *jsiiProxy_AdProvider)SetWinrmProto(val *string) {
	_jsii_.Set(
		j,
		"winrmProto",
		val,
	)
}

func (j *jsiiProxy_AdProvider)SetWinrmUseNtlm(val interface{}) {
	if err := j.validateSetWinrmUseNtlmParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"winrmUseNtlm",
		val,
	)
}

func (j *jsiiProxy_AdProvider)SetWinrmUsername(val *string) {
	_jsii_.Set(
		j,
		"winrmUsername",
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
func AdProvider_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAdProvider_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-ad.provider.AdProvider",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AdProvider_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAdProvider_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-ad.provider.AdProvider",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AdProvider_IsTerraformProvider(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAdProvider_IsTerraformProviderParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-ad.provider.AdProvider",
		"isTerraformProvider",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AdProvider_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-ad.provider.AdProvider",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AdProvider) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AdProvider) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AdProvider) ResetAlias() {
	_jsii_.InvokeVoid(
		a,
		"resetAlias",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AdProvider) ResetDomainController() {
	_jsii_.InvokeVoid(
		a,
		"resetDomainController",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AdProvider) ResetKrbConf() {
	_jsii_.InvokeVoid(
		a,
		"resetKrbConf",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AdProvider) ResetKrbKeytab() {
	_jsii_.InvokeVoid(
		a,
		"resetKrbKeytab",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AdProvider) ResetKrbRealm() {
	_jsii_.InvokeVoid(
		a,
		"resetKrbRealm",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AdProvider) ResetKrbSpn() {
	_jsii_.InvokeVoid(
		a,
		"resetKrbSpn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AdProvider) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AdProvider) ResetWinrmInsecure() {
	_jsii_.InvokeVoid(
		a,
		"resetWinrmInsecure",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AdProvider) ResetWinrmPassCredentials() {
	_jsii_.InvokeVoid(
		a,
		"resetWinrmPassCredentials",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AdProvider) ResetWinrmPort() {
	_jsii_.InvokeVoid(
		a,
		"resetWinrmPort",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AdProvider) ResetWinrmProto() {
	_jsii_.InvokeVoid(
		a,
		"resetWinrmProto",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AdProvider) ResetWinrmUseNtlm() {
	_jsii_.InvokeVoid(
		a,
		"resetWinrmUseNtlm",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AdProvider) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AdProvider) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AdProvider) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AdProvider) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

