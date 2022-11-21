package main

import (
	test "example/basic"
	// encode "example/encode"

	util "example/util"
	"fmt"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	startTime := time.Now()
	fmt.Println("main")
	util.LoadEnv()
	// db.ConnectDb()

	/* ######### test basic code  ######### */
	// test.MainTest()

	/* ######### test encode ######### */
	// encode.MainEncode()
	initRouter()
	endTime := time.Since(startTime)
	fmt.Printf("\n time : %g", endTime.Seconds())
}

func initRouter() {
	r := gin.New()
	// Set a lower memory limit for multipart forms (default is 32 MiB)
	r.MaxMultipartMemory = 1

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	r.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())

	// - PUT and PATCH methods
	// - Origin header
	// - Preflight requests cached for 20 min
	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type"},
		AllowCredentials: false,
		MaxAge:           10 * time.Minute,
	}))

	// LoggerWithFormatter middleware will write the logs to gin.DefaultWriter
	// By default gin.DefaultWriter = os.Stdout
	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {

		// your custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))

	v1 := r.Group("/api/v1")
	{
		mainApi := v1.Group("/main")
		mainApi.Use(AuthRequired)
		{
			mainApi.GET("/db", test.GetDataDbMongo)
			mainApi.GET("/test", test.MainTestGetApi)
			mainApi.GET("/test/:id", test.MainTestGetApiID)
			mainApi.POST("/test", test.MainTestPostApi)
			mainApi.GET("/excel", test.Excel)
			mainApi.POST("/excel", test.Excel)
		}
	}

	// r.GET("/test", test.MainTestGetApi)
	// r.GET("/test/:id", test.MainTestGetApiID)
	// r.POST("/test", test.MainTestPostApi)

	r.Run(":8080")
}

func AuthRequired(c *gin.Context) {

}
