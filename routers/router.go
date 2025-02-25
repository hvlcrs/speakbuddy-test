package routers

import (
	"speakbuddy/routers/api"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	_ "speakbuddy/docs"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Init() *gin.Engine {
	r := gin.Default()

	audio := r.Group("/audio")
	audio.GET("/user/:user_id/phrase/:phrase_id/:audio_format", api.GetAudio)
	audio.POST("/user/:user_id/phrase/:phrase_id", api.UploadAudio)

	// Use swagger for API documentation
	swagger := r.Group("/swagger")
	swagger.GET("/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.Use(cors.Default())

	return r
}
