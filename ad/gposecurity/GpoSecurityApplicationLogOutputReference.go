// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gposecurity

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-ad-go/ad/v8/jsii"

	"github.com/cdktf/cdktf-provider-ad-go/ad/v8/gposecurity/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GpoSecurityApplicationLogOutputReference interface {
	cdktf.ComplexObject
	AuditLogRetentionPeriod() *string
	SetAuditLogRetentionPeriod(val *string)
	AuditLogRetentionPeriodInput() *string
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
	// Experimental.
	Fqn() *string
	InternalValue() *GpoSecurityApplicationLog
	SetInternalValue(val *GpoSecurityApplicationLog)
	MaximumLogSize() *string
	SetMaximumLogSize(val *string)
	MaximumLogSizeInput() *string
	RestrictGuestAccess() *string
	SetRestrictGuestAccess(val *string)
	RestrictGuestAccessInput() *string
	RetentionDays() *string
	SetRetentionDays(val *string)
	RetentionDaysInput() *string
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
	ResetAuditLogRetentionPeriod()
	ResetMaximumLogSize()
	ResetRestrictGuestAccess()
	ResetRetentionDays()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GpoSecurityApplicationLogOutputReference
type jsiiProxy_GpoSecurityApplicationLogOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GpoSecurityApplicationLogOutputReference) AuditLogRetentionPeriod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"auditLogRetentionPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GpoSecurityApplicationLogOutputReference) AuditLogRetentionPeriodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"auditLogRetentionPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GpoSecurityApplicationLogOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GpoSecurityApplicationLogOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GpoSecurityApplicationLogOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GpoSecurityApplicationLogOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GpoSecurityApplicationLogOutputReference) InternalValue() *GpoSecurityApplicationLog {
	var returns *GpoSecurityApplicationLog
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GpoSecurityApplicationLogOutputReference) MaximumLogSize() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maximumLogSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GpoSecurityApplicationLogOutputReference) MaximumLogSizeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maximumLogSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GpoSecurityApplicationLogOutputReference) RestrictGuestAccess() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restrictGuestAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GpoSecurityApplicationLogOutputReference) RestrictGuestAccessInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restrictGuestAccessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GpoSecurityApplicationLogOutputReference) RetentionDays() *string {
	var returns *string
	_jsii_.Get(
		j,
		"retentionDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GpoSecurityApplicationLogOutputReference) RetentionDaysInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"retentionDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GpoSecurityApplicationLogOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GpoSecurityApplicationLogOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGpoSecurityApplicationLogOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GpoSecurityApplicationLogOutputReference {
	_init_.Initialize()

	if err := validateNewGpoSecurityApplicationLogOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GpoSecurityApplicationLogOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-ad.gpoSecurity.GpoSecurityApplicationLogOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGpoSecurityApplicationLogOutputReference_Override(g GpoSecurityApplicationLogOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-ad.gpoSecurity.GpoSecurityApplicationLogOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GpoSecurityApplicationLogOutputReference)SetAuditLogRetentionPeriod(val *string) {
	if err := j.validateSetAuditLogRetentionPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"auditLogRetentionPeriod",
		val,
	)
}

func (j *jsiiProxy_GpoSecurityApplicationLogOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GpoSecurityApplicationLogOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GpoSecurityApplicationLogOutputReference)SetInternalValue(val *GpoSecurityApplicationLog) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GpoSecurityApplicationLogOutputReference)SetMaximumLogSize(val *string) {
	if err := j.validateSetMaximumLogSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maximumLogSize",
		val,
	)
}

func (j *jsiiProxy_GpoSecurityApplicationLogOutputReference)SetRestrictGuestAccess(val *string) {
	if err := j.validateSetRestrictGuestAccessParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"restrictGuestAccess",
		val,
	)
}

func (j *jsiiProxy_GpoSecurityApplicationLogOutputReference)SetRetentionDays(val *string) {
	if err := j.validateSetRetentionDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retentionDays",
		val,
	)
}

func (j *jsiiProxy_GpoSecurityApplicationLogOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GpoSecurityApplicationLogOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GpoSecurityApplicationLogOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GpoSecurityApplicationLogOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GpoSecurityApplicationLogOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GpoSecurityApplicationLogOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GpoSecurityApplicationLogOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GpoSecurityApplicationLogOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GpoSecurityApplicationLogOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GpoSecurityApplicationLogOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GpoSecurityApplicationLogOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GpoSecurityApplicationLogOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GpoSecurityApplicationLogOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GpoSecurityApplicationLogOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GpoSecurityApplicationLogOutputReference) ResetAuditLogRetentionPeriod() {
	_jsii_.InvokeVoid(
		g,
		"resetAuditLogRetentionPeriod",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GpoSecurityApplicationLogOutputReference) ResetMaximumLogSize() {
	_jsii_.InvokeVoid(
		g,
		"resetMaximumLogSize",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GpoSecurityApplicationLogOutputReference) ResetRestrictGuestAccess() {
	_jsii_.InvokeVoid(
		g,
		"resetRestrictGuestAccess",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GpoSecurityApplicationLogOutputReference) ResetRetentionDays() {
	_jsii_.InvokeVoid(
		g,
		"resetRetentionDays",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GpoSecurityApplicationLogOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GpoSecurityApplicationLogOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

