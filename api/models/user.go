package models

import "time"

type User struct {
	ID                string    `json:"id"`
	Username          string    `json:"username"`
	PasswordHash      string    `json:"-"`
	Name              string    `json:"name"`
	Bio               string    `json:"bio"`
	ProfilePicture    string    `json:"profile_picture"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

type CreateUser struct {
	Username          string `json:"username"`
	Password          string `json:"password"`
	Name              string `json:"name"`
	Bio               string `json:"bio,omitempty"`
	ProfilePicture    string `json:"profile_picture,omitempty"`
}

type UpdateUser struct {
	ID                string  `json:"-"`
	Name              *string `json:"name,omitempty"`
	Bio               *string `json:"bio,omitempty"`
	ProfilePicture    *string `json:"profile_picture,omitempty"`
}

type UsersResponse struct {
	Users []User `json:"users"`
	Count int    `json:"count"`
}

type UpdateUserPassword struct {
	ID          string `json:"-"`
	NewPassword string `json:"new_password"`
	OldPassword string `json:"old_password"`
}

