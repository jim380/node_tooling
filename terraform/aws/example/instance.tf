resource "aws_key_pair" "main_key" {
  key_name   = "terra"
  public_key = file(var.PATH_TO_PUBLIC_KEY)
}

resource "aws_instance" "dev_server" {
  ami           = var.AMIS[var.AWS_REGION]
  instance_type = "t2.micro"
  key_name      = aws_key_pair.main_key.key_name
  # security group
  security_groups = ["aws_security_group.allow-ssh.id"]
  # user data
  user_data = data.template_cloudinit_config.cloudinit.rendered

  root_block_device {
    volume_size = 500
    volume_type = "gp2"
    delete_on_termination = true
  }

  provisioner "file" {
    source      = "script.sh"
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

