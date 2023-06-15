package gposecurity

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-ad-go/ad/v5/jsii"

	"github.com/cdktf/cdktf-provider-ad-go/ad/v5/gposecurity/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GpoSecurityEventAuditOutputReference interface {
	cdktf.ComplexObject
	AuditAccountLogon() *string
	SetAuditAccountLogon(val *string)
	AuditAccountLogonInput() *string
	AuditAccountManage() *string
	SetAuditAccountManage(val *string)
	AuditAccountManageInput() *string
	AuditDsAccess() *string
	SetAuditDsAccess(val *string)
	AuditDsAccessInput() *string
	AuditLogonEvents() *string
	SetAuditLogonEvents(val *string)
	AuditLogonEventsInput() *string
	AuditObjectAccess() *string
	SetAuditObjectAccess(val *string)
	AuditObjectAccessInput() *string
	AuditPolicyChange() *string
	SetAuditPolicyChange(val *string)
	AuditPolicyChangeInput() *string
	AuditPrivilegeUse() *string
	SetAuditPrivilegeUse(val *string)
	AuditPrivilegeUseInput() *string
	AuditProcessTracking() *string
	SetAuditProcessTracking(val *string)
	AuditProcessTrackingInput() *string
	AuditSystemEvents() *string
	SetAuditSystemEvents(val *string)
	AuditSystemEventsInput() *string
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
	InternalValue() *GpoSecurityEventAudit
	SetInternalValue(val *GpoSecurityEventAudit)
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
	ResetAuditAccountLogon()
	ResetAuditAccountManage()
	ResetAuditDsAccess()
	ResetAuditLogonEvents()
	ResetAuditObjectAccess()
	ResetAuditPolicyChange()
	ResetAuditPrivilegeUse()
	ResetAuditProcessTracking()
	ResetAuditSystemEvents()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GpoSecurityEventAuditOutputReference
type jsiiProxy_GpoSecurityEventAuditOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GpoSecurityEventAuditOutputReference) AuditAccountLogon() *string {
	var returns *string
	_jsii_.Get(
		j,
		"auditAccountLogon",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GpoSecurityEventAuditOutputReference) AuditAccountLogonInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"auditAccountLogonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GpoSecurityEventAuditOutputReference) AuditAccountManage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"auditAccountManage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GpoSecurityEventAuditOutputReference) AuditAccountManageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"auditAccountManageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GpoSecurityEventAuditOutputReference) AuditDsAccess() *string {
	var returns *string
	_jsii_.Get(
		j,
		"auditDsAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GpoSecurityEventAuditOutputReference) AuditDsAccessInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"auditDsAccessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GpoSecurityEventAuditOutputReference) AuditLogonEvents() *string {
	var returns *string
	_jsii_.Get(
		j,
		"auditLogonEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GpoSecurityEventAuditOutputReference) AuditLogonEventsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"auditLogonEventsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GpoSecurityEventAuditOutputReference) AuditObjectAccess() *string {
	var returns *string
	_jsii_.Get(
		j,
		"auditObjectAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GpoSecurityEventAuditOutputReference) AuditObjectAccessInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"auditObjectAccessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GpoSecurityEventAuditOutputReference) AuditPolicyChange() *string {
	var returns *string
	_jsii_.Get(
		j,
		"auditPolicyChange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GpoSecurityEventAuditOutputReference) AuditPolicyChangeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"auditPolicyChangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GpoSecurityEventAuditOutputReference) AuditPrivilegeUse() *string {
	var returns *string
	_jsii_.Get(
		j,
		"auditPrivilegeUse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GpoSecurityEventAuditOutputReference) AuditPrivilegeUseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"auditPrivilegeUseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GpoSecurityEventAuditOutputReference) AuditProcessTracking() *string {
	var returns *string
	_jsii_.Get(
		j,
		"auditProcessTracking",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GpoSecurityEventAuditOutputReference) AuditProcessTrackingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"auditProcessTrackingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GpoSecurityEventAuditOutputReference) AuditSystemEvents() *string {
	var returns *string
	_jsii_.Get(
		j,
		"auditSystemEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GpoSecurityEventAuditOutputReference) AuditSystemEventsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"auditSystemEventsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GpoSecurityEventAuditOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GpoSecurityEventAuditOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GpoSecurityEventAuditOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GpoSecurityEventAuditOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GpoSecurityEventAuditOutputReference) InternalValue() *GpoSecurityEventAudit {
	var returns *GpoSecurityEventAudit
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GpoSecurityEventAuditOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GpoSecurityEventAuditOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGpoSecurityEventAuditOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GpoSecurityEventAuditOutputReference {
	_init_.Initialize()

	if err := validateNewGpoSecurityEventAuditOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GpoSecurityEventAuditOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-ad.gpoSecurity.GpoSecurityEventAuditOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGpoSecurityEventAuditOutputReference_Override(g GpoSecurityEventAuditOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-ad.gpoSecurity.GpoSecurityEventAuditOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GpoSecurityEventAuditOutputReference)SetAuditAccountLogon(val *string) {
	if err := j.validateSetAuditAccountLogonParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"auditAccountLogon",
		val,
	)
}

func (j *jsiiProxy_GpoSecurityEventAuditOutputReference)SetAuditAccountManage(val *string) {
	if err := j.validateSetAuditAccountManageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"auditAccountManage",
		val,
	)
}

func (j *jsiiProxy_GpoSecurityEventAuditOutputReference)SetAuditDsAccess(val *string) {
	if err := j.validateSetAuditDsAccessParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"auditDsAccess",
		val,
	)
}

func (j *jsiiProxy_GpoSecurityEventAuditOutputReference)SetAuditLogonEvents(val *string) {
	if err := j.validateSetAuditLogonEventsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"auditLogonEvents",
		val,
	)
}

func (j *jsiiProxy_GpoSecurityEventAuditOutputReference)SetAuditObjectAccess(val *string) {
	if err := j.validateSetAuditObjectAccessParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"auditObjectAccess",
		val,
	)
}

func (j *jsiiProxy_GpoSecurityEventAuditOutputReference)SetAuditPolicyChange(val *string) {
	if err := j.validateSetAuditPolicyChangeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"auditPolicyChange",
		val,
	)
}

func (j *jsiiProxy_GpoSecurityEventAuditOutputReference)SetAuditPrivilegeUse(val *string) {
	if err := j.validateSetAuditPrivilegeUseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"auditPrivilegeUse",
		val,
	)
}

