provider "aws" {
  region = "us-east-1"
  shared_credentials_files = [
    "~/.aws/credentials"
  ]
}

resource "aws_security_group" "sg_web" {
  name        = "WebSG"
  description = "Allows HTTP and SSH traffic"
  vpc_id      = aws_default_vpc.vpc_default.id

  ingress {
    from_port   = 80
    to_port     = 80
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
    description = "Allow incoming HTTP connections"
  }
  ingress {
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
    description = "Allow incoming SSH connections"
  }
  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = {
    Name    = "sg_web"
    Project = local.Project
    Version = local.Version
  }
}
