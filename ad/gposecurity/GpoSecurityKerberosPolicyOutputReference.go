package gposecurity

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-ad-go/ad/v5/jsii"

	"github.com/cdktf/cdktf-provider-ad-go/ad/v5/gposecurity/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GpoSecurityKerberosPolicyOutputReference interface {
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
	// Experimental.
	Fqn() *string
	InternalValue() *GpoSecurityKerberosPolicy
	SetInternalValue(val *GpoSecurityKerberosPolicy)
	MaxClockSkew() *string
	SetMaxClockSkew(val *string)
	MaxClockSkewInput() *string
	MaxRenewAge() *string
	SetMaxRenewAge(val *string)
	MaxRenewAgeInput() *string
	MaxServiceAge() *string
	SetMaxServiceAge(val *string)
	MaxServiceAgeInput() *string
	MaxTicketAge() *string
	SetMaxTicketAge(val *string)
	MaxTicketAgeInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TicketValidateClient() *string
	SetTicketValidateClient(val *string)
	TicketValidateClientInput() *string
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
	ResetMaxClockSkew()
	ResetMaxRenewAge()
	ResetMaxServiceAge()
	ResetMaxTicketAge()
	ResetTicketValidateClient()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GpoSecurityKerberosPolicyOutputReference
type jsiiProxy_GpoSecurityKerberosPolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GpoSecurityKerberosPolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GpoSecurityKerberosPolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GpoSecurityKerberosPolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GpoSecurityKerberosPolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GpoSecurityKerberosPolicyOutputReference) InternalValue() *GpoSecurityKerberosPolicy {
	var returns *GpoSecurityKerberosPolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GpoSecurityKerberosPolicyOutputReference) MaxClockSkew() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxClockSkew",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GpoSecurityKerberosPolicyOutputReference) MaxClockSkewInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxClockSkewInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GpoSecurityKerberosPolicyOutputReference) MaxRenewAge() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxRenewAge",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GpoSecurityKerberosPolicyOutputReference) MaxRenewAgeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxRenewAgeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GpoSecurityKerberosPolicyOutputReference) MaxServiceAge() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxServiceAge",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GpoSecurityKerberosPolicyOutputReference) MaxServiceAgeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxServiceAgeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GpoSecurityKerberosPolicyOutputReference) MaxTicketAge() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxTicketAge",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GpoSecurityKerberosPolicyOutputReference) MaxTicketAgeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxTicketAgeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GpoSecurityKerberosPolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GpoSecurityKerberosPolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GpoSecurityKerberosPolicyOutputReference) TicketValidateClient() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ticketValidateClient",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GpoSecurityKerberosPolicyOutputReference) TicketValidateClientInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ticketValidateClientInput",
		&returns,
	)
	return returns
}


func NewGpoSecurityKerberosPolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GpoSecurityKerberosPolicyOutputReference {
	_init_.Initialize()

	if err := validateNewGpoSecurityKerberosPolicyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GpoSecurityKerberosPolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-ad.gpoSecurity.GpoSecurityKerberosPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGpoSecurityKerberosPolicyOutputReference_Override(g GpoSecurityKerberosPolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-ad.gpoSecurity.GpoSecurityKerberosPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GpoSecurityKerberosPolicyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GpoSecurityKerberosPolicyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GpoSecurityKerberosPolicyOutputReference)SetInternalValue(val *GpoSecurityKerberosPolicy) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GpoSecurityKerberosPolicyOutputReference)SetMaxClockSkew(val *string) {
	if err := j.validateSetMaxClockSkewParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxClockSkew",
		val,
	)
}

func (j *jsiiProxy_GpoSecurityKerberosPolicyOutputReference)SetMaxRenewAge(val *string) {
	if err := j.validateSetMaxRenewAgeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxRenewAge",
		val,
	)
}

func (j *jsiiProxy_GpoSecurityKerberosPolicyOutputReference)SetMaxServiceAge(val *string) {
	if err := j.validateSetMaxServiceAgeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxServiceAge",
		val,
	)
}

func (j *jsiiProxy_GpoSecurityKerberosPolicyOutputReference)SetMaxTicketAge(val *string) {
	if err := j.validateSetMaxTicketAgeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxTicketAge",
		val,
	)
}

func (j *jsiiProxy_GpoSecurityKerberosPolicyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GpoSecurityKerberosPolicyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GpoSecurityKerberosPolicyOutputReference)SetTicketValidateClient(val *string) {
	if err := j.validateSetTicketValidateClientParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ticketValidateClient",
		val,
	)
}

func (g *jsiiProxy_GpoSecurityKerberosPolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GpoSecurityKerberosPolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GpoSecurityKerberosPolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GpoSecurityKerberosPolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GpoSecurityKerberosPolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GpoSecurityKerberosPolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GpoSecurityKerberosPolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GpoSecurityKerberosPolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GpoSecurityKerberosPolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GpoSecurityKerberosPolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GpoSecurityKerberosPolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GpoSecurityKerberosPolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GpoSecurityKerberosPolicyOutputReference) ResetMaxClockSkew() {
	_jsii_.InvokeVoid(
		g,
		"resetMaxClockSkew",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GpoSecurityKerberosPolicyOutputReference) ResetMaxRenewAge() {
	_jsii_.InvokeVoid(
		g,
		"resetMaxRenewAge",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GpoSecurityKerberosPolicyOutputReference) ResetMaxServiceAge() {
	_jsii_.InvokeVoid(
		g,
		"resetMaxServiceAge",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GpoSecurityKerberosPolicyOutputReference) ResetMaxTicketAge() {
	_jsii_.InvokeVoid(
		g,
		"resetMaxTicketAge",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GpoSecurityKerberosPolicyOutputReference) ResetTicketValidateClient() {
	_jsii_.InvokeVoid(
		g,
		"resetTicketValidateClient",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GpoSecurityKerberosPolicyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GpoSecurityKerberosPolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

