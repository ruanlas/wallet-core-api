package configs

import (
	"github.com/caarlos0/env/v11"
)

var App struct {
	ServiceHost string `env:"SERVICE_HOST,required"`
	Database    struct {
		Host     string `env:"DATABASE_HOST,required"`
		Port     string `env:"DATABASE_PORT,required"`
		Name     string `env:"DATABASE_NAME,required"`
		Username string `env:"DATABASE_USERNAME,required"`
		Password string `env:"DATABASE_PASSWORD,required"`
	}
	Idp struct {
		Host             string `env:"IDP_HOST,required"`
		Port             string `env:"IDP_PORT,required"`
		MainRealm        string `env:"IDP_MAIN_REALM,required"`
		UserAdmin        string `env:"IDP_USER_ADMIN,required"`
		PasswordAdmin    string `env:"IDP_PASSWORD_ADMIN,required"`
		Realm            string `env:"IDP_REALM,required"`
		ClientIdentifier string `env:"IDP_CLIENT_IDENTIFIER,required"`
		ClientSecret     string `env:"IDP_CLIENT_SECRET,required"`
	}
	Monitoring struct {
	}
}

func Init() {
	err := env.Parse(&App)
	if err != nil {
		panic(err)
	}
}
