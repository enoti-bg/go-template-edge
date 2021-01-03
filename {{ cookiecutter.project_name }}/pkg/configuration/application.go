package configuration

import (
	"fmt"
	"os"
	"path/filepath"

	ozzo "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/spf13/viper"
)

// RepositoryConfiguration contains the adapter, as well as options for creating
// the specific repository type, if applicable.
type RepositoryConfiguration struct {
	Adapter string
	Options map[string]interface{}
}

// AppConfiguration contains application specific data.
type AppConfiguration struct {
	Env        string
	LogLevel   string `mapstructure:"log"`
	Port       int    `mapstructure:"port"`
	Repository *RepositoryConfiguration
}

// Validate performs basic validation on the contents of a configuration.
func (c *AppConfiguration) Validate() error {
	return ozzo.ValidateStruct(
		c,
		ozzo.Field(&c.Env, ozzo.Required, ozzo.In("acceptance", "development", "production", "test")),
		ozzo.Field(&c.Port, ozzo.Required, ozzo.Min(1), ozzo.Max(65535)),
	)
}

// LoadYAML loads configuration from a mix of environment variables and a yaml file and coerce to
// given target structure.
func LoadYAML(targetFormat interface{}, targetPath *string, targetFile *string, fromEnv []string) error {
	viper.SetDefault("env", "default")
	viper.SetEnvPrefix("service")
	viper.BindEnv("env")

	for _, suffix := range fromEnv {
		viper.BindEnv(suffix)
	}

	var configDirectory, configName = ".", viper.GetString("env")

	if targetPath != nil {
		configDirectory = filepath.Clean(*targetPath)
	}

	if targetFile != nil {
		configName = *targetFile
	}

	_, err := os.Stat(fmt.Sprintf("%s%s%s.yml", configDirectory, string(os.PathSeparator), configName))

	if err != nil {
		return err
	}

	viper.AddConfigPath(configDirectory)
	viper.SetConfigName(configName)

	err = viper.ReadInConfig()
	if err != nil {
		return err
	}

	err = viper.Unmarshal(targetFormat)
	if err != nil {
		return err
	}

	return nil
}
