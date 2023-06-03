package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"largo/bootstrap"
)

func main() {
	r := gin.New()

	bootstrap.SetupRoute(r)

	err := r.Run()
	if err != nil {
		fmt.Println(err.Error())
	}
}
