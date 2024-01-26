package main

import "github.com/CubicGames/cubic-btc-wallet-server/api/routes"

func main() {
	// Our server will live in the routes package
	routes.Run()
}
