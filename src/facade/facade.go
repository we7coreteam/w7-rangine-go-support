package facade

import (
	"github.com/asaskevich/EventBus"
	"github.com/gin-gonic/gin/binding"
	ut "github.com/go-playground/universal-translator"
	"github.com/golobby/container/v3/pkg/container"
	"github.com/spf13/viper"
	"github.com/we7coreteam/w7-rangine-go-support/src"
	"github.com/we7coreteam/w7-rangine-go-support/src/console"
	"github.com/we7coreteam/w7-rangine-go-support/src/database"
	"github.com/we7coreteam/w7-rangine-go-support/src/logger"
	"github.com/we7coreteam/w7-rangine-go-support/src/redis"
	"github.com/we7coreteam/w7-rangine-go-support/src/server"
)

var app support.App

func SetApp(appI support.App) {
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

func GetConsole() console.Console {
	return app.GetConsole()
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

func RegisterServer(sev server.Server) {
	server.RegisterServer(sev)
}

func GetAllServer() map[string]server.Server {
	return server.GetAllServer()
}

func GetServer(serverName string) server.Server {
	return server.GetServer(serverName)
}
