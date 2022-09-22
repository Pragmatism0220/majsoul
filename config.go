package majsoul

// ServerAddress majsoul server config
type ServerAddress struct {
	ServerAddress  string `json:"serverAddress"`
	GatewayAddress string `json:"gatewayAddress"`
	GameAddress    string `json:"gameAddress"`
}

func generateMajsoulServerAddressList() (ServerAddressList []*ServerAddress) {
	wsURLs, _ := GetMajsoulWebSocketURLs()
	for _, ws := range wsURLs {
		server := new(ServerAddress)
		server.ServerAddress = "https://game.maj-soul.net"
		server.GatewayAddress = ws + "gateway"
		server.GameAddress = ws + "game-gateway"
		ServerAddressList = append(ServerAddressList, server)
	}
	return
}

var ServerAddressList = generateMajsoulServerAddressList()
