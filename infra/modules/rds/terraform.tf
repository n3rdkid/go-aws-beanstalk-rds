variable "rds_allocated_storage" {
  type=number
}
variable "rds_engine" {
  type=string
}
variable "rds_engine_version" {
  type=string
}
variable "rds_instance_class" {
  type = string
}
variable "rds_parameter_group_name" {
  type = string
}
variable "rds_db_name" {
  type = string
}
variable "rds_port" {
  type = string
}

variable "rds_username" {
  type = string
}
variable "rds_password" {
  type = string
}
