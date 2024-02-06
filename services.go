package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	InfuraURL     string `mapstructure:"INFURA_URL"`
	AlchemyURL    string `mapstructure:"ALCHEMY_URL"`
	ChainstackURL string `mapstructure:"CHAINSTACK_URL"`
	TednerlyURL   string `mapstructure:"TENDERLY_URL"`

	InfuraKey     string `mapstructure:"INFURA_KEY"`
	AlchemyKey    string `mapstructure:"ALCHEMY_KEY"`
	ChainstackKey string `mapstructure:"CHAINSTACK_KEY"`
	TenderlyKey   string `mapstructure:"TENDERLY_KEY"`
}

func loadConfig() (config *Config, err error) {
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")

	err = viper.ReadInConfig()
	if err != nil {
		log.Fatal("Can't read the config file: ", err)
		return nil, fmt.Errorf("Can't read the config file: %e", err)
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatal("Unable to Unmarshal Config file: ", err)
		return nil, fmt.Errorf("Unable to Unmarshal Config file: %e", err)
	}
	return config, nil
}

func initClients(config *Config) ([]EthereumClient, error) {
	var clients []EthereumClient

	clients = append(clients, NewInfuraClient(config.InfuraURL, config.InfuraKey))
	clients = append(clients, NewAlchemyClient(config.AlchemyURL, config.AlchemyKey))
	clients = append(clients, NewChainStackClient(config.ChainstackURL, config.ChainstackKey))
	clients = append(clients, NewTenderlyClient(config.TednerlyURL, config.TenderlyKey))
	return clients, nil
}
