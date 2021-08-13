
terraform {
  required_version = ">= v1.0.3"
}

resource "aws_iam_role" "beanstalk_service" {
  name = var.aws_elastic_beanstalk_service_name
  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [{
    "Effect": "Allow",
    "Principal": {
      "Service": "elasticbeanstalk.amazonaws.com"
    },
    "Action": "sts:AssumeRole",
    "Condition": {
      "StringEquals": {
        "sts:ExternalId": "elasticbeanstalk"
      }
    }
  }]
}
EOF
}
resource "aws_elastic_beanstalk_application" "beanstalk_app" {
  name        = var.aws_elastic_beanstalk_application_name
  description = var.aws_elastic_beanstalk_application_description

  appversion_lifecycle {
    service_role          = aws_iam_role.beanstalk_service.arn
    max_age_in_days = 30
    delete_source_from_s3 = true
  }
}