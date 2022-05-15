package env

import (
	"os"
	"path/filepath"

	"github.com/jayp0521/Translate/encrypt"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

type EnvLoader interface {
	Load() error
}

type envSecret string

type envFileName string

type EnvLoad struct {
	log      *zap.SugaredLogger
	secret   envSecret
	fileName envFileName
}

func (env EnvLoad) Load() error {
	env.decryptEnvFile()
	env.log.Info("File decrypted")
	return godotenv.Load(string(env.fileName))
}

func (env EnvLoad) decryptEnvFile() error {
	path, _ := os.Getwd()
	_, err := encrypt.DecryptFile(filepath.Join(path, string(env.fileName)+".bin"), string(env.secret))
	return err
}