func (j *jsiiProxy_GpoSecurityEventAuditOutputReference)SetAuditProcessTracking(val *string) {
	if err := j.validateSetAuditProcessTrackingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"auditProcessTracking",
		val,
	)
}

func (j *jsiiProxy_GpoSecurityEventAuditOutputReference)SetAuditSystemEvents(val *string) {
	if err := j.validateSetAuditSystemEventsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"auditSystemEvents",
		val,
	)
}

func (j *jsiiProxy_GpoSecurityEventAuditOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GpoSecurityEventAuditOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GpoSecurityEventAuditOutputReference)SetInternalValue(val *GpoSecurityEventAudit) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GpoSecurityEventAuditOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GpoSecurityEventAuditOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GpoSecurityEventAuditOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GpoSecurityEventAuditOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GpoSecurityEventAuditOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GpoSecurityEventAuditOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GpoSecurityEventAuditOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GpoSecurityEventAuditOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GpoSecurityEventAuditOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GpoSecurityEventAuditOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GpoSecurityEventAuditOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GpoSecurityEventAuditOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GpoSecurityEventAuditOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GpoSecurityEventAuditOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GpoSecurityEventAuditOutputReference) ResetAuditAccountLogon() {
	_jsii_.InvokeVoid(
		g,
		"resetAuditAccountLogon",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GpoSecurityEventAuditOutputReference) ResetAuditAccountManage() {
	_jsii_.InvokeVoid(
		g,
		"resetAuditAccountManage",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GpoSecurityEventAuditOutputReference) ResetAuditDsAccess() {
	_jsii_.InvokeVoid(
		g,
		"resetAuditDsAccess",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GpoSecurityEventAuditOutputReference) ResetAuditLogonEvents() {
	_jsii_.InvokeVoid(
		g,
		"resetAuditLogonEvents",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GpoSecurityEventAuditOutputReference) ResetAuditObjectAccess() {
	_jsii_.InvokeVoid(
		g,
		"resetAuditObjectAccess",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GpoSecurityEventAuditOutputReference) ResetAuditPolicyChange() {
	_jsii_.InvokeVoid(
		g,
		"resetAuditPolicyChange",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GpoSecurityEventAuditOutputReference) ResetAuditPrivilegeUse() {
	_jsii_.InvokeVoid(
		g,
		"resetAuditPrivilegeUse",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GpoSecurityEventAuditOutputReference) ResetAuditProcessTracking() {
	_jsii_.InvokeVoid(
		g,
		"resetAuditProcessTracking",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GpoSecurityEventAuditOutputReference) ResetAuditSystemEvents() {
	_jsii_.InvokeVoid(
		g,
		"resetAuditSystemEvents",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GpoSecurityEventAuditOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GpoSecurityEventAuditOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

