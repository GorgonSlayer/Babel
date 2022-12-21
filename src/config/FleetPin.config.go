package config

import "github.com/spf13/viper"

// RetrieveEnvironmentalVariables /** This file should pick up our environmental variables for configuration. **/
//Retrieve the following.
//jwt string
//fleetPinUrl string
//host string
//port int
//key string
//agency string

func GetConfigFromFile() { //For Local Dev or Dev
	viper.SetConfigName("FleetPinConfig")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
}

func GetConfigFromEnv() { //For Prod.

}
