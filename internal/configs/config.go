package configs

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

type Config struct {
	Server struct {
		Port string `yaml:"port", envconfig:"SERVER_PORT"`
		Host string `yaml:"host", envconfig:"SERVER_HOST"`
	} `yaml:"server"`
}

// обработка ошибок при работе с конфигом
func processError(err error) {
	fmt.Println(err)
	os.Exit(2)
}

// чтение и парсинг файла .yml конфига
func ReadFile(cfg *Config) {
	s := string(os.PathSeparator)
	env := "dev"
	configName := env + "-config.yml"
	configPath := "." + s + "internal" + s + "configs" + s + configName
	f, err := os.Open(configPath)
	if err != nil {
		processError(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
	if err != nil {
		processError(err)
	}
}

// чтение и парсинг файла .env конфига
/*func ReadEnv(cfg *Config) {
	err := envconfig.Process("", cfg)
	if err != nil {
		processError(err)
	}
}*/
