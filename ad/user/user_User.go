package user

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-ad-go/ad/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-ad-go/ad/user/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/ad/r/user ad_user}.
type User interface {
	cdktf.TerraformResource
	CannotChangePassword() interface{}
	SetCannotChangePassword(val interface{})
	CannotChangePasswordInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	City() *string
	SetCity(val *string)
	CityInput() *string
	Company() *string
	SetCompany(val *string)
	CompanyInput() *string
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	Container() *string
	SetContainer(val *string)
	ContainerInput() *string
	// Experimental.
	Count() *float64
	// Experimental.
	SetCount(val *float64)
	Country() *string
	SetCountry(val *string)
	CountryInput() *string
	CustomAttributes() *string
	SetCustomAttributes(val *string)
	CustomAttributesInput() *string
	Department() *string
	SetDepartment(val *string)
	DepartmentInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
	Division() *string
	SetDivision(val *string)
	DivisionInput() *string
	Dn() *string
	EmailAddress() *string
	SetEmailAddress(val *string)
	EmailAddressInput() *string
	EmployeeId() *string
	SetEmployeeId(val *string)
	EmployeeIdInput() *string
	EmployeeNumber() *string
	SetEmployeeNumber(val *string)
	EmployeeNumberInput() *string
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	Fax() *string
	SetFax(val *string)
	FaxInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GivenName() *string
	SetGivenName(val *string)
	GivenNameInput() *string
	HomeDirectory() *string
	SetHomeDirectory(val *string)
	HomeDirectoryInput() *string
	HomeDrive() *string
	SetHomeDrive(val *string)
	HomeDriveInput() *string
	HomePage() *string
	SetHomePage(val *string)
	HomePageInput() *string
	HomePhone() *string
	SetHomePhone(val *string)
	HomePhoneInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	InitialPassword() *string
	SetInitialPassword(val *string)
	InitialPasswordInput() *string
	Initials() *string
	SetInitials(val *string)
	InitialsInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MobilePhone() *string
	SetMobilePhone(val *string)
	MobilePhoneInput() *string
	// The tree node.
	Node() constructs.Node
	Office() *string
	SetOffice(val *string)
	OfficeInput() *string
	OfficePhone() *string
	SetOfficePhone(val *string)
	OfficePhoneInput() *string
	Organization() *string
	SetOrganization(val *string)
	OrganizationInput() *string
	OtherName() *string
	SetOtherName(val *string)
	OtherNameInput() *string
	PasswordNeverExpires() interface{}
	SetPasswordNeverExpires(val interface{})
	PasswordNeverExpiresInput() interface{}
	PoBox() *string
	SetPoBox(val *string)
	PoBoxInput() *string
	PostalCode() *string
	SetPostalCode(val *string)
	PostalCodeInput() *string
	PrincipalName() *string
	SetPrincipalName(val *string)
	PrincipalNameInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	SamAccountName() *string
	SetSamAccountName(val *string)
	SamAccountNameInput() *string
	Sid() *string
	SmartCardLogonRequired() interface{}
	SetSmartCardLogonRequired(val interface{})
	SmartCardLogonRequiredInput() interface{}
	State() *string
	SetState(val *string)
	StateInput() *string
	StreetAddress() *string
	SetStreetAddress(val *string)
	StreetAddressInput() *string
	Surname() *string
	SetSurname(val *string)
	SurnameInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Title() *string
	SetTitle(val *string)
	TitleInput() *string
	TrustedForDelegation() interface{}
	SetTrustedForDelegation(val interface{})
	TrustedForDelegationInput() interface{}
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	ResetCannotChangePassword()
	ResetCity()
	ResetCompany()
	ResetContainer()
	ResetCountry()
	ResetCustomAttributes()
	ResetDepartment()
	ResetDescription()
	ResetDivision()
	ResetEmailAddress()
	ResetEmployeeId()
	ResetEmployeeNumber()
	ResetEnabled()
	ResetFax()
	ResetGivenName()
	ResetHomeDirectory()
	ResetHomeDrive()
	ResetHomePage()
	ResetHomePhone()
	ResetId()
	ResetInitialPassword()
	ResetInitials()
	ResetMobilePhone()
	ResetOffice()
	ResetOfficePhone()
	ResetOrganization()
	ResetOtherName()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPasswordNeverExpires()
	ResetPoBox()
	ResetPostalCode()
	ResetSmartCardLogonRequired()
	ResetState()
	ResetStreetAddress()
	ResetSurname()
	ResetTitle()
	ResetTrustedForDelegation()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for User
type jsiiProxy_User struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_User) CannotChangePassword() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cannotChangePassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) CannotChangePasswordInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cannotChangePasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) City() *string {
	var returns *string
	_jsii_.Get(
		j,
		"city",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) CityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Company() *string {
	var returns *string
	_jsii_.Get(
		j,
		"company",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) CompanyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"companyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Container() *string {
	var returns *string
	_jsii_.Get(
		j,
		"container",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) ContainerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Country() *string {
	var returns *string
	_jsii_.Get(
		j,
		"country",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) CountryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"countryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) CustomAttributes() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) CustomAttributesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customAttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Department() *string {
	var returns *string
	_jsii_.Get(
		j,
		"department",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) DepartmentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"departmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Division() *string {
	var returns *string
	_jsii_.Get(
		j,
		"division",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) DivisionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"divisionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Dn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) EmailAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) EmailAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) EmployeeId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"employeeId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) EmployeeIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"employeeIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) EmployeeNumber() *string {
	var returns *string
	_jsii_.Get(
		j,
		"employeeNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) EmployeeNumberInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"employeeNumberInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Fax() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fax",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) FaxInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"faxInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) GivenName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"givenName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) GivenNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"givenNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) HomeDirectory() *string {
	var returns *string
	_jsii_.Get(
		j,
		"homeDirectory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) HomeDirectoryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"homeDirectoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) HomeDrive() *string {
	var returns *string
	_jsii_.Get(
		j,
		"homeDrive",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) HomeDriveInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"homeDriveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) HomePage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"homePage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) HomePageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"homePageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) HomePhone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"homePhone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) HomePhoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"homePhoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) InitialPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"initialPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) InitialPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"initialPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Initials() *string {
	var returns *string
	_jsii_.Get(
		j,
		"initials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) InitialsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"initialsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) MobilePhone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mobilePhone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) MobilePhoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mobilePhoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Office() *string {
	var returns *string
	_jsii_.Get(
		j,
		"office",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) OfficeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"officeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) OfficePhone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"officePhone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) OfficePhoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"officePhoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Organization() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) OrganizationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organizationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) OtherName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"otherName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) OtherNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"otherNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) PasswordNeverExpires() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"passwordNeverExpires",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) PasswordNeverExpiresInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"passwordNeverExpiresInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) PoBox() *string {
	var returns *string
	_jsii_.Get(
		j,
		"poBox",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) PoBoxInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"poBoxInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) PostalCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"postalCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) PostalCodeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"postalCodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) PrincipalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"principalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) PrincipalNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"principalNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) SamAccountName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"samAccountName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) SamAccountNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"samAccountNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Sid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) SmartCardLogonRequired() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"smartCardLogonRequired",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) SmartCardLogonRequiredInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"smartCardLogonRequiredInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) StateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) StreetAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streetAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) StreetAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streetAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Surname() *string {
	var returns *string
	_jsii_.Get(
		j,
		"surname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) SurnameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"surnameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Title() *string {
	var returns *string
	_jsii_.Get(
		j,
		"title",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) TitleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"titleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) TrustedForDelegation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"trustedForDelegation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) TrustedForDelegationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"trustedForDelegationInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/ad/r/user ad_user} Resource.
func NewUser(scope constructs.Construct, id *string, config *UserConfig) User {
	_init_.Initialize()

	if err := validateNewUserParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_User{}

	_jsii_.Create(
		"@cdktf/provider-ad.user.User",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/ad/r/user ad_user} Resource.
func NewUser_Override(u User, scope constructs.Construct, id *string, config *UserConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-ad.user.User",
		[]interface{}{scope, id, config},
		u,
	)
}

func (j *jsiiProxy_User)SetCannotChangePassword(val interface{}) {
	if err := j.validateSetCannotChangePasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cannotChangePassword",
		val,
	)
}

func (j *jsiiProxy_User)SetCity(val *string) {
	if err := j.validateSetCityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"city",
		val,
	)
}

