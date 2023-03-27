resource "aws_vpc" "vpc_main" {
  cidr_block = "10.0.0.0/16"

  enable_dns_hostnames = true

  tags = {
    Name    = "vpc_main"
    Project = local.Project
    Version = local.Version
  }
}

resource "aws_subnet" "public01" {
  vpc_id            = aws_vpc.vpc_main.id
  cidr_block        = "10.0.1.0/24"
  availability_zone = "us-east-1a"

  map_public_ip_on_launch = true

  tags = {
    Name    = "public01"
    Project = local.Project
    Version = local.Version
  }
}

resource "aws_subnet" "public02" {
  vpc_id            = aws_vpc.vpc_main.id
  cidr_block        = "10.0.2.0/24"
  availability_zone = "us-east-1b"

  map_public_ip_on_launch = true

  tags = {
    Name    = "public02"
    Project = local.Project
    Version = local.Version
  }
}

resource "aws_subnet" "private01" {
  vpc_id            = aws_vpc.vpc_main.id
  cidr_block        = "10.0.3.0/24"
  availability_zone = "us-east-1a"

  tags = {
    Name    = "private01"
    Project = local.Project
    Version = local.Version
  }
}

resource "aws_subnet" "private02" {
  vpc_id            = aws_vpc.vpc_main.id
  cidr_block        = "10.0.4.0/24"
  availability_zone = "us-east-1b"

  tags = {
    Name    = "private02"
    Project = local.Project
    Version = local.Version
  }
}

resource "aws_default_route_table" "rt_default" {
  default_route_table_id = aws_vpc.vpc_main.default_route_table_id

  route {
    cidr_block = "0.0.0.0/0"
    gateway_id = aws_internet_gateway.igw_main.id
  }

  tags = {
    Name    = "rt_default"
    Project = local.Project
    Version = local.Version
  }
}

resource "aws_route_table" "rt_private" {
  vpc_id = aws_vpc.vpc_main.id

  route {
    cidr_block     = "0.0.0.0/0"
    nat_gateway_id = aws_nat_gateway.nat_public01.id
  }

  tags = {
    Name    = "rt_private"
    Project = local.Project
    Version = local.Version
  }
}

resource "aws_route_table_association" "rt_assoc_private01" {
  subnet_id      = aws_subnet.private01.id
  route_table_id = aws_route_table.rt_private.id
}

resource "aws_route_table_association" "rt_assoc_private02" {
  subnet_id      = aws_subnet.private02.id
  route_table_id = aws_route_table.rt_private.id
}

resource "aws_internet_gateway" "igw_main" {
  vpc_id = aws_vpc.vpc_main.id

  tags = {
    Name    = "igw_main"
    Project = local.Project
    Version = local.Version
  }
}

resource "aws_eip" "eip_public01" {
  vpc = true

  tags = {
    Name    = "eip_public01"
    Project = local.Project
    Version = local.Version
  }

  depends_on = [
    aws_internet_gateway.igw_main
  ]
}

resource "aws_nat_gateway" "nat_public01" {
  subnet_id     = aws_subnet.public01.id
  allocation_id = aws_eip.eip_public01.id

  connectivity_type = "public"

  tags = {
    Name    = "nat_public01"
    Project = local.Project
    Version = local.Version
  }

  depends_on = [
    aws_internet_gateway.igw_main
  ]
}
