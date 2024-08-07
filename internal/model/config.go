package model

type Config struct {
	Port                 string `yaml:"port"`
	DbUsername           string `yaml:"db-username"`
	DbPassword           string `yaml:"db-password"`
	DbName               string `yaml:"db-name"`
	DbHost               string `yaml:"db-host"`
	DbPort               string `yaml:"db-port"`
	DbEncryptionPassword string `yaml:"db-encryption-password"`
}