func (j *jsiiProxy_User)SetCompany(val *string) {
	if err := j.validateSetCompanyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"company",
		val,
	)
}

func (j *jsiiProxy_User)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_User)SetContainer(val *string) {
	if err := j.validateSetContainerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"container",
		val,
	)
}

func (j *jsiiProxy_User)SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_User)SetCountry(val *string) {
	if err := j.validateSetCountryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"country",
		val,
	)
}

func (j *jsiiProxy_User)SetCustomAttributes(val *string) {
	if err := j.validateSetCustomAttributesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customAttributes",
		val,
	)
}

func (j *jsiiProxy_User)SetDepartment(val *string) {
	if err := j.validateSetDepartmentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"department",
		val,
	)
}

func (j *jsiiProxy_User)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_User)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_User)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_User)SetDivision(val *string) {
	if err := j.validateSetDivisionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"division",
		val,
	)
}

func (j *jsiiProxy_User)SetEmailAddress(val *string) {
	if err := j.validateSetEmailAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"emailAddress",
		val,
	)
}

func (j *jsiiProxy_User)SetEmployeeId(val *string) {
	if err := j.validateSetEmployeeIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"employeeId",
		val,
	)
}

func (j *jsiiProxy_User)SetEmployeeNumber(val *string) {
	if err := j.validateSetEmployeeNumberParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"employeeNumber",
		val,
	)
}

