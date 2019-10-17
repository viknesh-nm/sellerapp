/*
Package repo provides definitions for various Repositories that has been available
to perform database operations.
*/
package repo

import "github.com/viknesh-nm/sellerapp/domain"

// User Repository defines a list functions available to perform Operations on User Repository
type User interface {
	UserList(domain.UserRequest) ([]*domain.List, error)
	UserListMongo(domain.UserRequest) (domain.UserList, error)
	MongoConvertions() error
}
