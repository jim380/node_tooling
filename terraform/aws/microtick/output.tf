output "ip" {
  description = "instance public ip"
  value = {
    instance_public_ip = aws_instance.dev_mt.public_ip
  }
}