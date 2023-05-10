package config

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

type Config struct {
	DBHost     string `mapstructure:"DB_HOST"`
	DBName     string `mapstructure:"DB_NAME"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
}

var envs = []string{
	"DB_HOST", "DB_NAME", "DB_USER", "DB_PORT", "DB_PASSWORD",
}

func LoadConfig() (Config, error) {
	var config Config

	viper.AddConfigPath("./")   //The AddConfigPath() function is used to specify the directories that Viper should search for configuration files. In this case, the function is called with the argument "./", which adds the current directory to the search path.
	viper.SetConfigFile(".env") //The SetConfigFile() function is used to specify the name of the configuration file. In this case, the function is called with the argument ".env", which sets the name of the configuration file to .env.
	viper.ReadInConfig()        //the ReadInConfig() function is called to read the configuration file and set the values in Viper. This function searches for the configuration file in the directories specified by AddConfigPath() and attempts to read it using the format specified by SetConfigType(). If the configuration file is found and read successfully, Viper sets the values in its internal data structure.

	fmt.Println("just studying viper, no need of this line", viper.GetString("DB_HOST")) //Viper uses this internal data structure to provide access to the configuration values via its API. For example, you can access a configuration value by calling viper.GetString("key") or viper.GetInt("key"), where "key" is the name of the configuration key.

	fmt.Println("config struct is 1", config)
	for _, env := range envs {
		if err := viper.BindEnv(env); err != nil {
			return config, err
		}
	}

	fmt.Println("config struct is 2", config)
	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}

	fmt.Println("config struct is 3", config)

	if err := validator.New().Struct(&config); err != nil {
		return config, err
	}

	fmt.Println("config struct is 4", config)

	return config, nil
}
