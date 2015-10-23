package api

// User is a foo.
type User struct {
	ID        string `json:"id,omitempty"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Address   string `json:"address"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	ApiKey    string `json:"api_key"`

	errors string
}

func (u *User) Valid() bool {
	if len(u.FirstName) == 0 {
		return false
	}

	if len(u.LastName) == 0 {
		return false
	}

	if len(u.Email) == 0 {
		return false
	}

	return true
}
