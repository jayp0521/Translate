//go:build wireinject
// +build wireinject

package env

import (
	"os"

	"github.com/google/wire"
	"github.com/jayp0521/Translate/utils"
)

var superset = wire.NewSet(
	utils.SuperSet,
)

var SuperSet = wire.NewSet(
	InjectEnvLoad,
	wire.Bind(new(EnvLoader), new(EnvLoad)),
)

func provideEnvSecret() envSecret {
	secret := os.Getenv("TRANSLATE_KEY")
	if len(secret) == 0 {
		panic("TRANSLATE_KEY is not defined!")
	}
	return envSecret(secret)
}

func provideEnvFileName() envFileName {
	envEnv := os.Getenv("ENV")
	switch envEnv {
	case "PROD":
		return envFileName(".env.production")
	default:
		return envFileName(".env.production")

	}
}

func InjectEnvLoad() EnvLoad {
	panic(wire.Build(
		superset,
		provideEnvSecret,
		provideEnvFileName,
		wire.Struct(new(EnvLoad), "*"),
	))
}
