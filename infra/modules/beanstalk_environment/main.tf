terraform {
  required_version = ">= v1.0.3"
}

resource "aws_elastic_beanstalk_environment" "beanstalk_environment" {
  name                = var.beanstalk_environment_name
  application         = var.aws_elastic_beanstalk_application_name
  solution_stack_name = var.beanstalk_environment_solution_stack_name
  setting {
    namespace = "aws:autoscaling:launchconfiguration"
    name      = "IamInstanceProfile"
    value   = "aws-elasticbeanstalk-ec2-role"
    }
    setting {
    namespace = "aws:autoscaling:launchconfiguration"
    name      = "SecurityGroups"
    value   = var.beanstalk_environment_security_group
    }
}