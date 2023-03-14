provider "aws" {
  region = "us-east-1"
  shared_credentials_files = [
    "~/.aws/credentials"
  ]
}

data "aws_ami" "amazon-linux-2" {
  most_recent = true
  owners      = ["amazon"]

  filter {
    name   = "name"
    values = ["amzn2-ami-hvm*"]
  }
}

resource "aws_security_group" "web_sg" {
  name        = "WebSG"
  description = "Allows HTTP and SSH traffic"

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
    Name    = "WebSG"
    Project = "TFHelloWorld"
  }
}

resource "aws_instance" "vm01" {
  ami                    = data.aws_ami.amazon-linux-2.id
  instance_type          = "t2.micro"
  key_name               = aws_key_pair.key_pair.key_name
  vpc_security_group_ids = [aws_security_group.web_sg.id]

  user_data = file("startup.sh")

  tags = {
    Name    = "Server01"
    Project = "TFHelloWorld"
  }
}

output "vm01_public_ip" {
  value = aws_instance.vm01.public_ip
}
