// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gposecurity

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-ad-go/ad/v7/jsii"

	"github.com/cdktf/cdktf-provider-ad-go/ad/v7/gposecurity/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GpoSecurityAccountLockoutOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	ForceLogoffWhenHourExpire() *string
	SetForceLogoffWhenHourExpire(val *string)
	ForceLogoffWhenHourExpireInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *GpoSecurityAccountLockout
	SetInternalValue(val *GpoSecurityAccountLockout)
	LockoutBadCount() *string
	SetLockoutBadCount(val *string)
	LockoutBadCountInput() *string
	LockoutDuration() *string
	SetLockoutDuration(val *string)
	LockoutDurationInput() *string
	ResetLockoutCount() *string
	SetResetLockoutCount(val *string)
	ResetLockoutCountInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetForceLogoffWhenHourExpire()
	ResetLockoutBadCount()
	ResetLockoutDuration()
	ResetResetLockoutCount()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GpoSecurityAccountLockoutOutputReference
type jsiiProxy_GpoSecurityAccountLockoutOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GpoSecurityAccountLockoutOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GpoSecurityAccountLockoutOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GpoSecurityAccountLockoutOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GpoSecurityAccountLockoutOutputReference) ForceLogoffWhenHourExpire() *string {
	var returns *string
	_jsii_.Get(
		j,
		"forceLogoffWhenHourExpire",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GpoSecurityAccountLockoutOutputReference) ForceLogoffWhenHourExpireInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"forceLogoffWhenHourExpireInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GpoSecurityAccountLockoutOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GpoSecurityAccountLockoutOutputReference) InternalValue() *GpoSecurityAccountLockout {
	var returns *GpoSecurityAccountLockout
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GpoSecurityAccountLockoutOutputReference) LockoutBadCount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lockoutBadCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GpoSecurityAccountLockoutOutputReference) LockoutBadCountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lockoutBadCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GpoSecurityAccountLockoutOutputReference) LockoutDuration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lockoutDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GpoSecurityAccountLockoutOutputReference) LockoutDurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lockoutDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GpoSecurityAccountLockoutOutputReference) ResetLockoutCount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resetLockoutCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GpoSecurityAccountLockoutOutputReference) ResetLockoutCountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resetLockoutCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GpoSecurityAccountLockoutOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GpoSecurityAccountLockoutOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGpoSecurityAccountLockoutOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GpoSecurityAccountLockoutOutputReference {
	_init_.Initialize()

	if err := validateNewGpoSecurityAccountLockoutOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GpoSecurityAccountLockoutOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-ad.gpoSecurity.GpoSecurityAccountLockoutOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGpoSecurityAccountLockoutOutputReference_Override(g GpoSecurityAccountLockoutOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-ad.gpoSecurity.GpoSecurityAccountLockoutOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GpoSecurityAccountLockoutOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GpoSecurityAccountLockoutOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GpoSecurityAccountLockoutOutputReference)SetForceLogoffWhenHourExpire(val *string) {
	if err := j.validateSetForceLogoffWhenHourExpireParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forceLogoffWhenHourExpire",
		val,
	)
}

func (j *jsiiProxy_GpoSecurityAccountLockoutOutputReference)SetInternalValue(val *GpoSecurityAccountLockout) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GpoSecurityAccountLockoutOutputReference)SetLockoutBadCount(val *string) {
	if err := j.validateSetLockoutBadCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lockoutBadCount",
		val,
	)
}

func (j *jsiiProxy_GpoSecurityAccountLockoutOutputReference)SetLockoutDuration(val *string) {
	if err := j.validateSetLockoutDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lockoutDuration",
		val,
	)
}

func (j *jsiiProxy_GpoSecurityAccountLockoutOutputReference)SetResetLockoutCount(val *string) {
	if err := j.validateSetResetLockoutCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resetLockoutCount",
		val,
	)
}

func (j *jsiiProxy_GpoSecurityAccountLockoutOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GpoSecurityAccountLockoutOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GpoSecurityAccountLockoutOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GpoSecurityAccountLockoutOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GpoSecurityAccountLockoutOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GpoSecurityAccountLockoutOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GpoSecurityAccountLockoutOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GpoSecurityAccountLockoutOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GpoSecurityAccountLockoutOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GpoSecurityAccountLockoutOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GpoSecurityAccountLockoutOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GpoSecurityAccountLockoutOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GpoSecurityAccountLockoutOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GpoSecurityAccountLockoutOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GpoSecurityAccountLockoutOutputReference) ResetForceLogoffWhenHourExpire() {
	_jsii_.InvokeVoid(
		g,
		"resetForceLogoffWhenHourExpire",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GpoSecurityAccountLockoutOutputReference) ResetLockoutBadCount() {
	_jsii_.InvokeVoid(
		g,
		"resetLockoutBadCount",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GpoSecurityAccountLockoutOutputReference) ResetLockoutDuration() {
	_jsii_.InvokeVoid(
		g,
		"resetLockoutDuration",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GpoSecurityAccountLockoutOutputReference) ResetResetLockoutCount() {
	_jsii_.InvokeVoid(
		g,
		"resetResetLockoutCount",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GpoSecurityAccountLockoutOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := g.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GpoSecurityAccountLockoutOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

