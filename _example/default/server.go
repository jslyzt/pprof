package main

import (
	"github.com/jslyzt/gin"
	"github.com/jslyzt/pprof"
)

func main() {
	router := gin.Default()
	pprof.Register(router)
	router.Run(":8080")
}
