resource "aws_cloudwatch_log_group" "lg_main" {
  name = "lg_main"

  tags = {
    Name    = "lg_main"
    Project = local.Project
    Version = local.Version
  }
}

resource "aws_flow_log" "fl_vpc_main" {
  traffic_type    = "ALL"
  iam_role_arn    = aws_iam_role.fl_vpc_main.arn
  log_destination = aws_cloudwatch_log_group.lg_main.arn
  vpc_id          = aws_vpc.vpc_main.id

  tags = {
    Name    = "fl_vpc_main"
    Project = local.Project
    Version = local.Version
  }
}
