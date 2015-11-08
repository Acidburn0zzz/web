package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strings"

	"upper.io/db"

	"github.com/asaskevich/govalidator"
)

// User is a foo.
type User struct {
	ID        string `json:"id,omitempty" db:"id,omitempty"`
	FirstName string `json:"first_name" valid:"required,alphanum" db:"first_name"`
	LastName  string `json:"last_name" valid:"required,alphanum" db:"last_name"`
	Username  string `json:"username" valid:"required" db:"username"`
	Address   string `json:"address" db:"address"`
	Phone     string `json:"phone" db:"phone"`
	Email     string `json:"email" valid:"required,email" db:"email"`
	APIKey    string `json:"api_key" db:"apikey"`
	Password  string `json:"-" db:"password"`
	Banned    bool   `db:"banned"`

	errors []string
	db     db.Database
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

func userCollection() db.Collection {
	col, err := database.Collection("users")

	if err != nil {
		log.Fatalln("Failed to find users collection", err.Error())
	}

	return col
}

func UserCount() {
	uc := userCollection()
	cnt, _ := uc.Find().Count()

	fmt.Println("User Count: ", cnt)
}

func UserFind(id string) (*User, error) {
	if len(id) == 0 {
		return nil, errors.New("ID unspecified")
	}

	var u User
	uc := userCollection()
	res := uc.Find(db.Cond{"id": id})
	res.One(&u)

	return &u, nil
}

// Save the user to the database
func (u *User) Save() bool {
	if len(u.ID) == 0 {
		return u.Create()
	}

	return u.Update()
}

// SetID allows us to update the struct after the DB sets the ID
func (u *User) SetID(values map[string]interface{}) error {
	if valueInterface, ok := values["id"]; ok {
		u.ID = valueInterface.(string)
	}
	return nil
}

// Create a user
func (u *User) Create() bool {
	if u.Valid() == false {
		return false
	}

	uc := userCollection()

	if uc == nil {
		return false
	}

	var err error
	if _, err = uc.Append(u); err != nil {
		log.Println("failed to create user", err.Error())
		return false
	}

	return true
}

// Update a user
func (u *User) Update() bool {
	uc := userCollection()

	if uc == nil {
		return false
	}

	res := uc.Find(db.Cond{"id": u.ID})

	if err := res.Update(u); err != nil {
		log.Println("failed to update user", err.Error())
		return false
	}

	return true
}

func (u *User) String() string {
	buf := new(bytes.Buffer)

	json.NewEncoder(buf).Encode(u)

	return buf.String()
}
