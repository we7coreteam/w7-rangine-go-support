package server

type Manager interface {
	RegisterServer(server Server)
	GetAllServer() map[string]Server
	GetServer(serverName string) Server
}
