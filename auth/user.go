package auth

import (
	"encoding/json"
	"fmt"

	"menteslibres.net/gosexy/redis"
)

const (
	RoleOrga  = "orga"
	RoleAdmin = "admin"
)

type User struct {
	ID       int64    `json:"id"`
	Login    string   `json:"login"`
	Password string   `json:"password"`
	Email    string   `json:"email"`
	Roles    []string `json:"roles"`
	Secret   string   `json:"secret"`
}

func (u User) HasRole(role string) bool {
	for _, r := range u.Roles {
		if r == role {
			return true
		}
	}
	return false
}

func (u User) IsAuthorized(role string) (bool, bool) {
	ok, admin := false, false
	for _, role := range u.Roles {
		if role == RoleAdmin {
			return true, true
		}
		if role == role {
			ok = true
		}
	}
	return ok, admin
}

func NewUserFromDatabase(db *redis.Client, id int64) (*User, error) {
	raw, err := db.Get(fmt.Sprintf("user:%d", id))
	if err != nil {
		return nil, err
	}

	var user User
	err = json.Unmarshal([]byte(raw), &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func NewUserFromClaims(raw interface{}) (*User, error) {
	encoded, err := json.Marshal(raw)
	if err != nil {
		return nil, err
	}

	var user User
	err = json.Unmarshal(encoded, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
