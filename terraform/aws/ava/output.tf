output "ip" {
  description = "instance public ip"
  value = {
    instance_public_ip = aws_instance.ava_node.public_ip
  }
}