package gposecurity

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-ad-go/ad/v4/jsii"

	"github.com/cdktf/cdktf-provider-ad-go/ad/v4/gposecurity/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GpoSecurityPasswordPoliciesOutputReference interface {
	cdktf.ComplexObject
	ClearTextPassword() *string
	SetClearTextPassword(val *string)
	ClearTextPasswordInput() *string
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
	InternalValue() *GpoSecurityPasswordPolicies
	SetInternalValue(val *GpoSecurityPasswordPolicies)
	MaximumPasswordAge() *string
	SetMaximumPasswordAge(val *string)
	MaximumPasswordAgeInput() *string
	MinimumPasswordAge() *string
	SetMinimumPasswordAge(val *string)
	MinimumPasswordAgeInput() *string
	MinimumPasswordLength() *string
	SetMinimumPasswordLength(val *string)
	MinimumPasswordLengthInput() *string
	PasswordComplexity() *string
	SetPasswordComplexity(val *string)
	PasswordComplexityInput() *string
	PasswordHistorySize() *string
	SetPasswordHistorySize(val *string)
	PasswordHistorySizeInput() *string
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
	ResetClearTextPassword()
	ResetMaximumPasswordAge()
	ResetMinimumPasswordAge()
	ResetMinimumPasswordLength()
	ResetPasswordComplexity()
	ResetPasswordHistorySize()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GpoSecurityPasswordPoliciesOutputReference
type jsiiProxy_GpoSecurityPasswordPoliciesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GpoSecurityPasswordPoliciesOutputReference) ClearTextPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clearTextPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GpoSecurityPasswordPoliciesOutputReference) ClearTextPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clearTextPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GpoSecurityPasswordPoliciesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GpoSecurityPasswordPoliciesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GpoSecurityPasswordPoliciesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GpoSecurityPasswordPoliciesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GpoSecurityPasswordPoliciesOutputReference) InternalValue() *GpoSecurityPasswordPolicies {
	var returns *GpoSecurityPasswordPolicies
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GpoSecurityPasswordPoliciesOutputReference) MaximumPasswordAge() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maximumPasswordAge",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GpoSecurityPasswordPoliciesOutputReference) MaximumPasswordAgeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maximumPasswordAgeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GpoSecurityPasswordPoliciesOutputReference) MinimumPasswordAge() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minimumPasswordAge",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GpoSecurityPasswordPoliciesOutputReference) MinimumPasswordAgeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minimumPasswordAgeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GpoSecurityPasswordPoliciesOutputReference) MinimumPasswordLength() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minimumPasswordLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GpoSecurityPasswordPoliciesOutputReference) MinimumPasswordLengthInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minimumPasswordLengthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GpoSecurityPasswordPoliciesOutputReference) PasswordComplexity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordComplexity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GpoSecurityPasswordPoliciesOutputReference) PasswordComplexityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordComplexityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GpoSecurityPasswordPoliciesOutputReference) PasswordHistorySize() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordHistorySize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GpoSecurityPasswordPoliciesOutputReference) PasswordHistorySizeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordHistorySizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GpoSecurityPasswordPoliciesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GpoSecurityPasswordPoliciesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGpoSecurityPasswordPoliciesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GpoSecurityPasswordPoliciesOutputReference {
	_init_.Initialize()

	if err := validateNewGpoSecurityPasswordPoliciesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GpoSecurityPasswordPoliciesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-ad.gpoSecurity.GpoSecurityPasswordPoliciesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGpoSecurityPasswordPoliciesOutputReference_Override(g GpoSecurityPasswordPoliciesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-ad.gpoSecurity.GpoSecurityPasswordPoliciesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GpoSecurityPasswordPoliciesOutputReference)SetClearTextPassword(val *string) {
	if err := j.validateSetClearTextPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clearTextPassword",
		val,
	)
}

func (j *jsiiProxy_GpoSecurityPasswordPoliciesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GpoSecurityPasswordPoliciesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GpoSecurityPasswordPoliciesOutputReference)SetInternalValue(val *GpoSecurityPasswordPolicies) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GpoSecurityPasswordPoliciesOutputReference)SetMaximumPasswordAge(val *string) {
	if err := j.validateSetMaximumPasswordAgeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maximumPasswordAge",
		val,
	)
}

func (j *jsiiProxy_GpoSecurityPasswordPoliciesOutputReference)SetMinimumPasswordAge(val *string) {
	if err := j.validateSetMinimumPasswordAgeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minimumPasswordAge",
		val,
	)
}

func (j *jsiiProxy_GpoSecurityPasswordPoliciesOutputReference)SetMinimumPasswordLength(val *string) {
	if err := j.validateSetMinimumPasswordLengthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minimumPasswordLength",
		val,
	)
}

func (j *jsiiProxy_GpoSecurityPasswordPoliciesOutputReference)SetPasswordComplexity(val *string) {
	if err := j.validateSetPasswordComplexityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"passwordComplexity",
		val,
	)
}

func (j *jsiiProxy_GpoSecurityPasswordPoliciesOutputReference)SetPasswordHistorySize(val *string) {
	if err := j.validateSetPasswordHistorySizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"passwordHistorySize",
		val,
	)
}

func (j *jsiiProxy_GpoSecurityPasswordPoliciesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GpoSecurityPasswordPoliciesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GpoSecurityPasswordPoliciesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GpoSecurityPasswordPoliciesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GpoSecurityPasswordPoliciesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GpoSecurityPasswordPoliciesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GpoSecurityPasswordPoliciesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GpoSecurityPasswordPoliciesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GpoSecurityPasswordPoliciesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GpoSecurityPasswordPoliciesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GpoSecurityPasswordPoliciesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GpoSecurityPasswordPoliciesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GpoSecurityPasswordPoliciesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GpoSecurityPasswordPoliciesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GpoSecurityPasswordPoliciesOutputReference) ResetClearTextPassword() {
	_jsii_.InvokeVoid(
		g,
		"resetClearTextPassword",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GpoSecurityPasswordPoliciesOutputReference) ResetMaximumPasswordAge() {
	_jsii_.InvokeVoid(
		g,
		"resetMaximumPasswordAge",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GpoSecurityPasswordPoliciesOutputReference) ResetMinimumPasswordAge() {
	_jsii_.InvokeVoid(
		g,
		"resetMinimumPasswordAge",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GpoSecurityPasswordPoliciesOutputReference) ResetMinimumPasswordLength() {
	_jsii_.InvokeVoid(
		g,
		"resetMinimumPasswordLength",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GpoSecurityPasswordPoliciesOutputReference) ResetPasswordComplexity() {
	_jsii_.InvokeVoid(
		g,
		"resetPasswordComplexity",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GpoSecurityPasswordPoliciesOutputReference) ResetPasswordHistorySize() {
	_jsii_.InvokeVoid(
		g,
		"resetPasswordHistorySize",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GpoSecurityPasswordPoliciesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GpoSecurityPasswordPoliciesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

