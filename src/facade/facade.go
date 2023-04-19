package facade

import (
	"github.com/asaskevich/EventBus"
	"github.com/gin-gonic/gin/binding"
	ut "github.com/go-playground/universal-translator"
	"github.com/golobby/container/v3/pkg/container"
	"github.com/spf13/viper"
	"github.com/we7coreteam/w7-rangine-go-support/src/contract"
	"github.com/we7coreteam/w7-rangine-go-support/src/contract/database"
	"github.com/we7coreteam/w7-rangine-go-support/src/contract/logger"
	"github.com/we7coreteam/w7-rangine-go-support/src/contract/redis"
)

var app contract.App

func SetApp(appI contract.App) {
	app = appI
}

func GetContainer() container.Container {
	return app.GetContainer()
}

func GetConfig() *viper.Viper {
	return app.GetConfig()
}

func GetEvent() EventBus.Bus {
	return app.GetEvent()
}

func GetLoggerFactory() logger.Factory {
	return app.GetLoggerFactory()
}

func GetRedisFactory() redis.Factory {
	var redisFactory redis.Factory
	_ = GetContainer().NamedResolve(&redisFactory, "redis-factory")

	return redisFactory
}

func GetDbFactory() database.Factory {
	var dbFactory database.Factory
	_ = GetContainer().NamedResolve(&dbFactory, "db-factory")

	return dbFactory
}

func GetTranslator() ut.Translator {
	var translator ut.Translator
	_ = GetContainer().NamedResolve(&translator, "translator")

	return translator
}

func GetValidator() binding.StructValidator {
	return binding.Validator
}