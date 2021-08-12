terraform {
  required_version = ">= v1.0.3"
}

resource "aws_db_instance" "default" {
  allocated_storage    = var.rds_allocated_storage
  engine               = var.rds_engine
  engine_version       = var.rds_engine_version
  instance_class       = var.rds_instance_class
  apply_immediately = true
  name                 = var.rds_db_name
  port = var.rds_port
  username             = var.rds_username
  password             = var.rds_password
  # CHECK THIS
  parameter_group_name = var.rds_parameter_group_name
  # NEED TO RESEARCH MORE ABOUT THESE THEREE
  db_subnet_group_name = var.rds_db_subnet_group_name
  # NEED MORE RESEARCH WHY EBS Inbound rules need to be addded to default 
  vpc_security_group_ids = var.rds_vpc_security_group_ids
  skip_final_snapshot  = true
  deletion_protection = true 
}