// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gposecurity

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-ad.gpoSecurity.GpoSecurity",
		reflect.TypeOf((*GpoSecurity)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "accountLockout", GoGetter: "AccountLockout"},
			_jsii_.MemberProperty{JsiiProperty: "accountLockoutInput", GoGetter: "AccountLockoutInput"},
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "applicationLog", GoGetter: "ApplicationLog"},
			_jsii_.MemberProperty{JsiiProperty: "applicationLogInput", GoGetter: "ApplicationLogInput"},
			_jsii_.MemberProperty{JsiiProperty: "auditLog", GoGetter: "AuditLog"},
			_jsii_.MemberProperty{JsiiProperty: "auditLogInput", GoGetter: "AuditLogInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "eventAudit", GoGetter: "EventAudit"},
			_jsii_.MemberProperty{JsiiProperty: "eventAuditInput", GoGetter: "EventAuditInput"},
			_jsii_.MemberProperty{JsiiProperty: "filesystem", GoGetter: "Filesystem"},
			_jsii_.MemberProperty{JsiiProperty: "filesystemInput", GoGetter: "FilesystemInput"},
			_jsii_.MemberProperty{JsiiProperty: "forEach", GoGetter: "ForEach"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "gpoContainer", GoGetter: "GpoContainer"},
			_jsii_.MemberProperty{JsiiProperty: "gpoContainerInput", GoGetter: "GpoContainerInput"},
			_jsii_.MemberMethod{JsiiMethod: "hasResourceMove", GoMethod: "HasResourceMove"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberMethod{JsiiMethod: "importFrom", GoMethod: "ImportFrom"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "kerberosPolicy", GoGetter: "KerberosPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "kerberosPolicyInput", GoGetter: "KerberosPolicyInput"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberMethod{JsiiMethod: "moveFromId", GoMethod: "MoveFromId"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberMethod{JsiiMethod: "moveToId", GoMethod: "MoveToId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "passwordPolicies", GoGetter: "PasswordPolicies"},
			_jsii_.MemberProperty{JsiiProperty: "passwordPoliciesInput", GoGetter: "PasswordPoliciesInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberMethod{JsiiMethod: "putAccountLockout", GoMethod: "PutAccountLockout"},
			_jsii_.MemberMethod{JsiiMethod: "putApplicationLog", GoMethod: "PutApplicationLog"},
			_jsii_.MemberMethod{JsiiMethod: "putAuditLog", GoMethod: "PutAuditLog"},
			_jsii_.MemberMethod{JsiiMethod: "putEventAudit", GoMethod: "PutEventAudit"},
			_jsii_.MemberMethod{JsiiMethod: "putFilesystem", GoMethod: "PutFilesystem"},
			_jsii_.MemberMethod{JsiiMethod: "putKerberosPolicy", GoMethod: "PutKerberosPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "putPasswordPolicies", GoMethod: "PutPasswordPolicies"},
			_jsii_.MemberMethod{JsiiMethod: "putRegistryKeys", GoMethod: "PutRegistryKeys"},
			_jsii_.MemberMethod{JsiiMethod: "putRegistryValues", GoMethod: "PutRegistryValues"},
			_jsii_.MemberMethod{JsiiMethod: "putRestrictedGroups", GoMethod: "PutRestrictedGroups"},
			_jsii_.MemberMethod{JsiiMethod: "putSystemLog", GoMethod: "PutSystemLog"},
			_jsii_.MemberMethod{JsiiMethod: "putSystemServices", GoMethod: "PutSystemServices"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberProperty{JsiiProperty: "registryKeys", GoGetter: "RegistryKeys"},
			_jsii_.MemberProperty{JsiiProperty: "registryKeysInput", GoGetter: "RegistryKeysInput"},
			_jsii_.MemberProperty{JsiiProperty: "registryValues", GoGetter: "RegistryValues"},
			_jsii_.MemberProperty{JsiiProperty: "registryValuesInput", GoGetter: "RegistryValuesInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetAccountLockout", GoMethod: "ResetAccountLockout"},
			_jsii_.MemberMethod{JsiiMethod: "resetApplicationLog", GoMethod: "ResetApplicationLog"},
			_jsii_.MemberMethod{JsiiMethod: "resetAuditLog", GoMethod: "ResetAuditLog"},
			_jsii_.MemberMethod{JsiiMethod: "resetEventAudit", GoMethod: "ResetEventAudit"},
			_jsii_.MemberMethod{JsiiMethod: "resetFilesystem", GoMethod: "ResetFilesystem"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetKerberosPolicy", GoMethod: "ResetKerberosPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPasswordPolicies", GoMethod: "ResetPasswordPolicies"},
			_jsii_.MemberMethod{JsiiMethod: "resetRegistryKeys", GoMethod: "ResetRegistryKeys"},
			_jsii_.MemberMethod{JsiiMethod: "resetRegistryValues", GoMethod: "ResetRegistryValues"},
			_jsii_.MemberMethod{JsiiMethod: "resetRestrictedGroups", GoMethod: "ResetRestrictedGroups"},
			_jsii_.MemberMethod{JsiiMethod: "resetSystemLog", GoMethod: "ResetSystemLog"},
			_jsii_.MemberMethod{JsiiMethod: "resetSystemServices", GoMethod: "ResetSystemServices"},
			_jsii_.MemberProperty{JsiiProperty: "restrictedGroups", GoGetter: "RestrictedGroups"},
			_jsii_.MemberProperty{JsiiProperty: "restrictedGroupsInput", GoGetter: "RestrictedGroupsInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "systemLog", GoGetter: "SystemLog"},
			_jsii_.MemberProperty{JsiiProperty: "systemLogInput", GoGetter: "SystemLogInput"},
			_jsii_.MemberProperty{JsiiProperty: "systemServices", GoGetter: "SystemServices"},
			_jsii_.MemberProperty{JsiiProperty: "systemServicesInput", GoGetter: "SystemServicesInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_GpoSecurity{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-ad.gpoSecurity.GpoSecurityAccountLockout",
		reflect.TypeOf((*GpoSecurityAccountLockout)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-ad.gpoSecurity.GpoSecurityAccountLockoutOutputReference",
		reflect.TypeOf((*GpoSecurityAccountLockoutOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "forceLogoffWhenHourExpire", GoGetter: "ForceLogoffWhenHourExpire"},
			_jsii_.MemberProperty{JsiiProperty: "forceLogoffWhenHourExpireInput", GoGetter: "ForceLogoffWhenHourExpireInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lockoutBadCount", GoGetter: "LockoutBadCount"},
			_jsii_.MemberProperty{JsiiProperty: "lockoutBadCountInput", GoGetter: "LockoutBadCountInput"},
			_jsii_.MemberProperty{JsiiProperty: "lockoutDuration", GoGetter: "LockoutDuration"},
			_jsii_.MemberProperty{JsiiProperty: "lockoutDurationInput", GoGetter: "LockoutDurationInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetForceLogoffWhenHourExpire", GoMethod: "ResetForceLogoffWhenHourExpire"},
			_jsii_.MemberMethod{JsiiMethod: "resetLockoutBadCount", GoMethod: "ResetLockoutBadCount"},
			_jsii_.MemberProperty{JsiiProperty: "resetLockoutCount", GoGetter: "ResetLockoutCount"},
			_jsii_.MemberProperty{JsiiProperty: "resetLockoutCountInput", GoGetter: "ResetLockoutCountInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetLockoutDuration", GoMethod: "ResetLockoutDuration"},
			_jsii_.MemberMethod{JsiiMethod: "resetResetLockoutCount", GoMethod: "ResetResetLockoutCount"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GpoSecurityAccountLockoutOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-ad.gpoSecurity.GpoSecurityApplicationLog",
		reflect.TypeOf((*GpoSecurityApplicationLog)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-ad.gpoSecurity.GpoSecurityApplicationLogOutputReference",
		reflect.TypeOf((*GpoSecurityApplicationLogOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "auditLogRetentionPeriod", GoGetter: "AuditLogRetentionPeriod"},
			_jsii_.MemberProperty{JsiiProperty: "auditLogRetentionPeriodInput", GoGetter: "AuditLogRetentionPeriodInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "maximumLogSize", GoGetter: "MaximumLogSize"},
			_jsii_.MemberProperty{JsiiProperty: "maximumLogSizeInput", GoGetter: "MaximumLogSizeInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetAuditLogRetentionPeriod", GoMethod: "ResetAuditLogRetentionPeriod"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaximumLogSize", GoMethod: "ResetMaximumLogSize"},
			_jsii_.MemberMethod{JsiiMethod: "resetRestrictGuestAccess", GoMethod: "ResetRestrictGuestAccess"},
			_jsii_.MemberMethod{JsiiMethod: "resetRetentionDays", GoMethod: "ResetRetentionDays"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "restrictGuestAccess", GoGetter: "RestrictGuestAccess"},
			_jsii_.MemberProperty{JsiiProperty: "restrictGuestAccessInput", GoGetter: "RestrictGuestAccessInput"},
			_jsii_.MemberProperty{JsiiProperty: "retentionDays", GoGetter: "RetentionDays"},
			_jsii_.MemberProperty{JsiiProperty: "retentionDaysInput", GoGetter: "RetentionDaysInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GpoSecurityApplicationLogOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-ad.gpoSecurity.GpoSecurityAuditLog",
		reflect.TypeOf((*GpoSecurityAuditLog)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-ad.gpoSecurity.GpoSecurityAuditLogOutputReference",
		reflect.TypeOf((*GpoSecurityAuditLogOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "auditLogRetentionPeriod", GoGetter: "AuditLogRetentionPeriod"},
			_jsii_.MemberProperty{JsiiProperty: "auditLogRetentionPeriodInput", GoGetter: "AuditLogRetentionPeriodInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "maximumLogSize", GoGetter: "MaximumLogSize"},
			_jsii_.MemberProperty{JsiiProperty: "maximumLogSizeInput", GoGetter: "MaximumLogSizeInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetAuditLogRetentionPeriod", GoMethod: "ResetAuditLogRetentionPeriod"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaximumLogSize", GoMethod: "ResetMaximumLogSize"},
			_jsii_.MemberMethod{JsiiMethod: "resetRestrictGuestAccess", GoMethod: "ResetRestrictGuestAccess"},
			_jsii_.MemberMethod{JsiiMethod: "resetRetentionDays", GoMethod: "ResetRetentionDays"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "restrictGuestAccess", GoGetter: "RestrictGuestAccess"},
			_jsii_.MemberProperty{JsiiProperty: "restrictGuestAccessInput", GoGetter: "RestrictGuestAccessInput"},
			_jsii_.MemberProperty{JsiiProperty: "retentionDays", GoGetter: "RetentionDays"},
			_jsii_.MemberProperty{JsiiProperty: "retentionDaysInput", GoGetter: "RetentionDaysInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GpoSecurityAuditLogOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-ad.gpoSecurity.GpoSecurityConfig",
		reflect.TypeOf((*GpoSecurityConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-ad.gpoSecurity.GpoSecurityEventAudit",
		reflect.TypeOf((*GpoSecurityEventAudit)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-ad.gpoSecurity.GpoSecurityEventAuditOutputReference",
		reflect.TypeOf((*GpoSecurityEventAuditOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "auditAccountLogon", GoGetter: "AuditAccountLogon"},
			_jsii_.MemberProperty{JsiiProperty: "auditAccountLogonInput", GoGetter: "AuditAccountLogonInput"},
			_jsii_.MemberProperty{JsiiProperty: "auditAccountManage", GoGetter: "AuditAccountManage"},
			_jsii_.MemberProperty{JsiiProperty: "auditAccountManageInput", GoGetter: "AuditAccountManageInput"},
			_jsii_.MemberProperty{JsiiProperty: "auditDsAccess", GoGetter: "AuditDsAccess"},
			_jsii_.MemberProperty{JsiiProperty: "auditDsAccessInput", GoGetter: "AuditDsAccessInput"},
			_jsii_.MemberProperty{JsiiProperty: "auditLogonEvents", GoGetter: "AuditLogonEvents"},
			_jsii_.MemberProperty{JsiiProperty: "auditLogonEventsInput", GoGetter: "AuditLogonEventsInput"},
			_jsii_.MemberProperty{JsiiProperty: "auditObjectAccess", GoGetter: "AuditObjectAccess"},
			_jsii_.MemberProperty{JsiiProperty: "auditObjectAccessInput", GoGetter: "AuditObjectAccessInput"},
			_jsii_.MemberProperty{JsiiProperty: "auditPolicyChange", GoGetter: "AuditPolicyChange"},
			_jsii_.MemberProperty{JsiiProperty: "auditPolicyChangeInput", GoGetter: "AuditPolicyChangeInput"},
			_jsii_.MemberProperty{JsiiProperty: "auditPrivilegeUse", GoGetter: "AuditPrivilegeUse"},
			_jsii_.MemberProperty{JsiiProperty: "auditPrivilegeUseInput", GoGetter: "AuditPrivilegeUseInput"},
			_jsii_.MemberProperty{JsiiProperty: "auditProcessTracking", GoGetter: "AuditProcessTracking"},
			_jsii_.MemberProperty{JsiiProperty: "auditProcessTrackingInput", GoGetter: "AuditProcessTrackingInput"},
			_jsii_.MemberProperty{JsiiProperty: "auditSystemEvents", GoGetter: "AuditSystemEvents"},
			_jsii_.MemberProperty{JsiiProperty: "auditSystemEventsInput", GoGetter: "AuditSystemEventsInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetAuditAccountLogon", GoMethod: "ResetAuditAccountLogon"},
			_jsii_.MemberMethod{JsiiMethod: "resetAuditAccountManage", GoMethod: "ResetAuditAccountManage"},
			_jsii_.MemberMethod{JsiiMethod: "resetAuditDsAccess", GoMethod: "ResetAuditDsAccess"},
			_jsii_.MemberMethod{JsiiMethod: "resetAuditLogonEvents", GoMethod: "ResetAuditLogonEvents"},
			_jsii_.MemberMethod{JsiiMethod: "resetAuditObjectAccess", GoMethod: "ResetAuditObjectAccess"},
			_jsii_.MemberMethod{JsiiMethod: "resetAuditPolicyChange", GoMethod: "ResetAuditPolicyChange"},
			_jsii_.MemberMethod{JsiiMethod: "resetAuditPrivilegeUse", GoMethod: "ResetAuditPrivilegeUse"},
			_jsii_.MemberMethod{JsiiMethod: "resetAuditProcessTracking", GoMethod: "ResetAuditProcessTracking"},
			_jsii_.MemberMethod{JsiiMethod: "resetAuditSystemEvents", GoMethod: "ResetAuditSystemEvents"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GpoSecurityEventAuditOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-ad.gpoSecurity.GpoSecurityFilesystem",
		reflect.TypeOf((*GpoSecurityFilesystem)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-ad.gpoSecurity.GpoSecurityFilesystemList",
		reflect.TypeOf((*GpoSecurityFilesystemList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_GpoSecurityFilesystemList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-ad.gpoSecurity.GpoSecurityFilesystemOutputReference",
		reflect.TypeOf((*GpoSecurityFilesystemOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "acl", GoGetter: "Acl"},
			_jsii_.MemberProperty{JsiiProperty: "aclInput", GoGetter: "AclInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "path", GoGetter: "Path"},
			_jsii_.MemberProperty{JsiiProperty: "pathInput", GoGetter: "PathInput"},
			_jsii_.MemberProperty{JsiiProperty: "propagationMode", GoGetter: "PropagationMode"},
			_jsii_.MemberProperty{JsiiProperty: "propagationModeInput", GoGetter: "PropagationModeInput"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GpoSecurityFilesystemOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-ad.gpoSecurity.GpoSecurityKerberosPolicy",
		reflect.TypeOf((*GpoSecurityKerberosPolicy)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-ad.gpoSecurity.GpoSecurityKerberosPolicyOutputReference",
		reflect.TypeOf((*GpoSecurityKerberosPolicyOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "maxClockSkew", GoGetter: "MaxClockSkew"},
			_jsii_.MemberProperty{JsiiProperty: "maxClockSkewInput", GoGetter: "MaxClockSkewInput"},
			_jsii_.MemberProperty{JsiiProperty: "maxRenewAge", GoGetter: "MaxRenewAge"},
			_jsii_.MemberProperty{JsiiProperty: "maxRenewAgeInput", GoGetter: "MaxRenewAgeInput"},
			_jsii_.MemberProperty{JsiiProperty: "maxServiceAge", GoGetter: "MaxServiceAge"},
			_jsii_.MemberProperty{JsiiProperty: "maxServiceAgeInput", GoGetter: "MaxServiceAgeInput"},
			_jsii_.MemberProperty{JsiiProperty: "maxTicketAge", GoGetter: "MaxTicketAge"},
			_jsii_.MemberProperty{JsiiProperty: "maxTicketAgeInput", GoGetter: "MaxTicketAgeInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaxClockSkew", GoMethod: "ResetMaxClockSkew"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaxRenewAge", GoMethod: "ResetMaxRenewAge"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaxServiceAge", GoMethod: "ResetMaxServiceAge"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaxTicketAge", GoMethod: "ResetMaxTicketAge"},
			_jsii_.MemberMethod{JsiiMethod: "resetTicketValidateClient", GoMethod: "ResetTicketValidateClient"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "ticketValidateClient", GoGetter: "TicketValidateClient"},
			_jsii_.MemberProperty{JsiiProperty: "ticketValidateClientInput", GoGetter: "TicketValidateClientInput"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GpoSecurityKerberosPolicyOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-ad.gpoSecurity.GpoSecurityPasswordPolicies",
		reflect.TypeOf((*GpoSecurityPasswordPolicies)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-ad.gpoSecurity.GpoSecurityPasswordPoliciesOutputReference",
		reflect.TypeOf((*GpoSecurityPasswordPoliciesOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "clearTextPassword", GoGetter: "ClearTextPassword"},
			_jsii_.MemberProperty{JsiiProperty: "clearTextPasswordInput", GoGetter: "ClearTextPasswordInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "maximumPasswordAge", GoGetter: "MaximumPasswordAge"},
			_jsii_.MemberProperty{JsiiProperty: "maximumPasswordAgeInput", GoGetter: "MaximumPasswordAgeInput"},
			_jsii_.MemberProperty{JsiiProperty: "minimumPasswordAge", GoGetter: "MinimumPasswordAge"},
			_jsii_.MemberProperty{JsiiProperty: "minimumPasswordAgeInput", GoGetter: "MinimumPasswordAgeInput"},
			_jsii_.MemberProperty{JsiiProperty: "minimumPasswordLength", GoGetter: "MinimumPasswordLength"},
			_jsii_.MemberProperty{JsiiProperty: "minimumPasswordLengthInput", GoGetter: "MinimumPasswordLengthInput"},
			_jsii_.MemberProperty{JsiiProperty: "passwordComplexity", GoGetter: "PasswordComplexity"},
			_jsii_.MemberProperty{JsiiProperty: "passwordComplexityInput", GoGetter: "PasswordComplexityInput"},
			_jsii_.MemberProperty{JsiiProperty: "passwordHistorySize", GoGetter: "PasswordHistorySize"},
			_jsii_.MemberProperty{JsiiProperty: "passwordHistorySizeInput", GoGetter: "PasswordHistorySizeInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetClearTextPassword", GoMethod: "ResetClearTextPassword"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaximumPasswordAge", GoMethod: "ResetMaximumPasswordAge"},
			_jsii_.MemberMethod{JsiiMethod: "resetMinimumPasswordAge", GoMethod: "ResetMinimumPasswordAge"},
			_jsii_.MemberMethod{JsiiMethod: "resetMinimumPasswordLength", GoMethod: "ResetMinimumPasswordLength"},
			_jsii_.MemberMethod{JsiiMethod: "resetPasswordComplexity", GoMethod: "ResetPasswordComplexity"},
			_jsii_.MemberMethod{JsiiMethod: "resetPasswordHistorySize", GoMethod: "ResetPasswordHistorySize"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GpoSecurityPasswordPoliciesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-ad.gpoSecurity.GpoSecurityRegistryKeys",
		reflect.TypeOf((*GpoSecurityRegistryKeys)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-ad.gpoSecurity.GpoSecurityRegistryKeysList",
		reflect.TypeOf((*GpoSecurityRegistryKeysList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_GpoSecurityRegistryKeysList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-ad.gpoSecurity.GpoSecurityRegistryKeysOutputReference",
		reflect.TypeOf((*GpoSecurityRegistryKeysOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "acl", GoGetter: "Acl"},
			_jsii_.MemberProperty{JsiiProperty: "aclInput", GoGetter: "AclInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "keyName", GoGetter: "KeyName"},
			_jsii_.MemberProperty{JsiiProperty: "keyNameInput", GoGetter: "KeyNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "propagationMode", GoGetter: "PropagationMode"},
			_jsii_.MemberProperty{JsiiProperty: "propagationModeInput", GoGetter: "PropagationModeInput"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GpoSecurityRegistryKeysOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-ad.gpoSecurity.GpoSecurityRegistryValues",
		reflect.TypeOf((*GpoSecurityRegistryValues)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-ad.gpoSecurity.GpoSecurityRegistryValuesList",
		reflect.TypeOf((*GpoSecurityRegistryValuesList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_GpoSecurityRegistryValuesList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-ad.gpoSecurity.GpoSecurityRegistryValuesOutputReference",
		reflect.TypeOf((*GpoSecurityRegistryValuesOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "keyName", GoGetter: "KeyName"},
			_jsii_.MemberProperty{JsiiProperty: "keyNameInput", GoGetter: "KeyNameInput"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
			_jsii_.MemberProperty{JsiiProperty: "valueInput", GoGetter: "ValueInput"},
			_jsii_.MemberProperty{JsiiProperty: "valueType", GoGetter: "ValueType"},
			_jsii_.MemberProperty{JsiiProperty: "valueTypeInput", GoGetter: "ValueTypeInput"},
		},
		func() interface{} {
			j := jsiiProxy_GpoSecurityRegistryValuesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-ad.gpoSecurity.GpoSecurityRestrictedGroups",
		reflect.TypeOf((*GpoSecurityRestrictedGroups)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-ad.gpoSecurity.GpoSecurityRestrictedGroupsList",
		reflect.TypeOf((*GpoSecurityRestrictedGroupsList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_GpoSecurityRestrictedGroupsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-ad.gpoSecurity.GpoSecurityRestrictedGroupsOutputReference",
		reflect.TypeOf((*GpoSecurityRestrictedGroupsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "groupMemberof", GoGetter: "GroupMemberof"},
			_jsii_.MemberProperty{JsiiProperty: "groupMemberofInput", GoGetter: "GroupMemberofInput"},
			_jsii_.MemberProperty{JsiiProperty: "groupMembers", GoGetter: "GroupMembers"},
			_jsii_.MemberProperty{JsiiProperty: "groupMembersInput", GoGetter: "GroupMembersInput"},
			_jsii_.MemberProperty{JsiiProperty: "groupName", GoGetter: "GroupName"},
			_jsii_.MemberProperty{JsiiProperty: "groupNameInput", GoGetter: "GroupNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GpoSecurityRestrictedGroupsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-ad.gpoSecurity.GpoSecuritySystemLog",
		reflect.TypeOf((*GpoSecuritySystemLog)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-ad.gpoSecurity.GpoSecuritySystemLogOutputReference",
		reflect.TypeOf((*GpoSecuritySystemLogOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "auditLogRetentionPeriod", GoGetter: "AuditLogRetentionPeriod"},
			_jsii_.MemberProperty{JsiiProperty: "auditLogRetentionPeriodInput", GoGetter: "AuditLogRetentionPeriodInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "maximumLogSize", GoGetter: "MaximumLogSize"},
			_jsii_.MemberProperty{JsiiProperty: "maximumLogSizeInput", GoGetter: "MaximumLogSizeInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetAuditLogRetentionPeriod", GoMethod: "ResetAuditLogRetentionPeriod"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaximumLogSize", GoMethod: "ResetMaximumLogSize"},
			_jsii_.MemberMethod{JsiiMethod: "resetRestrictGuestAccess", GoMethod: "ResetRestrictGuestAccess"},
			_jsii_.MemberMethod{JsiiMethod: "resetRetentionDays", GoMethod: "ResetRetentionDays"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "restrictGuestAccess", GoGetter: "RestrictGuestAccess"},
			_jsii_.MemberProperty{JsiiProperty: "restrictGuestAccessInput", GoGetter: "RestrictGuestAccessInput"},
			_jsii_.MemberProperty{JsiiProperty: "retentionDays", GoGetter: "RetentionDays"},
			_jsii_.MemberProperty{JsiiProperty: "retentionDaysInput", GoGetter: "RetentionDaysInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GpoSecuritySystemLogOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-ad.gpoSecurity.GpoSecuritySystemServices",
		reflect.TypeOf((*GpoSecuritySystemServices)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-ad.gpoSecurity.GpoSecuritySystemServicesList",
		reflect.TypeOf((*GpoSecuritySystemServicesList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_GpoSecuritySystemServicesList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-ad.gpoSecurity.GpoSecuritySystemServicesOutputReference",
		reflect.TypeOf((*GpoSecuritySystemServicesOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "acl", GoGetter: "Acl"},
			_jsii_.MemberProperty{JsiiProperty: "aclInput", GoGetter: "AclInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "serviceName", GoGetter: "ServiceName"},
			_jsii_.MemberProperty{JsiiProperty: "serviceNameInput", GoGetter: "ServiceNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "startupMode", GoGetter: "StartupMode"},
			_jsii_.MemberProperty{JsiiProperty: "startupModeInput", GoGetter: "StartupModeInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GpoSecuritySystemServicesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
}
