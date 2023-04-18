package provider


type AdProviderConfig struct {
	// The hostname of the server we will use to run powershell scripts over WinRM. (Environment variable: AD_HOSTNAME).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/ad/0.4.4/docs#winrm_hostname AdProvider#winrm_hostname}
	WinrmHostname *string `field:"required" json:"winrmHostname" yaml:"winrmHostname"`
	// The password used to authenticate to the server's WinRM service. (Environment variable: AD_PASSWORD).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/ad/0.4.4/docs#winrm_password AdProvider#winrm_password}
	WinrmPassword *string `field:"required" json:"winrmPassword" yaml:"winrmPassword"`
	// The username used to authenticate to the server's WinRM service. (Environment variable: AD_USER).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/ad/0.4.4/docs#winrm_username AdProvider#winrm_username}
	WinrmUsername *string `field:"required" json:"winrmUsername" yaml:"winrmUsername"`
	// Alias name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/ad/0.4.4/docs#alias AdProvider#alias}
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// Use a specific domain controller. (default: none, environment variable: AD_DC).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/ad/0.4.4/docs#domain_controller AdProvider#domain_controller}
	DomainController *string `field:"optional" json:"domainController" yaml:"domainController"`
	// Path to kerberos configuration file. (default: none, environment variable: AD_KRB_CONF).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/ad/0.4.4/docs#krb_conf AdProvider#krb_conf}
	KrbConf *string `field:"optional" json:"krbConf" yaml:"krbConf"`
	// Path to a keytab file to be used instead of a password.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/ad/0.4.4/docs#krb_keytab AdProvider#krb_keytab}
	KrbKeytab *string `field:"optional" json:"krbKeytab" yaml:"krbKeytab"`
	// The name of the kerberos realm (domain) we will use for authentication. (default: "", environment variable: AD_KRB_REALM).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/ad/0.4.4/docs#krb_realm AdProvider#krb_realm}
	KrbRealm *string `field:"optional" json:"krbRealm" yaml:"krbRealm"`
	// Alternative Service Principal Name. (default: none, environment variable: AD_KRB_SPN).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/ad/0.4.4/docs#krb_spn AdProvider#krb_spn}
	KrbSpn *string `field:"optional" json:"krbSpn" yaml:"krbSpn"`
	// Trust unknown certificates. (default: false, environment variable: AD_WINRM_INSECURE).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/ad/0.4.4/docs#winrm_insecure AdProvider#winrm_insecure}
	WinrmInsecure interface{} `field:"optional" json:"winrmInsecure" yaml:"winrmInsecure"`
	// Pass credentials in WinRM session to create a System.Management.Automation.PSCredential. (default: false, environment variable: AD_WINRM_PASS_CREDENTIALS).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/ad/0.4.4/docs#winrm_pass_credentials AdProvider#winrm_pass_credentials}
	WinrmPassCredentials interface{} `field:"optional" json:"winrmPassCredentials" yaml:"winrmPassCredentials"`
	// The port WinRM is listening for connections. (default: 5985, environment variable: AD_PORT).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/ad/0.4.4/docs#winrm_port AdProvider#winrm_port}
	WinrmPort *float64 `field:"optional" json:"winrmPort" yaml:"winrmPort"`
	// The WinRM protocol we will use. (default: http, environment variable: AD_PROTO).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/ad/0.4.4/docs#winrm_proto AdProvider#winrm_proto}
	WinrmProto *string `field:"optional" json:"winrmProto" yaml:"winrmProto"`
	// Use NTLM authentication. (default: false, environment variable: AD_WINRM_USE_NTLM).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/ad/0.4.4/docs#winrm_use_ntlm AdProvider#winrm_use_ntlm}
	WinrmUseNtlm interface{} `field:"optional" json:"winrmUseNtlm" yaml:"winrmUseNtlm"`
}

