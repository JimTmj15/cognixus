package main

import (
	"cognixus/api/routes"
	"cognixus/bootstrap"

	"github.com/gin-gonic/gin"
)

func main() {

	app := bootstrap.App()
	env := app.Env
	db := app.Db

	gin := gin.Default()

	routes.Setup(db, gin)

	gin.Run(env.ServerHost)
}
