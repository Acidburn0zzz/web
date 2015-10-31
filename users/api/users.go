package api

import (
	"github.com/asaskevich/govalidator"
	"strings"
)

// User is a foo.
type User struct {
	ID        string `json:"id,omitempty"`
	FirstName string `json:"first_name" valid:"required,alphanum"`
	LastName  string `json:"last_name" valid:"required,alphanum"`
	Address   string `json:"address"`
	Phone     string `json:"phone"`
	Email     string `json:"email" valid:"required,email"`
	APIKey    string `json:"api_key"`

	errors []string
}

// Valid return true or false depending on whether or not the User is valid. It
// additionally sets the errors field on the User to provide information about
// why the user is not valid
func (u *User) Valid() bool {
	result, err := govalidator.ValidateStruct(u)
	if err != nil {
		u.errors = strings.Split(strings.TrimRight(err.Error(), ";"), ";")
	}
	return result
}

func validationRequiredFieldBlank(field string) string {
	str := []string{
		"Required field",
		field,
		"is blank.",
	}
	return strings.Join(str, " ")
}
