//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package gposecurity

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_GpoSecurityRegistryValuesList) validateGetParameters(index *float64) error {
	return nil
}

func (g *jsiiProxy_GpoSecurityRegistryValuesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_GpoSecurityRegistryValuesList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_GpoSecurityRegistryValuesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_GpoSecurityRegistryValuesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_GpoSecurityRegistryValuesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewGpoSecurityRegistryValuesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

