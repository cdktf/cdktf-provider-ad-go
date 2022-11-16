//go:build no_runtime_type_checking

package gposecurity

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_GpoSecurityFilesystemList) validateGetParameters(index *float64) error {
	return nil
}

func (g *jsiiProxy_GpoSecurityFilesystemList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_GpoSecurityFilesystemList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_GpoSecurityFilesystemList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_GpoSecurityFilesystemList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_GpoSecurityFilesystemList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewGpoSecurityFilesystemListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

