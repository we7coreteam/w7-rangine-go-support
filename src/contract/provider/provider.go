package provider

import (
	"github.com/asaskevich/EventBus"
	"github.com/golobby/container/v3/pkg/container"
	"github.com/spf13/viper"
	"github.com/we7coreteam/w7-rangine-go-support/src/contract/console"
	"github.com/we7coreteam/w7-rangine-go-support/src/contract/logger"
)

type Provider interface {
	GetPackageName() string
	Register()
	SetContainer(container container.Container)
	GetContainer() container.Container
	SetConfig(config *viper.Viper)
	GetConfig() *viper.Viper
	SetConsole(console *console.Console)
	GetConsole() *console.Console
	SetLoggerFactory(loggerFactory *logger.Factory)
	GetLoggerFactory() *logger.Factory
	SetEvent(event EventBus.Bus)
	GetEvent() EventBus.Bus
}

type Abstract struct {
	Provider
	container     container.Container
	config        *viper.Viper
	console       console.Console
	loggerFactory logger.Factory
	event         EventBus.Bus
	PackageName   string
}

func (abstract *Abstract) GetPackageName() string {
	return abstract.PackageName
}

func (abstract *Abstract) SetContainer(container container.Container) {
	abstract.container = container
}

func (abstract *Abstract) GetContainer() container.Container {
	return abstract.container
}

func (abstract *Abstract) SetConfig(config *viper.Viper) {
	abstract.config = config
}

func (abstract *Abstract) GetConfig() *viper.Viper {
	return abstract.config
}

func (abstract *Abstract) SetConsole(console console.Console) {
	abstract.console = console
}

func (abstract *Abstract) GetConsole() console.Console {
	return abstract.console
}

func (abstract *Abstract) SetLoggerFactory(loggerFactory logger.Factory) {
	abstract.loggerFactory = loggerFactory
}

func (abstract *Abstract) GetLoggerFactory() logger.Factory {
	return abstract.loggerFactory
}

func (abstract *Abstract) SetEvent(event EventBus.Bus) {
	abstract.event = event
}

func (abstract *Abstract) GetEvent() EventBus.Bus {
	return abstract.event
}
