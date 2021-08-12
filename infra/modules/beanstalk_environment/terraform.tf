variable "beanstalk_environment_name" {
  type = string
  description = "Name of the Beanstalk environment"
}

variable "aws_elastic_beanstalk_application_name" {
  type = string
  description = "Name of the containing Beanstalk application"
}
variable "beanstalk_environment_solution_stack_name" {
  type = string
  description = "Beanstalk environment solution stack name"
  
}
variable "beanstalk_environment_security_group" {
  type = string
  description = "Security Group matching the RDS Ingress"
  
}