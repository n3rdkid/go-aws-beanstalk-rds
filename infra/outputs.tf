output "rds_db_host" {
  value = module.rds_go_infra_dev.rds_db_host
}
output "beanstalk_env_host" {
  value = module.beanstalk_environment_go_infra_dev.beanstalk_environment_endpoint
}