func (j *jsiiProxy_User)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_User)SetFax(val *string) {
	if err := j.validateSetFaxParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fax",
		val,
	)
}

func (j *jsiiProxy_User)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_User)SetGivenName(val *string) {
	if err := j.validateSetGivenNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"givenName",
		val,
	)
}

func (j *jsiiProxy_User)SetHomeDirectory(val *string) {
	if err := j.validateSetHomeDirectoryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"homeDirectory",
		val,
	)
}

func (j *jsiiProxy_User)SetHomeDrive(val *string) {
	if err := j.validateSetHomeDriveParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"homeDrive",
		val,
	)
}

func (j *jsiiProxy_User)SetHomePage(val *string) {
	if err := j.validateSetHomePageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"homePage",
		val,
	)
}

func (j *jsiiProxy_User)SetHomePhone(val *string) {
	if err := j.validateSetHomePhoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"homePhone",
		val,
	)
}

func (j *jsiiProxy_User)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_User)SetInitialPassword(val *string) {
	if err := j.validateSetInitialPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"initialPassword",
		val,
	)
}

func (j *jsiiProxy_User)SetInitials(val *string) {
	if err := j.validateSetInitialsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"initials",
		val,
	)
}

func (j *jsiiProxy_User)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_User)SetMobilePhone(val *string) {
	if err := j.validateSetMobilePhoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mobilePhone",
		val,
	)
}

func (j *jsiiProxy_User)SetOffice(val *string) {
	if err := j.validateSetOfficeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"office",
		val,
	)
}

func (j *jsiiProxy_User)SetOfficePhone(val *string) {
	if err := j.validateSetOfficePhoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"officePhone",
		val,
	)
}

func (j *jsiiProxy_User)SetOrganization(val *string) {
	if err := j.validateSetOrganizationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"organization",
		val,
	)
}

func (j *jsiiProxy_User)SetOtherName(val *string) {
	if err := j.validateSetOtherNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"otherName",
		val,
	)
}

func (j *jsiiProxy_User)SetPasswordNeverExpires(val interface{}) {
	if err := j.validateSetPasswordNeverExpiresParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"passwordNeverExpires",
		val,
	)
}

func (j *jsiiProxy_User)SetPoBox(val *string) {
	if err := j.validateSetPoBoxParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"poBox",
		val,
	)
}

func (j *jsiiProxy_User)SetPostalCode(val *string) {
	if err := j.validateSetPostalCodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"postalCode",
		val,
	)
}

func (j *jsiiProxy_User)SetPrincipalName(val *string) {
	if err := j.validateSetPrincipalNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"principalName",
		val,
	)
}

func (j *jsiiProxy_User)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_User)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_User)SetSamAccountName(val *string) {
	if err := j.validateSetSamAccountNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"samAccountName",
		val,
	)
}

func (j *jsiiProxy_User)SetSmartCardLogonRequired(val interface{}) {
	if err := j.validateSetSmartCardLogonRequiredParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"smartCardLogonRequired",
		val,
	)
}

func (j *jsiiProxy_User)SetState(val *string) {
	if err := j.validateSetStateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"state",
		val,
	)
}

func (j *jsiiProxy_User)SetStreetAddress(val *string) {
	if err := j.validateSetStreetAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"streetAddress",
		val,
	)
}

func (j *jsiiProxy_User)SetSurname(val *string) {
	if err := j.validateSetSurnameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"surname",
		val,
	)
}

func (j *jsiiProxy_User)SetTitle(val *string) {
	if err := j.validateSetTitleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"title",
		val,
	)
}

