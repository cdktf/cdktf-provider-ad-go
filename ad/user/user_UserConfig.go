package user

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type UserConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The Display Name of an Active Directory user.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ad/r/user#display_name User#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// The Principal Name of an Active Directory user.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ad/r/user#principal_name User#principal_name}
	PrincipalName *string `field:"required" json:"principalName" yaml:"principalName"`
	// The pre-win2k user logon name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ad/r/user#sam_account_name User#sam_account_name}
	SamAccountName *string `field:"required" json:"samAccountName" yaml:"samAccountName"`
	// If set to true, the user will not be allowed to change their password.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ad/r/user#cannot_change_password User#cannot_change_password}
	CannotChangePassword interface{} `field:"optional" json:"cannotChangePassword" yaml:"cannotChangePassword"`
	// Specifies the user's town or city. This parameter sets the City property of a user object.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ad/r/user#city User#city}
	City *string `field:"optional" json:"city" yaml:"city"`
	// Specifies the user's company. This parameter sets the Company property of a user object.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ad/r/user#company User#company}
	Company *string `field:"optional" json:"company" yaml:"company"`
	// A DN of the container object that will be holding the user.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ad/r/user#container User#container}
	Container *string `field:"optional" json:"container" yaml:"container"`
	// Specifies the country by setting the country code (refer to ISO 3166).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ad/r/user#country User#country}
	Country *string `field:"optional" json:"country" yaml:"country"`
	// JSON encoded map that represents key/value pairs for custom attributes.
	//
	// Please note that `terraform import` will not import these attributes.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ad/r/user#custom_attributes User#custom_attributes}
	CustomAttributes *string `field:"optional" json:"customAttributes" yaml:"customAttributes"`
	// Specifies the user's department. This parameter sets the Department property of a user object.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ad/r/user#department User#department}
	Department *string `field:"optional" json:"department" yaml:"department"`
	// Specifies a description of the object. This parameter sets the value of the Description property for the user object.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ad/r/user#description User#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Specifies the user's division. This parameter sets the Division property of a user object.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ad/r/user#division User#division}
	Division *string `field:"optional" json:"division" yaml:"division"`
	// Specifies the user's e-mail address. This parameter sets the EmailAddress property of a user object.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ad/r/user#email_address User#email_address}
	EmailAddress *string `field:"optional" json:"emailAddress" yaml:"emailAddress"`
	// Specifies the user's employee ID. This parameter sets the EmployeeID property of a user object.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ad/r/user#employee_id User#employee_id}
	EmployeeId *string `field:"optional" json:"employeeId" yaml:"employeeId"`
	// Specifies the user's employee number. This parameter sets the EmployeeNumber property of a user object.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ad/r/user#employee_number User#employee_number}
	EmployeeNumber *string `field:"optional" json:"employeeNumber" yaml:"employeeNumber"`
	// If set to false, the user will be disabled.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ad/r/user#enabled User#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Specifies the user's fax phone number. This parameter sets the Fax property of a user object.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ad/r/user#fax User#fax}
	Fax *string `field:"optional" json:"fax" yaml:"fax"`
	// Specifies the user's given name. This parameter sets the GivenName property of a user object.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ad/r/user#given_name User#given_name}
	GivenName *string `field:"optional" json:"givenName" yaml:"givenName"`
	// Specifies a user's home directory. This parameter sets the HomeDirectory property of a user object.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ad/r/user#home_directory User#home_directory}
	HomeDirectory *string `field:"optional" json:"homeDirectory" yaml:"homeDirectory"`
	// Specifies a drive that is associated with the UNC path defined by the HomeDirectory property.
	//
	// The drive letter is specified as <DriveLetter>: where <DriveLetter> indicates the letter of the drive to associate. The <DriveLetter> must be a single, uppercase letter and the colon is required. This parameter sets the HomeDrive property of the user object.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ad/r/user#home_drive User#home_drive}
	HomeDrive *string `field:"optional" json:"homeDrive" yaml:"homeDrive"`
	// Specifies the URL of the home page of the object.
	//
	// This parameter sets the homePage property of a user object.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ad/r/user#home_page User#home_page}
	HomePage *string `field:"optional" json:"homePage" yaml:"homePage"`
	// Specifies the user's home telephone number. This parameter sets the HomePhone property of a user object.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ad/r/user#home_phone User#home_phone}
	HomePhone *string `field:"optional" json:"homePhone" yaml:"homePhone"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ad/r/user#id User#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The user's initial password. This will be set on creation but will *not* be enforced in subsequent plans.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ad/r/user#initial_password User#initial_password}
	InitialPassword *string `field:"optional" json:"initialPassword" yaml:"initialPassword"`
	// Specifies the initials that represent part of a user's name. Maximum 6 char.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ad/r/user#initials User#initials}
	Initials *string `field:"optional" json:"initials" yaml:"initials"`
	// Specifies the user's mobile phone number. This parameter sets the MobilePhone property of a user object.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ad/r/user#mobile_phone User#mobile_phone}
	MobilePhone *string `field:"optional" json:"mobilePhone" yaml:"mobilePhone"`
	// Specifies the location of the user's office or place of business.
	//
	// This parameter sets the Office property of a user object.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ad/r/user#office User#office}
	Office *string `field:"optional" json:"office" yaml:"office"`
	// Specifies the user's office telephone number. This parameter sets the OfficePhone property of a user object.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ad/r/user#office_phone User#office_phone}
	OfficePhone *string `field:"optional" json:"officePhone" yaml:"officePhone"`
	// Specifies the user's organization. This parameter sets the Organization property of a user object.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ad/r/user#organization User#organization}
	Organization *string `field:"optional" json:"organization" yaml:"organization"`
	// Specifies a name in addition to a user's given name and surname, such as the user's middle name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ad/r/user#other_name User#other_name}
	OtherName *string `field:"optional" json:"otherName" yaml:"otherName"`
	// If set to true, the password for this user will not expire.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ad/r/user#password_never_expires User#password_never_expires}
	PasswordNeverExpires interface{} `field:"optional" json:"passwordNeverExpires" yaml:"passwordNeverExpires"`
	// Specifies the user's post office box number. This parameter sets the POBox property of a user object.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ad/r/user#po_box User#po_box}
	PoBox *string `field:"optional" json:"poBox" yaml:"poBox"`
	// Specifies the user's postal code or zip code. This parameter sets the PostalCode property of a user object.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ad/r/user#postal_code User#postal_code}
	PostalCode *string `field:"optional" json:"postalCode" yaml:"postalCode"`
	// If set to true, a smart card is required to logon.
	//
	// This parameter sets the SmartCardLoginRequired property for a user object.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ad/r/user#smart_card_logon_required User#smart_card_logon_required}
	SmartCardLogonRequired interface{} `field:"optional" json:"smartCardLogonRequired" yaml:"smartCardLogonRequired"`
	// Specifies the user's or Organizational Unit's state or province. This parameter sets the State property of a user object.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ad/r/user#state User#state}
	State *string `field:"optional" json:"state" yaml:"state"`
	// Specifies the user's street address. This parameter sets the StreetAddress property of a user object.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ad/r/user#street_address User#street_address}
	StreetAddress *string `field:"optional" json:"streetAddress" yaml:"streetAddress"`
	// Specifies the user's last name or surname. This parameter sets the Surname property of a user object.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ad/r/user#surname User#surname}
	Surname *string `field:"optional" json:"surname" yaml:"surname"`
	// Specifies the user's title. This parameter sets the Title property of a user object.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ad/r/user#title User#title}
	Title *string `field:"optional" json:"title" yaml:"title"`
	// If set to true, the user account is trusted for Kerberos delegation.
	//
	// A service that runs under an account that is trusted for Kerberos delegation can assume the identity of a client requesting the service. This parameter sets the TrustedForDelegation property of an account object.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ad/r/user#trusted_for_delegation User#trusted_for_delegation}
	TrustedForDelegation interface{} `field:"optional" json:"trustedForDelegation" yaml:"trustedForDelegation"`
}

