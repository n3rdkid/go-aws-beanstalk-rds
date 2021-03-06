package bootstrap

import (
	"context"
	"go-aws-beanstalk-rds/api/controller"
	"go-aws-beanstalk-rds/api/repository"
	"go-aws-beanstalk-rds/api/routes"
	"go-aws-beanstalk-rds/api/service"
	"go-aws-beanstalk-rds/lib"
	"log"

	"go.uber.org/fx"
)

// Module exported for initializing application
var Module = fx.Options(
	lib.Module,
	controller.Module,
	routes.Module,
	service.Module,
	repository.Module,
	fx.Invoke(bootstrap),
)

func bootstrap(
	lifecycle fx.Lifecycle,
	handler lib.RequestHandler,
	routes routes.Routes,
	env lib.Env,
	database lib.Database,
) {
	conn, _ := database.DB.DB()

	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			conn.SetMaxOpenConns(10)
			go func() {
				routes.Setup()
				err := handler.Gin.Run(":" + env.ServerPort)
				if err != nil {
					log.Print("failed to start gin")
				}
			}()
			return nil
		},
		OnStop: func(context.Context) error {
			err := conn.Close()
			if err != nil {
				log.Print("failed to stop db connection")
			}
			return nil
		},
	})
}
