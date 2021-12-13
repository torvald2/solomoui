package config

import (
	"fmt"
	"os"
	"reflect"
	"sync"

	"github.com/joho/godotenv"
)

// TODO  - marshaller with validator
type appConfig struct {
	DBConnectionString string `env:"DB_CONNECTION_STRING"`
	TCP_Port           string `env:"APP_PORT"`
}

var conf appConfig
var once sync.Once

func GetConfig() *appConfig {
	once.Do(setUp)
	return &conf
}

func setUp() {
	if env := os.Getenv("IS_DEV"); env != "PROD" {
		err := godotenv.Load()
		if err != nil {
			fmt.Println("No .env files found. Using real environment")
		}

	}
	v := reflect.ValueOf(&conf).Elem()
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		f := t.Field(i)
		varname, _ := f.Tag.Lookup("env")
		env, ok := os.LookupEnv(varname)
		if ok {
			v.Field(i).SetString(env)
		}

	}
}
