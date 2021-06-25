package users

import (
	"fmt"
	"net/url"

	"gopkg.in/go-playground/validator.v9"
)

type User struct {
	CreatePassword    bool          //Optional - True(createpassword=1) if password should be created and mailed to user.
	Username          string        `validate:"required"` //Username policy is defined in Moodle security config.
	Auth              string        //Default to "manual" //Auth plugins include manual, ldap, etc
	Password          string        //Optional - Plain text password consisting of any characters
	Firstname         string        `validate:"required"` //The first name(s) of the user
	Lastname          string        `validate:"required"` //The family name of the user
	Email             string        `validate:"required"` //A valid and unique email address
	MailDisplay       int           //Optional - Email display
	City              string        //Optional - Home city of the user
	Country           string        //Optional - Home country code of the user, such as AU or CZ
	Timezone          string        //Optional - Timezone code such as Australia/Perth, or 99 for default
	Description       string        //Optional - User profile description, no HTML
	FirstnamePhonetic string        //Optional - The first name(s) phonetically of the user
	LastnamePhonetic  string        //Optional - The family name phonetically of the user
	Middlename        string        //Optional - The middle name of the user
	AlternateName     string        //Optional - The alternate name of the user
	Interests         string        //Optional - User interests (separated by commas)
	Idnumber          string        //Default  - "" //An arbitrary ID code number perhaps from the institution
	Institution       string        //Optional - institution
	Department        string        //Optional - department
	Phone1            string        //Optional - Phone 1
	Phone2            string        //Optional - Phone 2
	Address           string        //Optional - Postal address
	Lang              string        //Default  - "en" //Language code such as "en", must exist on server
	CalendarType      string        //Default  - "gregorian" //Calendar type such as "gregorian", must exist on server
	Theme             string        //Optional - Theme name such as "standard", must exist on server
	MailFormat        int           //Optional - Mail format code is 0 for plain text, 1 for HTML etc
	CustomFields      []CustomField //Optional - User custom fields (also known as user profil fields)
	Preferences       []Preference  //Optional - User preferences
}

type CustomField struct {
	Type  string `validate:"required"` //The name of the custom field
	Value string `validate:"required"` //The value of the custom field
}

type Preference struct {
	Type  string `validate:"required"` //The name of the preference
	Value string `validate:"required"` //The value of the preference
}

func verifyUsersDataRequired(user User) error {
	return validator.New().Struct(user)
}

func verifyCustomFieldDataRequired(field CustomField) error {
	return validator.New().Struct(field)
}

func verifyPreferenceDataRequired(preference Preference) error {
	return validator.New().Struct(preference)
}

func ParseUserToFormData(indice int, user User) (string, error) {
	err := verifyUsersDataRequired(user)
	if err != nil {
		return "", err
	}
	//Required data
	data := fmt.Sprintf("&users[%d][username]=%s&users[%d][firstname]=%s&users[%d][lastname]=%s&users[%d][email]=%s", indice, url.QueryEscape(user.Username), indice, url.QueryEscape(user.Firstname), indice, url.QueryEscape(user.Lastname), indice, url.QueryEscape(user.Email))

	//Required data - Password
	if user.CreatePassword {
		data += fmt.Sprintf("&users[%d][createpassword]=1", indice)
	} else {
		data += fmt.Sprintf("&users[%d][password]=%s", indice, url.QueryEscape(user.Password))
	}

	//Optional data
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

	if user.FirstnamePhonetic != "" {
		data += fmt.Sprintf("&users[%d][firstnamephonetic]=%s", indice, url.QueryEscape(user.FirstnamePhonetic))
	}

	if user.LastnamePhonetic != "" {
		data += fmt.Sprintf("&users[%d][lastnamephonetic]=%s", indice, url.QueryEscape(user.LastnamePhonetic))
	}

	if user.Middlename != "" {
		data += fmt.Sprintf("&users[%d][middlename]=%s", indice, url.QueryEscape(user.Middlename))
	}

	if user.AlternateName != "" {
		data += fmt.Sprintf("&users[%d][alternatename]=%s", indice, url.QueryEscape(user.AlternateName))
	}

	if user.Interests != "" {
		data += fmt.Sprintf("&users[%d][interests]=%s", indice, url.QueryEscape(user.Interests))
	}

	if user.Idnumber != "" {
		data += fmt.Sprintf("&users[%d][idnumber]=%s", indice, url.QueryEscape(user.Idnumber))
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
	return data, nil
}
