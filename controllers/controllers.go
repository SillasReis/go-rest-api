package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/SillasReis/go-rest-api/database"
	"github.com/SillasReis/go-rest-api/models"
	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	c.JSON(http.StatusOK, "Home Page")
}

func ListPersonas(c *gin.Context) {
	var p []models.Persona
	database.DB.Find(&p)

	c.JSON(http.StatusOK, p)
}

func GetPersona(c *gin.Context) {
	id := c.Param("id")

	var p models.Persona
	database.DB.First(&p, id)

	c.JSON(http.StatusOK, p)
}

func SetPersona(c *gin.Context) {
	var persona models.Persona
	json.NewDecoder(c.Request.Body).Decode(&persona)

	database.DB.Create(&persona)

	c.JSON(http.StatusCreated, persona)
}

func DeletePersona(c *gin.Context) {
	id := c.Param("id")

	var persona models.Persona
	database.DB.Delete(&persona, id)

	c.JSON(http.StatusOK, persona)
}

func UpdatePersona(c *gin.Context) {
	id := c.Param("id")

	var persona models.Persona
	var updatedPersona models.Persona

	json.NewDecoder(c.Request.Body).Decode(&updatedPersona)

	database.DB.First(&persona, id)

	database.DB.Model(&persona).Update("name", updatedPersona.Name)
	database.DB.Model(&persona).Update("history", updatedPersona.History)

	c.JSON(http.StatusOK, persona)
}
