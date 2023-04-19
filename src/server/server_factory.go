package server

var servers = make(map[string]Server)

func RegisterServer(server Server) {
	servers[server.GetServerName()] = server
}

func GetServer(serverName string) Server {
	server, exists := servers[serverName]
	if !exists {
		return nil
	}

	return server
}
