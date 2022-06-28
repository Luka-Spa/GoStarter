package controller

import (
	"net/http"
	"strconv"

	"github.com/Luka-Spa/GoAPI/logic"
	"github.com/Luka-Spa/GoAPI/model"
	"github.com/Luka-Spa/GoAPI/repository"
	"github.com/gin-gonic/gin"
)

func (r *httpRouter) InitPerson(logic *logic.PersonLogic) {
	r.api.GET("/person", func(c *gin.Context) {
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
		var people = logic.All(qp)
		if people == nil {
			people = make([]model.Person, 0)
		}
		c.JSON(http.StatusOK, people)
	})
	r.api.POST("/person", func(c *gin.Context) {
		var input model.Person
		if err := c.ShouldBindJSON(&input); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		person, err := logic.Create(input)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Oops, something went wrong"})
			return
		}
		c.JSON(http.StatusOK, person)
	})
	r.api.PUT("/person/:id", func(c *gin.Context) {
		var id = c.Param("id")
		var input model.Person
		if err := c.ShouldBindJSON(&input); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		person, err := logic.Update(id, input)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Oops, something went wrong"})
			return
		}
		c.JSON(http.StatusOK, person)
	})
	r.api.DELETE("/person/:id", func(c *gin.Context) {
		err := logic.Delete(c.Param("id"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Oops, something went wrong"})
			return
		}
		c.AbortWithStatus(http.StatusOK)
	})
}
