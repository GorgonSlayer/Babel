package config

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"radiola.co.nz/babel/src/util/logger"
)

// RetrieveEnvironmentalVariables /** This file should pick up our environmental variables for configuration. **/
//Retrieve the following.
//jwt string
//fleetPinUrl string
//host string
//port int
//key string
//agency string

// FleetPinConfig /** Configuration Struct for storing our config in memory. **/
type FleetPinConfig struct {
	Jwt         string
	FleetPinUrl string
	Host        string
	Port        int
	Key         string
	Agency      string
}

// GetConfigFromFile /** Return a struct of config variables taken from the .env file.**/
func GetConfigFromFile(l logger.Logger) FleetPinConfig { //For Local Dev or Dev
	viper.AddConfigPath("./config")
	viper.SetConfigName("config")
	err := viper.ReadInConfig()
	if err != nil {
		l.Zap.Error("Error Reading the Config File for FleetPin", zap.Any("Error", err.Error()))
	}
	var fpc FleetPinConfig
	fpc.Jwt = viper.GetString("JWT")
	fpc.FleetPinUrl = viper.GetString("FLEETPINURL")
	fpc.Host = viper.GetString("HOST")
	fpc.Port = viper.GetInt("PORT")
	fpc.Key = viper.GetString("KEY")
	fpc.Agency = viper.GetString("AGENCY")
	return fpc
}

// GetConfigFromEnv /** Return a struct of config variables taken from the environment during CICD deployment. **/
func GetConfigFromEnv() { //For Prod.
	//TODO: Implement this.
}