func (j *jsiiProxy_User)SetTrustedForDelegation(val interface{}) {
	if err := j.validateSetTrustedForDelegationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trustedForDelegation",
		val,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func User_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateUser_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-ad.user.User",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func User_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-ad.user.User",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (u *jsiiProxy_User) AddOverride(path *string, value interface{}) {
	if err := u.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		u,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (u *jsiiProxy_User) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := u.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		u,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_User) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := u.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		u,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_User) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := u.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		u,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_User) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := u.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		u,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_User) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := u.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		u,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_User) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := u.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		u,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_User) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := u.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		u,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_User) GetStringAttribute(terraformAttribute *string) *string {
	if err := u.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		u,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_User) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := u.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		u,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_User) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := u.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		u,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_User) OverrideLogicalId(newLogicalId *string) {
	if err := u.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		u,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (u *jsiiProxy_User) ResetCannotChangePassword() {
	_jsii_.InvokeVoid(
		u,
		"resetCannotChangePassword",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetCity() {
	_jsii_.InvokeVoid(
		u,
		"resetCity",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetCompany() {
	_jsii_.InvokeVoid(
		u,
		"resetCompany",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetContainer() {
	_jsii_.InvokeVoid(
		u,
		"resetContainer",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetCountry() {
	_jsii_.InvokeVoid(
		u,
		"resetCountry",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetCustomAttributes() {
	_jsii_.InvokeVoid(
		u,
		"resetCustomAttributes",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetDepartment() {
	_jsii_.InvokeVoid(
		u,
		"resetDepartment",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetDescription() {
	_jsii_.InvokeVoid(
		u,
		"resetDescription",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetDivision() {
	_jsii_.InvokeVoid(
		u,
		"resetDivision",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetEmailAddress() {
	_jsii_.InvokeVoid(
		u,
		"resetEmailAddress",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetEmployeeId() {
	_jsii_.InvokeVoid(
		u,
		"resetEmployeeId",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetEmployeeNumber() {
	_jsii_.InvokeVoid(
		u,
		"resetEmployeeNumber",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetEnabled() {
	_jsii_.InvokeVoid(
		u,
		"resetEnabled",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetFax() {
	_jsii_.InvokeVoid(
		u,
		"resetFax",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetGivenName() {
	_jsii_.InvokeVoid(
		u,
		"resetGivenName",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetHomeDirectory() {
	_jsii_.InvokeVoid(
		u,
		"resetHomeDirectory",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetHomeDrive() {
	_jsii_.InvokeVoid(
		u,
		"resetHomeDrive",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetHomePage() {
	_jsii_.InvokeVoid(
		u,
		"resetHomePage",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetHomePhone() {
	_jsii_.InvokeVoid(
		u,
		"resetHomePhone",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetId() {
	_jsii_.InvokeVoid(
		u,
		"resetId",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetInitialPassword() {
	_jsii_.InvokeVoid(
		u,
		"resetInitialPassword",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetInitials() {
	_jsii_.InvokeVoid(
		u,
		"resetInitials",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetMobilePhone() {
	_jsii_.InvokeVoid(
		u,
		"resetMobilePhone",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetOffice() {
	_jsii_.InvokeVoid(
		u,
		"resetOffice",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetOfficePhone() {
	_jsii_.InvokeVoid(
		u,
		"resetOfficePhone",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetOrganization() {
	_jsii_.InvokeVoid(
		u,
		"resetOrganization",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetOtherName() {
	_jsii_.InvokeVoid(
		u,
		"resetOtherName",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		u,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetPasswordNeverExpires() {
	_jsii_.InvokeVoid(
		u,
		"resetPasswordNeverExpires",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetPoBox() {
	_jsii_.InvokeVoid(
		u,
		"resetPoBox",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetPostalCode() {
	_jsii_.InvokeVoid(
		u,
		"resetPostalCode",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetSmartCardLogonRequired() {
	_jsii_.InvokeVoid(
		u,
		"resetSmartCardLogonRequired",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetState() {
	_jsii_.InvokeVoid(
		u,
		"resetState",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetStreetAddress() {
	_jsii_.InvokeVoid(
		u,
		"resetStreetAddress",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetSurname() {
	_jsii_.InvokeVoid(
		u,
		"resetSurname",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetTitle() {
	_jsii_.InvokeVoid(
		u,
		"resetTitle",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetTrustedForDelegation() {
	_jsii_.InvokeVoid(
		u,
		"resetTrustedForDelegation",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		u,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (u *jsiiProxy_User) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		u,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (u *jsiiProxy_User) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		u,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (u *jsiiProxy_User) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		u,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

