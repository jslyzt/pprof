package main

import (
	"github.com/jslyzt/pprof"
	"github.com/lerryxiao/gin"
)

func main() {
	router := gin.Default()
	pprof.Register(router)
	router.Run(":8080")
}
