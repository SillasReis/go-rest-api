package routes

import (
	"log"
	"net/http"

	"github.com/SillasReis/go-rest-api/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func HandleRequest() {
	r := gin.Default()
	r.Use(cors.Default())

	r.GET("/", controllers.Home)
	r.GET("/api/personas", controllers.ListPersonas)
	r.GET("/api/personas/:id", controllers.GetPersona)

	r.POST("/api/personas", controllers.SetPersona)

	r.DELETE("/api/personas/:id", controllers.DeletePersona)

	r.PUT("/api/personas/:id", controllers.UpdatePersona)

	log.Fatal(http.ListenAndServe(":8000", r))
}
