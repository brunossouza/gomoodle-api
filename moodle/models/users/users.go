package users

import (
	"fmt"
	"net/url"

	"gopkg.in/go-playground/validator.v9"
)

// User model
type User struct {
	ID                int           //ID of the user. Used only on update request.
	CreatePassword    bool          //Optional - True(createpassword=1) if password should be created and mailed to user. Used only on create request.
	Username          string        `validate:"required"` //Username policy is defined in Moodle security config.
	Auth              string        //Default to "manual" //Auth plugins include manual, ldap, etc.
	Suspended         int           //Optional - Suspend user account, either false to enable user login or true to disable it. Used only on update request.
	Password          string        //Optional - Plain text password consisting of any characters.
	Firstname         string        `validate:"required"` //The first name(s) of the user.
	Lastname          string        `validate:"required"` //The family name of the user.
	Email             string        `validate:"required"` //A valid and unique email address.
	MailDisplay       int           //Optional - Email display.
	City              string        //Optional - Home city of the user.
	Country           string        //Optional - Home country code of the user, such as AU or CZ.
	Timezone          string        //Optional - Timezone code such as Australia/Perth, or 99 for default.
	Description       string        //Optional - User profile description, no HTML.
	UserPicture       int           //Optional - The itemid where the new user picture has been uploaded to, 0 to delete. Used only on update request.
	FirstNamePhonetic string        //Optional - The first name(s) phonetically of the user.
	LastNamePhonetic  string        //Optional - The family name phonetically of the user.
	MiddleName        string        //Optional - The middle name of the user.
	AlternateName     string        //Optional - The alternate name of the user.
	Interests         string        //Optional - User interests (separated by commas).
	IDNumber          string        //Default  - "" //An arbitrary ID code number perhaps from the institution.
	Institution       string        //Optional - institution.
	Department        string        //Optional - department.
	Phone1            string        //Optional - Phone 1.
	Phone2            string        //Optional - Phone 2.
	Address           string        //Optional - Postal address.
	Lang              string        //Default  - "en" //Language code such as "en", must exist on server.
	CalendarType      string        //Default  - "gregorian" //Calendar type such as "gregorian", must exist on server.
	Theme             string        //Optional - Theme name such as "standard", must exist on server.
	MailFormat        int           //Optional - Mail format code is 0 for plain text, 1 for HTML etc.
	CustomFields      []CustomField //Optional - User custom fields (also known as user profil fields).
	Preferences       []Preference  //Optional - User preferences.
}

//CustomField user model
type CustomField struct {
	Type  string `validate:"required"` //The name of the custom field
	Value string `validate:"required"` //The value of the custom field
}

//Prefernce model
type Preference struct {
	Type  string `validate:"required"` //The name of the preference
	Value string `validate:"required"` //The value of the preference
}

//verifyUsersDataRequired Verify if user has all required fields.
func verifyUsersDataRequired(user User) error {
	return validator.New().Struct(user)
}

//verifyCustomFieldDataRequired Verify if CustomField has all fields with value.
func verifyCustomFieldDataRequired(field CustomField) error {
	return validator.New().Struct(field)
}

//verifyPreferenceDataRequired Verify if Preference has all fields with value.
func verifyPreferenceDataRequired(preference Preference) error {
	return validator.New().Struct(preference)
}

