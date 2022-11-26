package utils

import (
	"github.com/ilyakaznacheev/cleanenv"
	"go.uber.org/zap"
	"io"
	"os"
)

type TurnConfig struct {
	ExternalIp    string `yaml:"external_ip" env:"EXTERNAL_IP" env-descriptions:"External IP of a L4 LoadBalancer"`
	TurnPort      string `yaml:"port" env:"TURN_PORT" env-descriptions:"Listening port of a coturn server"`
	UserName      string `yaml:"user_name" env:"USER_NAME" env-descriptions:"username for the lt-cred-mech"`
	TTLSeconds    uint64 `yaml:"ttl_seconds" env:"TTL_SECONDS" env-descriptions:"Expiration of password"`
	UseLtCredMech bool `yaml:"use_lt_cred_mech" env:"USE_LT_CRED_MECH" env-descriptions:"whether to use lt-cred-mech"`
}

type TurnSecret struct {
	Password         string `yaml:"password" env:"PASSWORD" env-descriptions:"password for the lt-cred-mech"`
	TurnSharedSecret string `yaml:"shared_secret" env:"TURN_SHARED_SECRET" env-descriptions:"Shared encrypted secret for Turn REST API"`
}

// the following struct is to map from hashicorp vault
type TurnSecretData struct {
	Data TurnSecret `json:"data"`
}

type AppConfig struct {
	TurnConfig TurnConfig     `yaml:"turn_config"`
	TurnSecret TurnSecretData `json:"data"`
}

// parse multiple config files and combine to one global configuraion object
func ParseConfigFiles(files ...string) (*AppConfig, error) {
	var cfg AppConfig

	for i := 0; i < len(files); i++ {
		err := cleanenv.ReadConfig(files[i], &cfg)
		if err != nil {
			return nil, err
		}
	}

	return &cfg, nil
}

// generate the json file that cleanenv lib could support
func GenerateSecretJson(secretPath string, logger *zap.SugaredLogger) string {
	original, err := os.Open((secretPath))
	if err != nil {
		logger.Fatal(err)
	}
	defer original.Close()
	newPath := "./secret.json"
	new, err := os.Create(newPath)
	if err != nil {
		logger.Fatal(err)
	}
	defer new.Close()

	bytesWritten, err := io.Copy(new, original)
	if err != nil {
		logger.Fatal(err)
	}
	logger.Infof("Bytes Written: %d\n", bytesWritten)
	return newPath
}
