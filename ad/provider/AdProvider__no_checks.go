// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package provider

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AdProvider) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (a *jsiiProxy_AdProvider) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateAdProvider_IsConstructParameters(x interface{}) error {
	return nil
}

func validateAdProvider_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateAdProvider_IsTerraformProviderParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_AdProvider) validateSetWinrmInsecureParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AdProvider) validateSetWinrmPassCredentialsParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AdProvider) validateSetWinrmUseNtlmParameters(val interface{}) error {
	return nil
}

func validateNewAdProviderParameters(scope constructs.Construct, id *string, config *AdProviderConfig) error {
	return nil
}

