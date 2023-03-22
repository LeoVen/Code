provider "aws" {
  region = "us-east-1"
  shared_credentials_files = [
    "~/.aws/credentials"
  ]
}

resource "aws_security_group" "sg_web" {
  name        = "WebSG"
  description = "Allows HTTP and SSH traffic"
  vpc_id      = aws_vpc.vpc_main.id

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
    Project = "Networking"
  }
}

resource "aws_security_group" "sg_all" {
  name        = "Allow all SG"
  description = "Allow everything Security Group"
  vpc_id      = aws_vpc.vpc_main.id

  ingress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = {
    Name    = "sg_all"
    Project = "Networking"
  }
}
