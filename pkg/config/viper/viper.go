package viper

import (
	"reflect"
	"strings"

	"github.com/spf13/viper"

	"github.com/dittonetwork/executor-avs/pkg/log"
)

func Load(confPath string, receiver interface{}) {
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yml")

	configType := "default"
	if confPath != "" {
		viper.SetConfigFile(confPath)
		configType = "supplied"
	}

	err := viper.ReadInConfig()
	if err != nil {
		log.Default().Fatal("Read config",
			log.String("config_type", configType),
			log.Err(err),
		)
	}

	log.Default().Info("Viper using config",
		log.String("config", viper.ConfigFileUsed()),
		log.String("config_type", configType),
	)

	bindEnvs(reflect.ValueOf(receiver))
	if err = viper.Unmarshal(receiver); err != nil {
		log.Panic("Error Unmarshal Viper Config File", log.Err(err))
	}
}

//nolint:exhaustive // later
func bindEnvs(v reflect.Value, parts ...string) {
	if v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return
		}

		bindEnvs(v.Elem(), parts...)
		return
	}

	ift := v.Type()
	for i := 0; i < ift.NumField(); i++ {
		val := v.Field(i)
		t := ift.Field(i)
		tv, ok := t.Tag.Lookup("mapstructure")
		if !ok {
			continue
		}

		switch val.Kind() {
		case reflect.Struct:
			bindEnvs(val, append(parts, tv)...)
		default:
			if err := viper.BindEnv(strings.Join(append(parts, tv), ".")); err != nil {
				log.Default().Fatal("Bind env", log.Err(err))
			}
		}
	}
}
