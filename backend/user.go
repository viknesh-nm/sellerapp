package backend

import (
	"context"
	"database/sql"
	"sort"

	"github.com/viknesh-nm/sellerapp/domain"
	"go.mongodb.org/mongo-driver/bson"
)

type users struct {
	*Infra
}

type AuthAccess struct {
	Auth struct {
		AuthID       sql.NullString
		AuthProvider sql.NullString
		AuthName     sql.NullString
	}
	AuthID       string
	AuthProvider string
	AuthName     string

	Access struct {
		AccessName sql.NullString
		View       sql.NullBool
		Edit       sql.NullBool
		Remove     sql.NullBool
	}

	AccessName string
	View       bool
	Edit       bool
	Remove     bool
}

// UserList ...
func (u users) UserList(req domain.UserRequest) ([]*domain.List, error) {

	list := make(domain.UserList, 0)

	access, err := u.UserAccessList(req)
	if err != nil {
		return nil, err
	}

	auth, err := u.UserAuthList(req)
	if err != nil {
		return nil, err
	}

	for user, accessDetail := range access {
		list = append(list,
			&domain.List{
				UserDetail: user,
				Access:     accessDetail,
				Auth:       auth[user],
			})
	}

	sort.Sort(list)
	return list, nil
}

// UserAccessList ...
func (u users) UserAccessList(req domain.UserRequest) (map[domain.UserDetail][]domain.Access, error) {
	var selectField = []string{
		"profile.id",
		"profile.name",
		"profile.email",
		"profile.gender",
		"profile.martial_status",
		"profile.phone",
		"access.name AS access_name",
		"view",
		"edit",
		"remove",
	}
	db := u.MySQLDB.Table("profile").
		Select(selectField).
		Joins("LEFT JOIN access_profile_mapping apm ON apm.profile_id = profile.id").
		Joins("LEFT JOIN access ON access.id = apm.access_id")

	if req.Name != "" {
		db = db.Where("profile.name LIKE ?", "%"+req.Name+"%")
	}

	rows, err := db.Rows()
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	accessMap := make(map[domain.UserDetail][]domain.Access, 0)
	for rows.Next() {
		g := AuthAccess{}
		tp := domain.UserDetail{}
		err := rows.Scan(
			&tp.ID,
			&tp.Name,
			&tp.Email,
			&tp.Gender,
			&tp.MartialStatus,
			&tp.Phone,
			&g.Access.AccessName,
			&g.Access.View,
			&g.Access.Edit,
			&g.Access.Remove,
		)
		if err != nil {
			return nil, err
		}

		if g.Access.AccessName.Valid &&
			g.Access.View.Valid && g.Access.Edit.Valid && g.Access.Remove.Valid {
			g.AccessName = g.Access.AccessName.String
			g.View = g.Access.View.Bool
			g.Edit = g.Access.Edit.Bool
			g.Remove = g.Access.Remove.Bool

			accessMap[tp] = append(
				accessMap[tp],
				domain.Access{
					Name:   g.AccessName,
					View:   g.View,
					Edit:   g.Edit,
					Remove: g.Remove,
				},
			)
		} else {
			accessMap[tp] = make([]domain.Access, 0)
		}

	}

	return accessMap, nil
}

// UserAuthList ...
func (u users) UserAuthList(req domain.UserRequest) (map[domain.UserDetail][]domain.Auth, error) {
	var selectField = []string{
		"profile.id",
		"profile.name",
		"profile.email",
		"profile.gender",
		"profile.martial_status",
		"profile.phone",
		"auth_id",
		"auth_provider",
		"auth.name AS auth_name",
	}
	db := u.MySQLDB.Table("profile").
		Select(selectField).
		Joins("LEFT JOIN auth ON auth.profile_id = profile.id")

	if req.Name != "" {
		db = db.Where("profile.name LIKE ?", "%"+req.Name+"%")
	}

	rows, err := db.Rows()
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	authMap := make(map[domain.UserDetail][]domain.Auth, 0)
	for rows.Next() {
		g := AuthAccess{}
		tp := domain.UserDetail{}
		err := rows.Scan(
			&tp.ID,
			&tp.Name,
			&tp.Email,
			&tp.Gender,
			&tp.MartialStatus,
			&tp.Phone,
			&g.Auth.AuthID,
			&g.Auth.AuthProvider,
			&g.Auth.AuthName,
		)
		if err != nil {
			return nil, err
		}

		if g.Auth.AuthID.Valid &&
			g.Auth.AuthName.Valid && g.Auth.AuthProvider.Valid {
			g.AuthID = g.Auth.AuthID.String
			g.AuthName = g.Auth.AuthName.String
			g.AuthProvider = g.Auth.AuthProvider.String

			authMap[tp] = append(
				authMap[tp],
				domain.Auth{
					Provider: g.AuthProvider,
					AuthID:   g.AuthID,
					Name:     g.AuthName,
				},
			)
		} else {
			authMap[tp] = make([]domain.Auth, 0)
		}

	}

	return authMap, nil
}

// MongoConvertions ...
func (u users) MongoConvertions() error {
	list, err := u.UserList(domain.UserRequest{})
	if err != nil {
		return err
	}

	dataInsertion := make([]interface{}, 0)
	for _, l := range list {
		dataInsertion = append(dataInsertion, l)
	}

	collection := u.MongoDB.Database("mydb").Collection("users")

	_, err = collection.InsertMany(context.Background(), dataInsertion)
	if err != nil {
		return err
	}

	return nil
}

// UserListMongo ...
func (u users) UserListMongo(req domain.UserRequest) (domain.UserList, error) {
	var (
		filter = bson.M{}
		ctx    = context.Background()
		list   = make(domain.UserList, 0)
	)

	if req.Name != "" {
		filter = bson.M{
			"name": bson.M{
				"$regex": req.Name,
			},
		}
	}

	collection := u.MongoDB.Database("mydb").Collection("users")
	rows, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	for rows.Next(ctx) {
		temp := domain.List{}
		err := rows.Decode(&temp)
		if err != nil {
			return nil, err
		}

		list = append(list, &temp)
	}

	return list, nil
}
