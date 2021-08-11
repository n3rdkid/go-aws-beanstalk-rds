module "beanstalk_app_project" {
    source = "./modules/beanstalk_application"
    aws_elastic_beanstalk_application_name=var.aws_elastic_beanstalk_application_name
    aws_elastic_beanstalk_application_description= var.aws_elastic_beanstalk_application_description
}

module "beanstalk_environment_go_infra_dev" {
      source = "./modules/beanstalk_environment"
      aws_elastic_beanstalk_application_name = module.beanstalk_app_project.beanstalk_application_name
      beanstalk_environment_name = "go-infra-dev"   
      beanstalk_environment_solution_stack_name=var.beanstalk_environment_solution_stack_name
}

module "rds_go_infra_dev" {
      source = "./modules/rds"
      allocated_storage    = 20
  engine               = "mysql"
  engine_version       = "8.0"
  instance_class       = "db.t2.micro"
  apply_immediately = true
  name                 = "aws_test"
  port = 3306
  username             = "FAKE_USER"
  password             = "FAKE_PASSWORD"
  # CHECK THIS
#   parameter_group_name = var.rds_parameter_group_name
  # NEED TO RESEARCH MORE ABOUT THESE THEREE
  db_subnet_group_name = "default-vpc-*"
  # NEED MORE RESEARCH WHY EBS Inbound rules need to be addded to default 
  vpc_security_group_ids = ["CREATE YOUR OWN SECURITY GROUP MAY BE" ]
  skip_final_snapshot  = true
  deletion_protection = true
    
}
