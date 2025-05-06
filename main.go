package main

import (
	"todo-case/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default();

	routes.Setup(r)
	r.Run(":8080")
}