//parseUserToFormData Parses all optional data entered to the expected pattern of moodle rest API.
func parseUserToFormData(indice int, user User, data string) string {

	if user.Auth != "" {
		data += fmt.Sprintf("&users[%d][auth]=%s", indice, url.QueryEscape(user.Auth))
	}

	if user.MailDisplay != 0 {
		data += fmt.Sprintf("&users[%d][maildisplay]=%d", indice, user.MailDisplay)
	}

	if user.City != "" {
		data += fmt.Sprintf("&users[%d][city]=%s", indice, url.QueryEscape(user.City))
	}

	if user.Country != "" {
		data += fmt.Sprintf("&users[%d][country]=%s", indice, url.QueryEscape(user.Country))
	}

	if user.Timezone != "" {
		data += fmt.Sprintf("&users[%d][timezone]=%s", indice, url.QueryEscape(user.Timezone))
	}

	if user.Description != "" {
		data += fmt.Sprintf("&users[%d][description]=%s", indice, url.QueryEscape(user.Description))
	}

	if user.FirstNamePhonetic != "" {
		data += fmt.Sprintf("&users[%d][firstnamephonetic]=%s", indice, url.QueryEscape(user.FirstNamePhonetic))
	}

	if user.LastNamePhonetic != "" {
		data += fmt.Sprintf("&users[%d][lastnamephonetic]=%s", indice, url.QueryEscape(user.LastNamePhonetic))
	}

	if user.MiddleName != "" {
		data += fmt.Sprintf("&users[%d][middlename]=%s", indice, url.QueryEscape(user.MiddleName))
	}

	if user.AlternateName != "" {
		data += fmt.Sprintf("&users[%d][alternatename]=%s", indice, url.QueryEscape(user.AlternateName))
	}

	if user.Interests != "" {
		data += fmt.Sprintf("&users[%d][interests]=%s", indice, url.QueryEscape(user.Interests))
	}

	if user.IDNumber != "" {
		data += fmt.Sprintf("&users[%d][idnumber]=%s", indice, url.QueryEscape(user.IDNumber))
	}

	if user.Institution != "" {
		data += fmt.Sprintf("&users[%d][institution]=%s", indice, url.QueryEscape(user.Institution))
	}

	if user.Department != "" {
		data += fmt.Sprintf("&users[%d][department]=%s", indice, url.QueryEscape(user.Department))
	}

	if user.Phone1 != "" {
		data += fmt.Sprintf("&users[%d][phone1]=%s", indice, url.QueryEscape(user.Phone1))
	}

	if user.Phone2 != "" {
		data += fmt.Sprintf("&users[%d][phone2]=%s", indice, url.QueryEscape(user.Phone2))
	}

	if user.Address != "" {
		data += fmt.Sprintf("&users[%d][address]=%s", indice, url.QueryEscape(user.Address))
	}

	if user.Lang != "" {
		data += fmt.Sprintf("&users[%d][lang]=%s", indice, url.QueryEscape(user.Lang))
	}

	if user.CalendarType != "" {
		data += fmt.Sprintf("&users[%d][calendartype]=%s", indice, url.QueryEscape(user.CalendarType))
	}

	if user.Theme != "" {
		data += fmt.Sprintf("&users[%d][theme]=%s", indice, url.QueryEscape(user.Theme))
	}

	if user.MailFormat != 0 {
		data += fmt.Sprintf("&users[%d][mailformat]=%d", indice, user.MailFormat)
	}

	for idx, field := range user.CustomFields {
		if err := verifyCustomFieldDataRequired(field); err == nil {
			data += fmt.Sprintf("&users[%d][customfields][%d][type]=%s&users[%d][customfields][%d][value]=%s", indice, idx, url.QueryEscape(field.Type), indice, idx, url.QueryEscape(field.Value))
		}
	}

	for idx, preference := range user.Preferences {
		if err := verifyPreferenceDataRequired(preference); err == nil {
			data += fmt.Sprintf("&users[%d][preferences][%d][type]=%s&users[%d][preferences][%d][value]=%s", indice, idx, url.QueryEscape(preference.Type), indice, idx, url.QueryEscape(preference.Value))
		}
	}
	return data
}

//ParseUserToCreateFormData Parse all mandatory and optional data for creating a new user to the expected default of moodle rest API.
func ParseUserToCreateFormData(indice int, user User) (string, error) {
	//Parse commom data
	err := verifyUsersDataRequired(user)
	if err != nil {
		return "", err
	}

	//Required data
	data := fmt.Sprintf("&users[%d][username]=%s&users[%d][firstname]=%s&users[%d][lastname]=%s&users[%d][email]=%s", indice, url.QueryEscape(user.Username), indice, url.QueryEscape(user.Firstname), indice, url.QueryEscape(user.Lastname), indice, url.QueryEscape(user.Email))

	//Optional data create form
	if user.CreatePassword {
		data += fmt.Sprintf("&users[%d][createpassword]=1", indice)
	} else {
		data += fmt.Sprintf("&users[%d][password]=%s", indice, url.QueryEscape(user.Password))
	}

	data = parseUserToFormData(indice, user, data)

	return data, nil
}

//ParseUserToUpdateFormData Parse all mandatory and optional data to update the informed user data to the expected default of moodle rest API.
func ParseUserToUpdateFormData(indice int, user User) (string, error) {
	if user.ID == 0 {
		return "", fmt.Errorf("the id field must be informed to perform the update")
	}

	//Parse commom data
	err := verifyUsersDataRequired(user)
	if err != nil {
		return "", err
	}

	//Requered data to update
	data := fmt.Sprintf("users[%d][id]=%d", indice, user.ID)

	//Optional data update form
	if user.Username == "" {
		data += fmt.Sprintf("&users[%d][username]=%s", indice, url.QueryEscape(user.Username))
	}

	if user.Firstname == "" {
		data += fmt.Sprintf("&users[%d][firstname]=%s", indice, url.QueryEscape(user.Firstname))
	}

	if user.Lastname == "" {
		data += fmt.Sprintf("&users[%d][lastname]=%s", indice, url.QueryEscape(user.Lastname))
	}

	if user.Email == "" {
		data += fmt.Sprintf("&users[%d][email]=%s", indice, url.QueryEscape(user.Email))
	}

	if user.Suspended == 0 {
		data += fmt.Sprintf("&users[%d][suspended]=1", indice)
	}

	if user.UserPicture == 0 {
		data += fmt.Sprintf("&users[%d][userpicture]=%d", indice, user.UserPicture)
	}

	data = parseUserToFormData(indice, user, data)

	return data, nil
}
