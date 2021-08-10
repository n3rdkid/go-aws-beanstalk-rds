package controller

import (
	"go-aws-beanstalk-rds/api/service"
	"go-aws-beanstalk-rds/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// UserController ->
type UserController struct {
	service service.UserService
}

// NewUserController -> NewUserController constructor
func NewUserController(userService service.UserService) UserController {
	return UserController{
		service: userService,
	}
}

// Save ->
func (controller *UserController) Save(c *gin.Context) {
	log.Println("Inside Save User")
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err1 := controller.service.Save(user)
	if err1 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err1.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": user})
}

// GetAll ->
func (controller *UserController) GetAll(c *gin.Context) {
	log.Println("Inside GetAll users")
	users, err := controller.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": users,
	})
}

// GetByID ->
func (controller *UserController) GetByID(c *gin.Context) {
	log.Println("Inside GetByID")
	id := c.Param("id")
	idx, err := (strconv.ParseUint(id, 10, 64))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	user, err := controller.service.GetByID(uint(idx))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}
