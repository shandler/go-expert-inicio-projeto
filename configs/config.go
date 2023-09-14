package configs

import (
	"github.com/go-chi/jwtauth"
	"github.com/spf13/viper"
)

type conf struct {
	DBDriver      string `mapstruture:"DB_DRIVER"`
	DBHost        string `mapstruture:"DB_HOST"`
	DBPort        string `mapstruture:"DB_PORT"`
	DBUser        string `mapstruture:"DB_USER"`
	DBPassword    string `mapstruture:"DB_PASSWORD"`
	DBName        string `mapstruture:"DB_NAME"`
	WebServerPort string `mapstruture:"WEB_SERVER_PORT"`
	JWTSecret     string `mapstruture:"JWT_SECRET"`
	JWTExpiresIn  string `mapstruture:"JWT_EXPIRESIN"`
	TokenAuth     *jwtauth.JWTAuth
}

func LoadConfig(path string) (*conf, error) {
	var cfg *conf
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}
	cfg.TokenAuth = jwtauth.New("HS256", []byte(cfg.JWTSecret), nil)
	return cfg, err

}
