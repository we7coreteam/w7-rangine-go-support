package server

type Factory interface {
	RegisterServer(server Server)
	GetAllServer() map[string]Server
	GetServer(serverName string) Server
}
