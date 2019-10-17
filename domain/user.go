package domain

type UserRequest struct {
	Size   int64  `query:"size"`
	Offset int64  `query:"offset"`
	Name   string `query:"name"`
}

// UserList ...
type UserList []*List

func (s UserList) Len() int           { return len(s) }
func (s UserList) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s UserList) Less(i, j int) bool { return s[i].ID < s[j].ID }

type List struct {
	UserDetail `bson:",inline"`
	Auth       []Auth   `json:"auth,omitempty" bson:"auth"`
	Access     []Access `json:"access,omitempty" bson:"access"`
}
type UserDetail struct {
	ID            int64  `json:"id" bson:"id"`
	Name          string `json:"name" bson:"name"`
	Email         string `json:"email" bson:"email"`
	Gender        string `json:"gender" bson:"gender"`
	Phone         string `json:"phone" bson:"phone"`
	MartialStatus string `json:"martialStatus" bson:"martialStatus"`
}

type Auth struct {
	ID       int64  `json:"-" bson:"id"`
	Provider string `json:"provider,omitempty" bson:"provider"`
	AuthID   string `json:"authID,omitempty" bson:"authID"`
	Name     string `json:"name,omitempty" bson:"name"`
	Password string `json:"-" bson:"password"`
}

type Access struct {
	ID     int64  `json:"-" bson:"id"`
	Name   string `json:"name" bson:"name"`
	View   bool   `json:"view" bson:"view"`
	Edit   bool   `json:"edit" bson:"edit"`
	Remove bool   `json:"delete" bson:"delete"`
}
