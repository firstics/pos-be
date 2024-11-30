package config

import (
	"encoding/json"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

type Config struct {
	API APIConfig `mapstructure:"api" json:"api" validate:"required"`
}

type APIConfig struct {
	Log LogConfig `mapstructure:"log" json:"log" validate:"required"`
}

type LogConfig struct {
	Tracing bool      `mapstructure:"tracing" json:"tracing"`
	Zap     ZapConfig `mapstructure:"zap" json:"zap" validate:"required"`
}

type ZapConfig struct {
	Level            string        `mapstructure:"level" json:"level" validate:"required"`
	Encoding         string        `mapstructure:"encoding" json:"encoding" validate:"required"`
	OutputPaths      []string      `mapstructure:"outputPaths" json:"outputPaths" validate:"required"`
	ErrorOutputPaths []string      `mapstructure:"errorOutputPaths" json:"errorOutputPaths" validate:"required"`
	EncoderConfig    EncoderConfig `mapstructure:"encoderConfig" json:"encoderConfig" validate:"required"`
}

type EncoderConfig struct {
	MessageKey   string `mapstructure:"messageKey" json:"messageKey" validate:"required"`
	LevelKey     string `mapstructure:"levelKey" json:"levelKey" validate:"required"`
	LevelEncoder string `mapstructure:"levelEncoder" json:"levelEncoder" validate:"required"`
}

func LoadConfig() (Config, error) {
	var config Config

	viper.SetConfigName("config")
	viper.AddConfigPath("./")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		return config, err
	}

	if err := viper.Unmarshal(&config, func(dc *mapstructure.DecoderConfig) {
		dc.DecodeHook = mapstructure.ComposeDecodeHookFunc(
			StringToStructHookFunc(),
			StringToSliceWithBracketHookFunc(),
			dc.DecodeHook)
	}); err != nil {
		return config, err
	}

	if err := validator.New().Struct(&config); err != nil {
		return config, err
	}

	return config, nil
}

func StringToSliceWithBracketHookFunc() mapstructure.DecodeHookFunc {
	return func(
		f reflect.Kind,
		t reflect.Kind,
		data interface{}) (interface{}, error) {
		if f != reflect.String || t != reflect.Slice {
			return data, nil
		}

		raw := data.(string)
		if raw == "" {
			return []string{}, nil
		}
		var slice []json.RawMessage
		err := json.Unmarshal([]byte(raw), &slice)
		if err != nil {
			return data, nil
		}

		var strSlice []string
		for _, v := range slice {
			strSlice = append(strSlice, string(v))
		}
		return strSlice, nil
	}
}

func StringToStructHookFunc() mapstructure.DecodeHookFunc {
	return func(
		f reflect.Type,
		t reflect.Type,
		data interface{},
	) (interface{}, error) {
		if f.Kind() != reflect.String ||
			(t.Kind() != reflect.Struct && !(t.Kind() == reflect.Pointer && t.Elem().Kind() == reflect.Struct)) {
			return data, nil
		}
		raw := data.(string)
		var val reflect.Value
		// Struct or the pointer to a struct
		if t.Kind() == reflect.Struct {
			val = reflect.New(t)
		} else {
			val = reflect.New(t.Elem())
		}

		if raw == "" {
			return val, nil
		}
		err := json.Unmarshal([]byte(raw), val.Interface())
		if err != nil {
			return data, nil
		}
		return val.Interface(), nil
	}
}
