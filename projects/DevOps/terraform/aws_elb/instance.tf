data "aws_ami" "amazon_linux_2" {
  most_recent = true
  owners      = ["amazon"]

  filter {
    name   = "name"
    values = ["amzn2-ami-hvm*"]
  }
}

resource "aws_instance" "instance_public01" {
  ami               = data.aws_ami.amazon_linux_2.id
  instance_type     = "t2.micro"
  availability_zone = "us-east-1a"
  key_name          = aws_key_pair.key_pair.key_name

  vpc_security_group_ids = [aws_security_group.sg_web.id]

  user_data = file("startup.sh")

  tags = {
    Name    = "Public01"
    Project = local.Project
    Version = local.Version
  }
}

resource "aws_instance" "instance_public02" {
  ami               = data.aws_ami.amazon_linux_2.id
  instance_type     = "t2.micro"
  availability_zone = "us-east-1b"
  key_name          = aws_key_pair.key_pair.key_name

  vpc_security_group_ids = [aws_security_group.sg_web.id]

  user_data = file("startup.sh")

  tags = {
    Name    = "Public02"
    Project = local.Project
    Version = local.Version
  }
}

resource "aws_instance" "instance_public03" {
  ami               = data.aws_ami.amazon_linux_2.id
  instance_type     = "t2.micro"
  availability_zone = "us-east-1c"
  key_name          = aws_key_pair.key_pair.key_name

  vpc_security_group_ids = [aws_security_group.sg_web.id]

  user_data = file("startup.sh")

  tags = {
    Name    = "Public03"
    Project = local.Project
    Version = local.Version
  }
}
