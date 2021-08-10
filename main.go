package main

import (
	"go-aws-beanstalk-rds/bootstrap"

	"go.uber.org/fx"
)

func main() {
	fx.New(bootstrap.Module).Run()
}
