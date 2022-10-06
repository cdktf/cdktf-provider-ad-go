//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package gposecurity

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_GpoSecuritySystemServicesList) validateGetParameters(index *float64) error {
	return nil
}

func (g *jsiiProxy_GpoSecuritySystemServicesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_GpoSecuritySystemServicesList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_GpoSecuritySystemServicesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_GpoSecuritySystemServicesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_GpoSecuritySystemServicesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewGpoSecuritySystemServicesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

