//go:build no_runtime_type_checking

package gposecurity

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_GpoSecurityRestrictedGroupsList) validateGetParameters(index *float64) error {
	return nil
}

func (g *jsiiProxy_GpoSecurityRestrictedGroupsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_GpoSecurityRestrictedGroupsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_GpoSecurityRestrictedGroupsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_GpoSecurityRestrictedGroupsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_GpoSecurityRestrictedGroupsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewGpoSecurityRestrictedGroupsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

