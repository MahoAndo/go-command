package conf

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

type (
	Config struct {
		BookDB     DB
	}

	DB struct {
		Hostname string
		Port     int
		User     string
		Password string
	}

)

func Load() (*Config, error) {
	v := viper.New()
	v.SetConfigName("application")
	confPath := getConfPath()
	v.AddConfigPath(confPath)

	//Dynamic loading of configuration files
	v.AutomaticEnv()

	err := v.ReadInConfig()
	if err != nil {
		return nil, err
	}

	c := &Config{}
	err = v.Unmarshal(c)
	return c, err
}

// get setting file (application.toml)
func getConfPath() string {
	exe, err := os.Executable()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	
	return filepath.Dir(exe)
}
