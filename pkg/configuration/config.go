package configuration

import (
	"os"

	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
)

// Configuration
type Configuration struct {
	ListenPort      string `default:":9670" split_words:"true"`
	RootURL         string `default:"/tokped" split_words:"true"`
	OriginHost      string `default:"tokped" split_words:"true"`
	Timeout         int64  `default:"60000" split_words:"true"`
	Addr            string `default:"localhost" split_words:"true"`
	MariaDBAddr     string `default:"localhost" split_words:"true"`
	MariaDBPort     string `default:"3306" split_words:"true"`
	MariaDBUser     string `default:"root" split_words:"true"`
	MariaDBPassword string `default:"" split_words:"true"`
	MariaDBDatabase string `default:"order_history" split_words:"true"`
	LimitQuery      int64  `default:"10" split_words:"true"`
	ClientSecret    string `default:"PlatformSecretdev" split_words:"true"`
	TokenLifeTime   int64  `default:"10800" split_words:"true"`
	JobListUrl      string `default:"http://dev3.dansmultipro.co.id/api/recruitment/positions.json" split_words:"true"`
	JobDetaillUrl   string `default:"http://dev3.dansmultipro.co.id/api/recruitment/positions/" split_words:"true"`
}

// Config .
var Config Configuration

// LoadConfig .
func LoadConfig() {
	if err := envconfig.Process("TP", &Config); err != nil {
		log.Error(err)
		os.Exit(1)
	}
}
