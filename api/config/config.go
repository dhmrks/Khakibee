package config

import (
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/neko-neko/echo-logrus/v2/log"
)

const (
	Port         = "PORT"
	SQLString    = "SQLSTRING"
	JWTSecret    = "JWTSECRET"
	AllowOrigins = "ALLOWORIGINS"

	DefaultPort    = "8080"
	EnvFilePath    = ".env"
	EnvFileWarning = ".env file not found - working with values from system."

	SMTPUsername = "SMTPUSERNAME"
	SMTPPassword = "SMTPPASSWORD"
	SMTPHost     = "SMTPHOST"
	SMTPPort     = "SMTPPORT"
)

type Config struct {
	Port         string
	SQLString    string
	JWTSecret    string
	AllowOrigins []string
	SMTPDtls     SMTPDtls
}

type SMTPDtls struct {
	Username string
	Password string
	Host     string
	Port     string
}

// LoadEnv loads the .env file and returns config object
func LoadEnv() (config Config) {
	err := godotenv.Load(EnvFilePath)
	if err != nil {
		log.Logger().Warning(EnvFileWarning)
	}

	config.Port = os.Getenv(Port)
	if len(config.Port) == 0 {
		config.Port = DefaultPort
	}

	config.SQLString = os.Getenv(SQLString)
	config.JWTSecret = os.Getenv(JWTSecret)

	allowOrigins := os.Getenv(AllowOrigins)
	config.AllowOrigins = strings.Split(allowOrigins, ",")

	config.SMTPDtls.Username = os.Getenv(SMTPUsername)
	config.SMTPDtls.Password = os.Getenv(SMTPPassword)
	config.SMTPDtls.Host = os.Getenv(SMTPHost)
	config.SMTPDtls.Port = os.Getenv(SMTPPort)

	return
}
