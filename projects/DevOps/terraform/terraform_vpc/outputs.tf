
output "instance_public01_ip" {
  value = aws_instance.instance_public01.public_ip
}

output "instance_public02_ip" {
  value = aws_instance.instance_public02.public_ip
}
