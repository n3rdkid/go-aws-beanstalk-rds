output "rds_db_host" {
  value=aws_db_instance.default.endpoint
}