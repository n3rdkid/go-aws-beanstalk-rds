variable "aws_elastic_beanstalk_application_name" {
  default = "go-terraform"
}

variable "aws_elastic_beanstalk_application_description" {
  default = "Go Infrastructure building"
}

variable "beanstalk_environment_solution_stack_name" {
  default = "64bit Amazon Linux 2 v3.3.4 running Go 1"
}