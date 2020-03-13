package main

import "github.com/y-transport-server/router"

func main() {
	r := router.InitRouter()
	r.Run(":9090")
}
