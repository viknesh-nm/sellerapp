package backend

import (
	"context"
	"fmt"
	"sort"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/viknesh-nm/sellerapp/domain"
	"go.mongodb.org/mongo-driver/bson"
)

type users struct {
	*Infra
}

// UserList returns the entire user profile list
// Search key by name has been implemented so far
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

// UserAccessList returns the access data with the key of profile detail
func (u users) UserAccessList(req domain.UserRequest) (map[domain.UserDetail][]domain.Access, error) {
	var selectField = []string{
		"profile.id",
		"profile.name",
		"profile.email",
		"profile.gender",
		"profile.martial_status",
		"profile.phone",
		"profile.password",
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
		tempAccess := domain.Access{}
		temp := domain.UserDetail{}
		err := rows.Scan(
			&temp.IDMysql,
			&temp.Name,
			&temp.Email,
			&temp.Gender,
			&temp.MartialStatus,
			&temp.Phone,
			&temp.Password,
			&tempAccess.Access.AccessName,
			&tempAccess.Access.View,
			&tempAccess.Access.Edit,
			&tempAccess.Access.Remove,
		)
		if err != nil {
			return nil, err
		}

		temp.ID = fmt.Sprintf("%d", temp.IDMysql)

		if tempAccess.Access.AccessName.Valid &&
			tempAccess.Access.View.Valid && tempAccess.Access.Edit.Valid && tempAccess.Access.Remove.Valid {
			tempAccess.Name = tempAccess.Access.AccessName.String
			tempAccess.View = tempAccess.Access.View.Bool
			tempAccess.Edit = tempAccess.Access.Edit.Bool
			tempAccess.Remove = tempAccess.Access.Remove.Bool

			accessMap[temp] = append(
				accessMap[temp],
				tempAccess,
			)
		} else {
			accessMap[temp] = make([]domain.Access, 0)
		}

	}

	return accessMap, nil
}

// UserAuthList eturns the auth data with the key of profile detail
func (u users) UserAuthList(req domain.UserRequest) (map[domain.UserDetail][]domain.Auth, error) {
	var selectField = []string{
		"profile.id",
		"profile.name",
		"profile.email",
		"profile.gender",
		"profile.martial_status",
		"profile.phone",
		"profile.password",
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
		temp := domain.UserDetail{}
		tempAuth := domain.Auth{}
		err := rows.Scan(
			&temp.IDMysql,
			&temp.Name,
			&temp.Email,
			&temp.Gender,
			&temp.MartialStatus,
			&temp.Phone,
			&temp.Password,
			&tempAuth.Auth.AuthID,
			&tempAuth.Auth.AuthProvider,
			&tempAuth.Auth.AuthName,
		)
		if err != nil {
			return nil, err
		}

		temp.ID = fmt.Sprintf("%d", temp.IDMysql)

		if tempAuth.Auth.AuthID.Valid &&
			tempAuth.Auth.AuthName.Valid && tempAuth.Auth.AuthProvider.Valid {
			tempAuth.AuthID = tempAuth.Auth.AuthID.String
			tempAuth.Name = tempAuth.Auth.AuthName.String
			tempAuth.Provider = tempAuth.Auth.AuthProvider.String

			authMap[temp] = append(
				authMap[temp],
				tempAuth,
			)
		} else {
			authMap[temp] = make([]domain.Auth, 0)
		}

	}

	return authMap, nil
}

// MongoConvertions trasfers the aggregated list -> mongoDB "user" collection
func (u users) MongoConvertions() error {
	list, err := u.UserList(domain.UserRequest{})
	if err != nil {
		return err
	}

	dataInsertion := make([]interface{}, 0)
	for _, l := range list {
		l.IDBson = primitive.NewObjectID()
		dataInsertion = append(dataInsertion, l)
	}

	collection := u.MongoDB.Database("mydb").Collection("users")

	_, err = collection.InsertMany(context.Background(), dataInsertion)
	if err != nil {
		return err
	}

	return nil
}

// UserListMongo retuns the user list and error from mongoDB
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

		temp.ID = temp.IDBson.Hex()

		list = append(list, &temp)
	}

	return list, nil
}
