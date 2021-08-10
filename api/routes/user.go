package routes

import (
	"go-aws-beanstalk-rds/api/controller"
	"go-aws-beanstalk-rds/api/repository"
	"go-aws-beanstalk-rds/lib"
	"log"
)

// UserRoutes ->
type UserRoutes struct {
	handler        lib.RequestHandler
	userController controller.UserController
	userRepository repository.UserRepository
}

// Setup -> user routes
func (s UserRoutes) Setup() {

	if err := s.userRepository.Migrate(); err != nil {
		log.Fatal("User migrate err", err)
	}
	api := s.handler.Gin.Group("/api")
	{
		api.GET("/user", s.userController.GetAll)
		api.POST("/user", s.userController.Save)
		api.GET("/user/:id", s.userController.GetByID)
	}
}

// NewUserRoutes -> Create new user controller
func NewUserRoutes(handler lib.RequestHandler, userController controller.UserController, userRepository repository.UserRepository) UserRoutes {
	return UserRoutes{
		handler:        handler,
		userController: userController,
		userRepository: userRepository,
	}
}
