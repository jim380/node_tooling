resource "aws_security_group" "dev-default" {
  name        = "ava-default"
  //description = "security group that allows ssh and all egress traffic"
  // vpc_id      = aws_vpc.default.id

  ingress {
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"] # use vpc next
    description = "allow ssh"
  }

  ingress {
    from_port   = "80"
    to_port     = "80"
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"] # use vpc next
    description = "allow http"
  }

  ingress {
    from_port   = "9090"
    to_port     = "9090"
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"] # use vpc next
    description = "prometheus metrics port"
  }

  ingress {
    from_port   = "9100"
    to_port     = "9100"
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"] # use vpc next
    description = "node exporter port"
  }

  ingress {
    from_port   = "9652"
    to_port     = "9652"
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"] # use vpc next
    description = "ava p2p port"
  }

  egress {
    from_port   = "80"
    to_port     = "80"
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
    description = "http"
  }

  egress {
    from_port   = "443"
    to_port     = "443"
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
    description = "https"
  }

  tags = {
    Name = "dev-default"
  }

  lifecycle {
    create_before_destroy = true
  }
}

