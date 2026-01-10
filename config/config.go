package config

import (
	"fmt"
	"log"
	"path/filepath"
	"runtime"

	libConfig "github.com/oceaninov/naeco-go-lib/config"
	"github.com/spf13/viper"
)

func readViperConfig() (obj ConfigObject) {
	v := viper.New()
	v.AddConfigPath(".")
	v.AddConfigPath("./params")
	v.AddConfigPath("/opt/params")
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)
	v.AddConfigPath(fmt.Sprintf("%s/../params", basepath))
	v.SetConfigName("app")
	v.SetConfigType("env")
	v.AutomaticEnv()

	err := v.ReadInConfig()
	if err == nil {
		log.Printf("Using config file: %s", v.ConfigFileUsed())
		err = v.Unmarshal(&obj)
		if err != nil {
			log.Panicf("Config error: %s", err)
		}
	} else {
		log.Printf("Config file not found or error: %s", err)
		// load from env as fallback
		libConfig.New(&obj)
	}

	return
}

// Config return provider so that you can read config anywhere
func Config() ConfigObject {
	return readViperConfig()
}
