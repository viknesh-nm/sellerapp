/*Package backend provides various Respoistories to handle database operations related to Domain
  It has various implementations to satisfy backend/repo interafaces.
*/
package backend

import (
	"github.com/jinzhu/gorm"
	"github.com/viknesh-nm/sellerapp/backend/repo"
	"go.mongodb.org/mongo-driver/mongo"
)

//Infra defines the infrastructure for all the database connection and logger objects.
type Infra struct {
	MongoDB *mongo.Client
	MySQLDB *gorm.DB
}

// Repository to handle domain specific database operations
var (
	User repo.User
)

// Init all the Repository with an Infrastructure
func Init(infra *Infra) {
	User = &users{infra}
}
