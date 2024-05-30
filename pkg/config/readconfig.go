package config

import (
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"lavaprovider-monitor/pkg/log"
)

//read alert config

type AlertConfig struct {
	Alert struct {
		Enable bool `mapstructure:"enable"`
	}

	Discord DiscordAlertConfig `mapstructure:"discord"`
}

type DiscordAlertConfig struct {
	Enable      bool     `mapstructure:"enable"`
	Webhook     string   `mapstructure:"webhook"`
	Alertuserid []string `mapstructure:"alertuserid"`
	Alertroleid []string `mapstructure:"alertroleid"`
}

// GetAlertConfig Read alert config file and return in alertConfig
func GetAlertConfig() AlertConfig {
	viper.SetConfigName("alert")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	var config AlertConfig
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		log.Log.Error("read alert config file alert.toml.example failed", zap.Error(err))
		panic(fmt.Errorf("read alert config file alert.toml.example file failed: %s \n", err))
	}

	// unmarshal the config
	err = viper.Unmarshal(&config)
	if err != nil {
		log.Log.Error("unmarshal alert config file alert.toml.example failed", zap.Error(err))
		panic(fmt.Errorf("unmarshal alert config file alert.toml.example file failed: %s \n", err))
	}
	return config

}

type LavaProviderMonitorConfig struct {
	LavaGrpc            string   `mapstructure:"lavagrpc"`
	ChainID             string   `mapstructure:"chainid"`
	LavaProviderAddress string   `mapstructure:"lavaprovideraddress"`
	Chains              []string `mapstructure:"chains"`
}

// GetLavaProviderMonitorConfig Read lava provider monitor config file and return in LavaProviderMonitorConfig
func GetLavaProviderMonitorConfig() LavaProviderMonitorConfig {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	var config LavaProviderMonitorConfig
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		log.Log.Error("read lava provider monitor config config.toml.example failed", zap.Error(err))
		panic(fmt.Errorf("read lava provider monitor config config.toml.example failed: %s \n", err))
	}

	// unmarshal the config
	err = viper.Unmarshal(&config)
	if err != nil {
		log.Log.Error("unmarshal lava provider monitor config config.toml.example failed", zap.Error(err))
		panic(fmt.Errorf("unmarshal lava provider monitor config config.toml.example failed: %s \n", err))
	}

	return config

}
