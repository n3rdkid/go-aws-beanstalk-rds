module "beanstalk_app_project" {
    source = "./modules/beanstalk_application"
    aws_elastic_beanstalk_application_name=var.aws_elastic_beanstalk_application_name
    aws_elastic_beanstalk_application_description= var.aws_elastic_beanstalk_application_description
    aws_elastic_beanstalk_service_name = var.aws_elastic_beanstalk_service_name
}

module "beanstalk_environment_go_infra_dev" {
      source = "./modules/beanstalk_environment"
      aws_elastic_beanstalk_application_name = module.beanstalk_app_project.beanstalk_application_name
      beanstalk_environment_name = "go-infra-dev"   
      beanstalk_environment_solution_stack_name=var.beanstalk_environment_solution_stack_name
      beanstalk_environment_security_group = var.beanstalk_environment_security_group

}

module "rds_go_infra_dev" {
  source = "./modules/rds"
  rds_allocated_storage    = var.rds_allocated_storage
  rds_engine               = var.rds_engine
  rds_engine_version       = var.rds_engine_version
  rds_instance_class       = var.rds_instance_class
  rds_db_name                 = var.rds_db_name
  rds_port = var.rds_port
  rds_username             = var.rds_username
  rds_password             = var.rds_password
  rds_parameter_group_name = var.rds_parameter_group_name
  rds_db_subnet_group_name = var.rds_db_subnet_group_name
  rds_vpc_security_group_ids = var.rds_vpc_security_group_ids  
}
