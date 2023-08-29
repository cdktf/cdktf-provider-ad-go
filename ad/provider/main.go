// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-ad.provider.AdProvider",
		reflect.TypeOf((*AdProvider)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "alias", GoGetter: "Alias"},
			_jsii_.MemberProperty{JsiiProperty: "aliasInput", GoGetter: "AliasInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "domainController", GoGetter: "DomainController"},
			_jsii_.MemberProperty{JsiiProperty: "domainControllerInput", GoGetter: "DomainControllerInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberProperty{JsiiProperty: "krbConf", GoGetter: "KrbConf"},
			_jsii_.MemberProperty{JsiiProperty: "krbConfInput", GoGetter: "KrbConfInput"},
			_jsii_.MemberProperty{JsiiProperty: "krbKeytab", GoGetter: "KrbKeytab"},
			_jsii_.MemberProperty{JsiiProperty: "krbKeytabInput", GoGetter: "KrbKeytabInput"},
			_jsii_.MemberProperty{JsiiProperty: "krbRealm", GoGetter: "KrbRealm"},
			_jsii_.MemberProperty{JsiiProperty: "krbRealmInput", GoGetter: "KrbRealmInput"},
			_jsii_.MemberProperty{JsiiProperty: "krbSpn", GoGetter: "KrbSpn"},
			_jsii_.MemberProperty{JsiiProperty: "krbSpnInput", GoGetter: "KrbSpnInput"},
			_jsii_.MemberProperty{JsiiProperty: "metaAttributes", GoGetter: "MetaAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetAlias", GoMethod: "ResetAlias"},
			_jsii_.MemberMethod{JsiiMethod: "resetDomainController", GoMethod: "ResetDomainController"},
			_jsii_.MemberMethod{JsiiMethod: "resetKrbConf", GoMethod: "ResetKrbConf"},
			_jsii_.MemberMethod{JsiiMethod: "resetKrbKeytab", GoMethod: "ResetKrbKeytab"},
			_jsii_.MemberMethod{JsiiMethod: "resetKrbRealm", GoMethod: "ResetKrbRealm"},
			_jsii_.MemberMethod{JsiiMethod: "resetKrbSpn", GoMethod: "ResetKrbSpn"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetWinrmInsecure", GoMethod: "ResetWinrmInsecure"},
			_jsii_.MemberMethod{JsiiMethod: "resetWinrmPassCredentials", GoMethod: "ResetWinrmPassCredentials"},
			_jsii_.MemberMethod{JsiiMethod: "resetWinrmPort", GoMethod: "ResetWinrmPort"},
			_jsii_.MemberMethod{JsiiMethod: "resetWinrmProto", GoMethod: "ResetWinrmProto"},
			_jsii_.MemberMethod{JsiiMethod: "resetWinrmUseNtlm", GoMethod: "ResetWinrmUseNtlm"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformProviderSource", GoGetter: "TerraformProviderSource"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "winrmHostname", GoGetter: "WinrmHostname"},
			_jsii_.MemberProperty{JsiiProperty: "winrmHostnameInput", GoGetter: "WinrmHostnameInput"},
			_jsii_.MemberProperty{JsiiProperty: "winrmInsecure", GoGetter: "WinrmInsecure"},
			_jsii_.MemberProperty{JsiiProperty: "winrmInsecureInput", GoGetter: "WinrmInsecureInput"},
			_jsii_.MemberProperty{JsiiProperty: "winrmPassCredentials", GoGetter: "WinrmPassCredentials"},
			_jsii_.MemberProperty{JsiiProperty: "winrmPassCredentialsInput", GoGetter: "WinrmPassCredentialsInput"},
			_jsii_.MemberProperty{JsiiProperty: "winrmPassword", GoGetter: "WinrmPassword"},
			_jsii_.MemberProperty{JsiiProperty: "winrmPasswordInput", GoGetter: "WinrmPasswordInput"},
			_jsii_.MemberProperty{JsiiProperty: "winrmPort", GoGetter: "WinrmPort"},
			_jsii_.MemberProperty{JsiiProperty: "winrmPortInput", GoGetter: "WinrmPortInput"},
			_jsii_.MemberProperty{JsiiProperty: "winrmProto", GoGetter: "WinrmProto"},
			_jsii_.MemberProperty{JsiiProperty: "winrmProtoInput", GoGetter: "WinrmProtoInput"},
			_jsii_.MemberProperty{JsiiProperty: "winrmUseNtlm", GoGetter: "WinrmUseNtlm"},
			_jsii_.MemberProperty{JsiiProperty: "winrmUseNtlmInput", GoGetter: "WinrmUseNtlmInput"},
			_jsii_.MemberProperty{JsiiProperty: "winrmUsername", GoGetter: "WinrmUsername"},
			_jsii_.MemberProperty{JsiiProperty: "winrmUsernameInput", GoGetter: "WinrmUsernameInput"},
		},
		func() interface{} {
			j := jsiiProxy_AdProvider{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformProvider)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-ad.provider.AdProviderConfig",
		reflect.TypeOf((*AdProviderConfig)(nil)).Elem(),
	)
}
