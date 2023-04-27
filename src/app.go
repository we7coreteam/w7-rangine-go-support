package support

import (
	"github.com/asaskevich/EventBus"
	"github.com/golobby/container/v3/pkg/container"
	"github.com/spf13/viper"
	"github.com/we7coreteam/w7-rangine-go-support/src/console"
	"github.com/we7coreteam/w7-rangine-go-support/src/logger"
	"github.com/we7coreteam/w7-rangine-go-support/src/server"
)

type App interface {
	GetContainer() container.Container
	GetConfig() *viper.Viper
	GetEvent() EventBus.Bus
	GetLoggerFactory() logger.Factory
	GetConsole() console.Console
	GetServerFactory() server.Factory
}
