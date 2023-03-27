data "aws_ami" "amazon_linux_2" {
  most_recent = true
  owners      = ["amazon"]

  filter {
    name   = "name"
    values = ["amzn2-ami-hvm*"]
  }
}

resource "aws_instance" "instance_public01" {
  ami           = data.aws_ami.amazon_linux_2.id
  instance_type = "t2.micro"
  key_name      = aws_key_pair.key_pair.key_name

  subnet_id              = aws_subnet.public01.id
  vpc_security_group_ids = [aws_security_group.sg_all.id]

  user_data = file("startup.sh")

  tags = {
    Name    = "Public01"
    Project = local.Project
    Version = local.Version
  }
}

resource "aws_instance" "instance_public02" {
  ami           = data.aws_ami.amazon_linux_2.id
  instance_type = "t2.micro"
  key_name      = aws_key_pair.key_pair.key_name

  subnet_id              = aws_subnet.public02.id
  vpc_security_group_ids = [aws_security_group.sg_all.id]

  user_data = file("startup.sh")

  tags = {
    Name    = "Public02"
    Project = local.Project
    Version = local.Version
  }
}

resource "aws_instance" "instance_private01" {
  ami           = data.aws_ami.amazon_linux_2.id
  instance_type = "t2.micro"
  key_name      = aws_key_pair.key_pair.key_name

  subnet_id              = aws_subnet.private01.id
  vpc_security_group_ids = [aws_security_group.sg_all.id]

  user_data = file("startup.sh")

  tags = {
    Name    = "Private01"
    Project = local.Project
    Version = local.Version
  }
}

resource "aws_instance" "instance_private02" {
  ami           = data.aws_ami.amazon_linux_2.id
  instance_type = "t2.micro"
  key_name      = aws_key_pair.key_pair.key_name

  subnet_id              = aws_subnet.private02.id
  vpc_security_group_ids = [aws_security_group.sg_all.id]

  user_data = file("startup.sh")

  tags = {
    Name    = "Private02"
    Project = local.Project
    Version = local.Version
  }
}
