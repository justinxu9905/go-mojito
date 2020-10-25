package controller

import (
	"net/http"
	"mojito/model"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user model.User
	c.BindJSON(&user)
	err := model.CreateUser(&user)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func UpdateUser(c *gin.Context) {
	var user model.User
	id := c.Params.ByName("id")
	err := model.GetUser(&user, id)
	if err != nil {
		c.JSON(http.StatusNotFound, user)
	}
	c.BindJSON(&user)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func GetUser(c *gin.Context) {
	id := c.Params.ByName("id")
	var user model.User
	err := model.GetUser(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func DeleteUser(c *gin.Context) {
	var user model.User
	id := c.Params.ByName("id")
	err := model.DeleteUser(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id:" + id: "deleted"})
	}
}


func CreateLost(c *gin.Context) {
	var lost model.Lost
	c.BindJSON(&lost)
	err := model.CreateLost(&lost)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, lost)
	}
}

func UpdateLost(c *gin.Context) {
	var lost model.Lost
	id := c.Params.ByName("id")
	err := model.GetLost(&lost, id)
	if err != nil {
		c.JSON(http.StatusNotFound, lost)
	}
	c.BindJSON(&lost)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, lost)
	}
}

func GetLost(c *gin.Context) {
	id := c.Params.ByName("id")
	var lost model.Lost
	err := model.GetLost(&lost, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, lost)
	}
}

func DeleteLost(c *gin.Context) {
	var lost model.Lost
	id := c.Params.ByName("id")
	err := model.DeleteLost(&lost, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id:" + id: "deleted"})
	}
}


func CreateFound(c *gin.Context) {
	var found model.Found
	c.BindJSON(&found)
	err := model.CreateFound(&found)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, found)
	}
}

func UpdateFound(c *gin.Context) {
	var found model.Found
	id := c.Params.ByName("id")
	err := model.GetFound(&found, id)
	if err != nil {
		c.JSON(http.StatusNotFound, found)
	}
	c.BindJSON(&found)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, found)
	}
}

func GetFound(c *gin.Context) {
	id := c.Params.ByName("id")
	var found model.Found
	err := model.GetFound(&found, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, found)
	}
}

func DeleteFound(c *gin.Context) {
	var found model.Found
	id := c.Params.ByName("id")
	err := model.DeleteFound(&found, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id:" + id: "deleted"})
	}
}