resource "aws_key_pair" "main_key" {
  key_name   = "terra"
  public_key = file(var.PATH_TO_PUBLIC_KEY)
}

resource "aws_instance" "dev_server" {
  ami           = var.AMIS[var.AWS_REGION]
  instance_type = var.NODE_INSTANCE_MODEL
  key_name      = aws_key_pair.main_key.key_name
  # security group
  vpc_security_group_ids = [aws_security_group.dev-default.id]
  # user data
  user_data = data.template_cloudinit_config.cloudinit.rendered

  root_block_device {
    volume_size = var.NODE_INSTANCE_VOLUME
    volume_type = "gp2"
    delete_on_termination = true
  }

  provisioner "file" {
    source      = "scripts/script.sh"
    destination = "~/script.sh"
  }
  provisioner "remote-exec" {
    inline = [
      "chmod +x ~/script.sh",
      "sudo sed -i -e 's/\r$//' ~/script.sh",  # Remove the spurious CR characters.
      "sudo ~/script.sh",
    ]
  }
  connection {
    host        = coalesce(self.public_ip, self.private_ip)
    type        = "ssh"
    user        = var.INSTANCE_USERNAME
    private_key = file(var.PATH_TO_PRIVATE_KEY)
  }
}

