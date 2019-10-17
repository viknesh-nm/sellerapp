/*Package conf is the configuration infrastructure for the applciation.
It provides utilites to read environment variables and Initializes the global configuration object.
*/
package conf

import "os"

// Vars represents the environment variables used by this App during Startup.
type Vars struct {
	Port string

	MySQL struct {
		User     string
		Password string
		DSN      string
	}

	Mongo struct {
		DSN string
	}
}

// Read parses the configuration from environment variables and populates the Conf struct
func Read() (*Vars, error) {
	vars := &Vars{}
	vars.Port = os.Getenv("SELLERAPP_PORT")
	vars.MySQL.DSN = os.Getenv("SELLERAPP_MYSQL_DSN")
	vars.MySQL.User = os.Getenv("SELLERAPP_MYSQL_USER")
	vars.MySQL.Password = os.Getenv("SELLERAPP_MYSQL_PASSWORD")
	vars.Mongo.DSN = os.Getenv("SELLERAPP_MONGO_DSN")
	return vars, nil
}

type configuration struct {
	*Vars
}

// Config is the global configuration objects for the App
//	It contains Backend Repository, Environments and Loggers
var Config configuration

//Init initializes New Global Configuration
func Init(c *Vars) {
	Config.Vars = c
}
