package controller

import (
	"net/http"
	"strconv"

	"github.com/Luka-Spa/GoAPI/logic"
	"github.com/Luka-Spa/GoAPI/model"
	"github.com/Luka-Spa/GoAPI/repository"
	"github.com/gin-gonic/gin"
)

var lc *logic.PersonLogic

func (r *httpRouter) InitPerson(logic *logic.PersonLogic) {
	lc = logic
	// Example Authorization: UseAuthorisation([]string{"read:users"})
	r.api.GET("/person", handleGetAll)
	r.api.GET("/person/:id", handleGetById)
	r.api.POST("/person", handleCreate)
	r.api.PUT("/person/:id", handleUpdate)
	r.api.DELETE("/person/:id", handleDelete)
}

func handleGetAll(c *gin.Context) {
	var qp = repository.QueryParams{}
	l := c.Query("limit")
	if len(l) > 0 {
		limit, err := strconv.Atoi(l)
		if err != nil {
			c.AbortWithStatusJSON(400, "invalid query parameter value: limit")
			return
		}
		qp.Limit = limit
	}
	p := c.Query("page")
	if len(p) > 0 {
		page, err := strconv.Atoi(p)
		if err != nil {
			c.AbortWithStatusJSON(400, "invalid query parameter value: page")
			return
		}
		qp.Page = page
	}
	var people = lc.All(qp)
	if people == nil {
		people = make([]model.Person, 0)
	}
	c.JSON(http.StatusOK, people)
}

func handleGetById(c *gin.Context) {
	var id = c.Param("id")
	var person, err = lc.ById(id)
	if err != nil {
		c.AbortWithStatusJSON(404, gin.H{"message": model.ErrNotFound.Error()})
		return
	}
	c.JSON(http.StatusOK, person)
}
func handleCreate(c *gin.Context) {
	var input model.Person
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	person, err := lc.Create(input)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Oops, something went wrong"})
		return
	}
	c.JSON(http.StatusOK, person)
}
func handleUpdate(c *gin.Context) {
	var id = c.Param("id")
	var input model.Person
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	person, err := lc.Update(id, input)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Oops, something went wrong"})
		return
	}
	c.JSON(http.StatusOK, person)
}
func handleDelete(c *gin.Context) {
	err := lc.Delete(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Oops, something went wrong"})
		return
	}
	c.AbortWithStatus(http.StatusOK)
}
