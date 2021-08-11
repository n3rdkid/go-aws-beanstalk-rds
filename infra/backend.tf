terraform {
  backend "s3" {
    bucket = "go-infra-tfstate"
    key    = "terraform.tfstate"
    region = "us-east-1"
    profile = "default"
  }
}