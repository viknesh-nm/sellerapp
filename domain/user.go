package domain

import (
	"database/sql"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// UserRequest holds the input request from the client
type UserRequest struct {
	Size   int64  `query:"size"`
	Offset int64  `query:"offset"`
	Name   string `query:"name"`
}

// UserList holds the set of all list data
// It helps to sort the data with user id
type UserList []*List

func (s UserList) Len() int           { return len(s) }
func (s UserList) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s UserList) Less(i, j int) bool { return s[i].ID < s[j].ID }

// List holds the user detail fields
type List struct {
	UserDetail `bson:",inline"`
	Auth       []Auth   `json:"auth,omitempty" bson:"auth"`
	Access     []Access `json:"access,omitempty" bson:"access"`
}

// UserDetail holds the profile basic information
type UserDetail struct {
	// separate fields to get the ID -> moving to common field ID
	IDMysql int64              `json:"-" bson:"-"`
	IDBson  primitive.ObjectID `json:"-" bson:"_id"`

	ID            string `json:"id" bson:"-"`
	Name          string `json:"name" bson:"name"`
	Email         string `json:"email" bson:"email"`
	Gender        string `json:"gender" bson:"gender"`
	Phone         string `json:"phone" bson:"phone"`
	MartialStatus string `json:"martialStatus" bson:"martialStatus"`
	Password      string `json:"-" bson:"password"`
}

// Auth holds the profile auth
type Auth struct {
	Provider string `json:"provider" bson:"provider"`
	AuthID   string `json:"authID" bson:"authID"`
	Name     string `json:"name" bson:"name"`

	Auth struct {
		AuthID       sql.NullString `json:"-" bson:"-"`
		AuthProvider sql.NullString `json:"-" bson:"-"`
		AuthName     sql.NullString `json:"-" bson:"-"`
	} `json:"-" bson:"-"`
}

// Access holds the profile access
type Access struct {
	Name   string `json:"name" bson:"name"`
	View   bool   `json:"view" bson:"view"`
	Edit   bool   `json:"edit" bson:"edit"`
	Remove bool   `json:"delete" bson:"delete"`

	Access struct {
		AccessName sql.NullString `json:"-" bson:"-"`
		View       sql.NullBool   `json:"-" bson:"-"`
		Edit       sql.NullBool   `json:"-" bson:"-"`
		Remove     sql.NullBool   `json:"-" bson:"-"`
	} `json:"-" bson:"-"`
}
