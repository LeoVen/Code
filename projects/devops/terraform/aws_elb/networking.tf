resource "aws_default_vpc" "vpc_default" {
  tags = {
    Name    = "vpc_default"
    Project = local.Project
    Version = local.Version
  }
}

resource "aws_default_subnet" "sn_default_1" {
  availability_zone = "us-east-1a"
  tags = {
    Name    = "sn_default_1"
    Project = local.Project
    Version = local.Version
  }
}

resource "aws_default_subnet" "sn_default_2" {
  availability_zone = "us-east-1b"
  tags = {
    Name    = "sn_default_2"
    Project = local.Project
    Version = local.Version
  }
}

resource "aws_default_subnet" "sn_default_3" {
  availability_zone = "us-east-1c"
  tags = {
    Name    = "sn_default_3"
    Project = local.Project
    Version = local.Version
  }
}
