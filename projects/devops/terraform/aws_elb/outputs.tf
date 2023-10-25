
output "instance_public01_subnet_ip" {
  value = aws_instance.instance_public01.subnet_id
}

output "instance_public02_subnet_ip" {
  value = aws_instance.instance_public02.subnet_id
}

output "instance_public03_subnet_ip" {
  value = aws_instance.instance_public03.subnet_id
}

output "load_balancer_dns" {
  value = aws_lb.lb_http_main.dns_name
}
