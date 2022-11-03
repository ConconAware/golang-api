package main

import (
	test "example/basic"
	// encode "example/encode"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	startTime := time.Now()
	fmt.Println("main")

	/* ######### test basic code  ######### */
	// test.MainTest()

	/* ######### test encode ######### */
	// encode.MainEncode()

	router := gin.Default()

	router.GET("/test", test.MainTestGetApi)
	router.GET("/test/:id", test.MainTestGetApiID)
	router.POST("/test", test.MainTestPostApi)

	router.Run("localhost:8080")
	endTime := time.Since(startTime)
	fmt.Printf("\n time : %g", endTime.Seconds())
}
