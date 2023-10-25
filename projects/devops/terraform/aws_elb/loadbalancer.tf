resource "aws_lb_target_group" "lb_tg_http_main" {
  name        = "lb-tg-http-main"
  target_type = "instance"
  vpc_id      = aws_default_vpc.vpc_default.id

  port     = 80
  protocol = "HTTP"

  tags = {
    Name    = "lb_tg_http_main"
    Project = local.Project
    Version = local.Version
  }
}

resource "aws_lb_target_group_attachment" "lb_group_attach_1" {
  target_group_arn = aws_lb_target_group.lb_tg_http_main.arn
  target_id        = aws_instance.instance_public01.id
  port             = 80
}

resource "aws_lb_target_group_attachment" "lb_group_attach_2" {
  target_group_arn = aws_lb_target_group.lb_tg_http_main.arn
  target_id        = aws_instance.instance_public02.id
  port             = 80
}

resource "aws_lb_target_group_attachment" "lb_group_attach_3" {
  target_group_arn = aws_lb_target_group.lb_tg_http_main.arn
  target_id        = aws_instance.instance_public03.id
  port             = 80
}

resource "aws_lb" "lb_http_main" {
  name               = "lb-http-main"
  internal           = false
  load_balancer_type = "application"
  security_groups    = [aws_security_group.sg_web.id]
  subnets = [
    aws_default_subnet.sn_default_1.id,
    aws_default_subnet.sn_default_2.id,
    aws_default_subnet.sn_default_3.id,
  ]

  tags = {
    Name    = "lb_http_main"
    Project = local.Project
    Version = local.Version
  }
}

resource "aws_lb_listener" "lis_http_main" {
  load_balancer_arn = aws_lb.lb_http_main.arn
  port              = "80"
  protocol          = "HTTP"

  default_action {
    type             = "forward"
    target_group_arn = aws_lb_target_group.lb_tg_http_main.arn
  }

  tags = {
    Name    = "lis_http_main"
    Project = local.Project
    Version = local.Version
  }
}
