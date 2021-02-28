package main

import (
	"api-gateway/pkg/gateway"
	"api-gateway/pkg/tokenserver"
)

func main() {
	tokenServer := tokenserver.CreateTokenServer()
	gatewayServer := gateway.CreateGatewayServer()

	go tokenServer.Run(":8006")
	gatewayServer.Run(":8004")
}
