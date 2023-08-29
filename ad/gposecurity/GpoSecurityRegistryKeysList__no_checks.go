// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package gposecurity

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_GpoSecurityRegistryKeysList) validateGetParameters(index *float64) error {
	return nil
}

func (g *jsiiProxy_GpoSecurityRegistryKeysList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_GpoSecurityRegistryKeysList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_GpoSecurityRegistryKeysList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_GpoSecurityRegistryKeysList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_GpoSecurityRegistryKeysList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewGpoSecurityRegistryKeysListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

