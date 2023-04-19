package console

type Console interface {
	RegisterCommand(command Command)
	Run()
}
